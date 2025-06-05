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

package vwechatpay

import (
	"context"
	"crypto/x509"

	"github.com/vogo/vogo/vlog"
	"github.com/wechatpay-apiv3/wechatpay-go/services/certificates"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func LoadPlatformCert(cli *WechatPayManager) *x509.Certificate {
	vlog.Infof("load wechat platform merchantCert")
	ctx := context.Background()
	// 发送请求，以下载微信支付平台证书为例
	// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
	svc := certificates.CertificatesApiService{Client: cli.Client}
	resp, result, err := svc.DownloadCertificates(ctx)
	if err != nil {
		vlog.Fatalf("load wechat platform merchantCert failed: %v", err)
		return nil
	}

	vlog.Infof("status=%d resp=%s", result.Response.StatusCode, resp)

	// 解析返回数据获得公钥证书
	encryptCert := resp.Data[0].EncryptCertificate
	keyText, err := utils.DecryptAES256GCM(cli.Config.MerchantAPIv3Key, *encryptCert.AssociatedData,
		*encryptCert.Nonce, *encryptCert.Ciphertext)
	if err != nil {
		vlog.Fatalf("decrypt wechat platform merchantCert failed: %v", err)
		return nil
	}

	// 解码证书
	platCert, err := utils.LoadCertificate(keyText)
	if err != nil {
		vlog.Fatalf("load wechat platform merchantCert failed: %v", err)
		return nil
	}

	return platCert
}
