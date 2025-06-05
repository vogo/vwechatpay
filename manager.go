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
	"crypto/rsa"
	"crypto/x509"
	"time"

	"github.com/vogo/vogo/vlog"
	"github.com/vogo/vogo/vos"
	"github.com/vogo/vogo/vsync/vrunner"
	"github.com/vogo/vwechatpay/vpayutils"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/certificates"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type WechatPayManager struct {
	runner             *vrunner.Runner
	Config             *Config
	merchantPrivateKey *rsa.PrivateKey
	merchantCert       *x509.Certificate
	platformCert       *x509.Certificate
	Client             *core.Client
	JsApi              jsapi.JsapiApiService
	certificateApi     certificates.CertificatesApiService
	Verifier           *verifiers.SHA256WithRSAVerifier
}

func NewManager() *WechatPayManager {
	manager := NewManagerFromEnv()

	manager.runner.Interval(func() {
		manager.platformCert = LoadPlatformCert(manager)
		manager.Verifier = verifiers.NewSHA256WithRSAVerifier(
			core.NewCertificateMapWithList([]*x509.Certificate{manager.platformCert}))
	}, 24*time.Hour)

	return manager
}

func NewManagerFromEnv() *WechatPayManager {
	cfg := &Config{
		MerchantID:           vos.EnsureEnvString("WECHAT_PAY_MERCHANT_ID"),
		MerchantCertSerialNO: vos.EnsureEnvString("WECHAT_PAY_MERCHANT_CERT_SERIAL_NO"),
		MerchantAPIv3Key:     vos.EnsureEnvString("WECHAT_PAY_MERCHANT_API_V3_KEY"),
		PrivateKeyPath:       vos.EnsureEnvString("WECHAT_PAY_PRIVATE_KEY_PATH"),
		CertPath:             vos.EnsureEnvString("WECHAT_PAY_CERT_PATH"),
		AppID:                vos.EnsureEnvString("WECHAT_PAY_APP_ID"),
	}

	manager := &WechatPayManager{
		runner: vrunner.New(),
		Config: cfg,
	}

	manager.merchantPrivateKey = buildMerchantPrivateKey(cfg)
	manager.merchantCert = buildMerchantCert(cfg)
	manager.Client = buildWechatPayClient(cfg, manager.merchantPrivateKey)
	manager.certificateApi = certificates.CertificatesApiService{Client: manager.Client}
	return manager
}

func buildWechatPayClient(cfg *Config, key *rsa.PrivateKey) *core.Client {
	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(cfg.MerchantID,
			cfg.MerchantCertSerialNO,
			key,
			cfg.MerchantAPIv3Key),
	}

	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		vlog.Fatalf("new wechat pay client err:%s", err)
	}

	return client
}

func buildMerchantPrivateKey(cfg *Config) *rsa.PrivateKey {
	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	privateKey, err := utils.LoadPrivateKeyWithPath(cfg.PrivateKeyPath)
	if err != nil {
		vlog.Fatalf("load merchant private key error: %v", err)
	}
	return privateKey
}

func buildMerchantCert(cfg *Config) *x509.Certificate {
	cert, err := utils.LoadCertificateWithPath(cfg.CertPath)
	if err != nil {
		vlog.Fatalf("load merchant merchantCert error:%s", err)
	}
	return cert
}

// EncryptSensitiveInfo 使用微信支付平台证书加密敏感信息
func (m *WechatPayManager) EncryptSensitiveInfo(ctx context.Context, plaintext string) (string, error) {
	// 确保平台证书已加载
	if m.platformCert == nil {
		m.platformCert = LoadPlatformCert(m)
	}

	// 获取平台证书公钥
	publicKey := m.platformCert.PublicKey.(*rsa.PublicKey)

	// 使用RSA公钥加密
	ciphertext, err := vpayutils.EncryptRSA(plaintext, publicKey)
	if err != nil {
		return "", err
	}

	return ciphertext, nil
}
