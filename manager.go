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
	"fmt"

	"github.com/vogo/vogo/vlog"
	"github.com/vogo/vogo/vos"
	"github.com/vogo/vogo/vsync/vrun"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type PlatManager interface {
	LoadCert() *x509.Certificate
	LoadVerifier() *verifiers.SHA256WithRSAVerifier
	Encrypt(plaintext string) (string, error)
}

var PlatManagerInit func(mgr *Manager) PlatManager

// Manager 微信支付管理类,包含微信支付客户端和商户信息.
type Manager struct {
	runner             *vrun.Runner
	Config             *Config
	merchantPrivateKey *rsa.PrivateKey
	merchantCert       *x509.Certificate
	PlatManager        PlatManager
	Client             *core.Client
}

func NewManager(cfg *Config) (*Manager, error) {
	mgr := &Manager{
		runner: vrun.New(),
		Config: cfg,
	}

	privateKey, err := utils.LoadPrivateKeyWithPath(cfg.PrivateKeyPath)
	if err != nil {
		vlog.Errorf("load merchant private key error: %v", err)
		return nil, err
	}
	mgr.merchantPrivateKey = privateKey

	cert, err := utils.LoadCertificateWithPath(cfg.CertPath)
	if err != nil {
		vlog.Errorf("load merchant cert error: %v", err)
		return nil, err
	}
	mgr.merchantCert = cert

	cli, err := buildWechatPayClient(cfg, mgr.merchantPrivateKey)
	if err != nil {
		vlog.Errorf("build wechat pay client error: %v", err)
		return nil, err
	}
	mgr.Client = cli

	mgr.PlatManager = PlatManagerInit(mgr)

	return mgr, nil
}

func NewManagerFromEnv() (_mgr *Manager, _err error) {
	defer func() {
		if err := recover(); err != nil {
			if e, ok := err.(error); ok {
				_err = e
			} else {
				_err = fmt.Errorf("new manager from env error: %v", err)
			}
		}
	}()

	cfg := &Config{
		MerchantID:           vos.EnsureEnvString("WECHAT_PAY_MERCHANT_ID"),
		MerchantCertSerialNO: vos.EnsureEnvString("WECHAT_PAY_MERCHANT_CERT_SERIAL_NO"),
		MerchantAPIv3Key:     vos.EnsureEnvString("WECHAT_PAY_MERCHANT_API_V3_KEY"),
		PrivateKeyPath:       vos.EnsureEnvString("WECHAT_PAY_PRIVATE_KEY_PATH"),
		CertPath:             vos.EnsureEnvString("WECHAT_PAY_CERT_PATH"),
		AppID:                vos.EnsureEnvString("WECHAT_PAY_APP_ID"),
	}

	return NewManager(cfg)
}

func buildWechatPayClient(cfg *Config, key *rsa.PrivateKey) (*core.Client, error) {
	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(cfg.MerchantID,
			cfg.MerchantCertSerialNO,
			key,
			cfg.MerchantAPIv3Key),
	}

	return core.NewClient(ctx, opts...)
}
