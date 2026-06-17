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

package vwxnative

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/vogo/vogo/vlog"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// NativeNotifyParse 解析 Native 支付回调通知。
// 包含 HTTP 头验签 + AES-GCM 解密，返回 notify.Request 与解密后的明文内容。
func (s *NativeClient) NativeNotifyParse(headerFetcher func(string) string, body []byte) (*notify.Request, map[string]interface{}, error) {
	ctx := context.Background()

	err := s.ValidateHTTPMessage(ctx, headerFetcher, body)
	if err != nil {
		vlog.Errorf("validate http message failed | err: %v", err)
		return nil, nil, err
	}

	return s.NativeNotifyParseBody(body)
}

// NativeNotifyParseBody 解析回调通知报文体（不含验签），完成 AES-GCM 解密并返回明文内容。
func (s *NativeClient) NativeNotifyParseBody(body []byte) (*notify.Request, map[string]interface{}, error) {
	ret := new(notify.Request)
	if err := json.Unmarshal(body, ret); err != nil {
		return nil, nil, fmt.Errorf("parse request body error: %v", err)
	}

	plaintext, err := utils.DecryptAES256GCM(
		s.mgr.Config.MerchantAPIv3Key, ret.Resource.AssociatedData, ret.Resource.Nonce, ret.Resource.Ciphertext,
	)
	if err != nil {
		return ret, nil, fmt.Errorf("decrypt request error: %v", err)
	}

	ret.Resource.Plaintext = plaintext

	content := map[string]interface{}{}
	if err = json.Unmarshal([]byte(plaintext), &content); err != nil {
		return ret, nil, fmt.Errorf("unmarshal plaintext to content failed: %v", err)
	}

	return ret, content, nil
}
