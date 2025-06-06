/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package vwxcert

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"sync"
	"time"

	"github.com/vogo/vogo/vlog"
	"github.com/vogo/vwechatpay"
	"github.com/vogo/vwechatpay/vwxutils"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/services/certificates"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func init() {
	vwechatpay.PlatManagerInit = func(mgr *vwechatpay.Manager) vwechatpay.PlatManager {
		return NewPlatformManager(mgr)
	}
}

type WxPlatManager struct {
	mux            sync.Mutex
	mgr            *vwechatpay.Manager
	certificateApi certificates.CertificatesApiService
	platformCert   *x509.Certificate
	verifier       *verifiers.SHA256WithRSAVerifier
	expireTime     time.Time
}

func NewPlatformManager(mgr *vwechatpay.Manager) *WxPlatManager {
	return &WxPlatManager{
		mux:            sync.Mutex{},
		mgr:            mgr,
		certificateApi: certificates.CertificatesApiService{Client: mgr.Client},
	}
}

func (c *WxPlatManager) LoadCert() *x509.Certificate {
	if c.platformCert != nil && c.expireTime.After(time.Now()) {
		return c.platformCert
	}

	c.reloadCert()

	return c.platformCert
}

func (c *WxPlatManager) LoadVerifier() *verifiers.SHA256WithRSAVerifier {
	if c.verifier != nil && c.expireTime.After(time.Now()) {
		return c.verifier
	}

	c.reloadCert()

	return c.verifier
}

func (c *WxPlatManager) reloadCert() {
	c.mux.Lock()
	defer c.mux.Unlock()

	vlog.Infof("load wechat platform merchantCert")
	ctx := context.Background()
	// 发送请求，以下载微信支付平台证书为例
	// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
	resp, result, err := c.certificateApi.DownloadCertificates(ctx)
	if err != nil {
		vlog.Fatalf("load wechat platform merchantCert failed: %v", err)
		return
	}

	vlog.Infof("status=%d resp=%s", result.Response.StatusCode, resp)

	// 解析返回数据获得公钥证书
	encryptCert := resp.Data[0].EncryptCertificate
	keyText, err := utils.DecryptAES256GCM(c.mgr.Config.MerchantAPIv3Key, *encryptCert.AssociatedData,
		*encryptCert.Nonce, *encryptCert.Ciphertext)
	if err != nil {
		vlog.Fatalf("decrypt wechat platform merchantCert failed: %v", err)
		return
	}

	// 解码证书
	platCert, err := utils.LoadCertificate(keyText)
	if err != nil {
		vlog.Fatalf("load wechat platform merchantCert failed: %v", err)
		return
	}

	c.platformCert = platCert
	c.verifier = verifiers.NewSHA256WithRSAVerifier(core.NewCertificateMapWithList([]*x509.Certificate{platCert}))
	c.expireTime = (*resp.Data[0].ExpireTime).Add(-60 * time.Second)
}

// EncryptSensitiveInfo 使用微信支付平台证书加密敏感信息
func (c *WxPlatManager) Encrypt(plaintext string) (string, error) {
	// 确保平台证书已加载
	platformCert := c.LoadCert()

	// 获取平台证书公钥
	publicKey := platformCert.PublicKey.(*rsa.PublicKey)

	// 使用RSA公钥加密
	return vwxutils.EncryptRSA(plaintext, publicKey)
}
