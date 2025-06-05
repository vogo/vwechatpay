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

package vjsapipays

import (
	"testing"

	"github.com/vogo/vogo/vos"
	"github.com/vogo/vwechatpay"
)

func TestJsApiNotifyParseBody(t *testing.T) {
	MerchantID := vos.EnvString("WECHAT_PAY_MERCHANT_ID")
	MerchantAPIv3Key := vos.EnvString("WECHAT_PAY_MERCHANT_API_V3_KEY")

	if MerchantID == "" || MerchantAPIv3Key == "" {
		t.Skip("skip test, env not found")
	}

	cli := vwechatpay.NewManagerFromEnv()
	jsApiClient := NewJsApiClient(cli)

	body := `{"id":"a0bc582c-c965-57f8-93cf-aeebc93347a3","create_time":"2024-08-19T22:53:47+08:00","resource_type":"encrypt-resource","event_type":"TRANSACTION.SUCCESS","summary":"支付成功","resource":{"original_type":"transaction","algorithm":"AEAD_AES_256_GCM","ciphertext":"QLVG50QrjlglPOJWjaPL9ITk+IxJYLnABxWSDALJmfy2kByQzrzjdURn03iSU0TZyoZSZjgV5KIijfan6WqshlaKCcOpKMNNFIIXvG/y1KqN1YFQbpydD+6pflHex3nivGnCGNQvbkbJGvIuyln6z6gjmEdo33iqDVxDlWwaRALo+M1/4uLub/5fZETKhrubNOW/hko7P/Y6pTqYf8jiSi82vCuqmr3yPRcsfUZqBHkhZ2hJBbCzMCoLvfz5JehXKtqDse0TLkwAy4dEm+a84YiOkHpXoE5yq++DLEuFPI/JQtfGQTq7BloeKsQiS8/GcYO9HLzXlk5IOFgXBX6mS9owq1+GD8RE+4ELRke2yGpsxnuMpmi94AaNaD/NqTi4fa6WSZ9qpa2c9UVXO82eJ7tQQEjZI53HjyxWJ7oYi+eBjAGjRXqd9VEsyUnNHBCvGYRN7flIfrxSlZX8mgGAnQYd/y6Po6F0L8B4M+5gnFNfCmuNtlWq2iVwwNWLOynAWZWT2u3KZyYz//di3fPBgZoWBFN632F3bnUS0QACfdCRQrJb4Wc1wABXBDAodyT6FI1E3Rl62yWNSGD6NphexZsIXyc3pO5Gudm5EqB+FWX8P1ABgXJ9pAAe6wlK","associated_data":"transaction","nonce":"ET5S6I72TJdA"}}`

	req, content, err := jsApiClient.JsApiNotifyParseBody([]byte(body))
	if err != nil {
		t.Error(err)
	}
	t.Log(req)
	t.Log(content)
}
