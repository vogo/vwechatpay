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
	"fmt"

	"github.com/vogo/vogo/vos"
)

// Config 微信支付配置
type Config struct {
	MerchantID           string `json:"merchant_id"`             // 商户号
	MerchantCertSerialNO string `json:"merchant_cert_serial_no"` // 商户证书序列号
	MerchantAPIv3Key     string `json:"merchant_api_v3_key"`     // 商户APIv3密钥
	PrivateKeyPath       string `json:"private_key_path"`        // 私钥文件路径
	PrivateKeyContent    string `json:"private_key_content"`     // 私钥内容
	CertPath             string `json:"cert_path"`               // 证书文件路径
	CertContent          string `json:"cert_content"`            // 证书内容
	AppID                string `json:"app_id"`                  // 应用ID
}

func LoadConfigFromEnv() (*Config, error) {
	cfg := &Config{
		MerchantID:           vos.EnsureEnvString("WECHAT_PAY_MERCHANT_ID"),
		MerchantCertSerialNO: vos.EnsureEnvString("WECHAT_PAY_MERCHANT_CERT_SERIAL_NO"),
		MerchantAPIv3Key:     vos.EnsureEnvString("WECHAT_PAY_MERCHANT_APIV3_KEY"),
		AppID:                vos.EnsureEnvString("WECHAT_PAY_APP_ID"),

		PrivateKeyPath:    vos.EnvString("WECHAT_PAY_PRIVATE_KEY_PATH"),
		PrivateKeyContent: vos.EnvString("WECHAT_PAY_PRIVATE_KEY_CONTENT"),
		CertPath:          vos.EnvString("WECHAT_PAY_CERT_PATH"),
		CertContent:       vos.EnvString("WECHAT_PAY_CERT_CONTENT"),
	}

	if cfg.PrivateKeyContent == "" && cfg.PrivateKeyPath == "" {
		return nil, fmt.Errorf("private key content or path is empty")
	}

	if cfg.CertContent == "" && cfg.CertPath == "" {
		return nil, fmt.Errorf("cert content or path is empty")
	}

	return cfg, nil
}
