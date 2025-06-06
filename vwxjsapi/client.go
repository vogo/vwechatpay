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

package vwxjsapi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/vogo/vogo/vlog"
	"github.com/vogo/vwechatpay"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type JsApiClient struct {
	mgr   *vwechatpay.Manager
	JsApi *jsapi.JsapiApiService
}

func NewJsApiClient(mgr *vwechatpay.Manager) *JsApiClient {
	return &JsApiClient{
		mgr:   mgr,
		JsApi: &jsapi.JsapiApiService{Client: mgr.Client},
	}
}

func (s *JsApiClient) JsApiPayRequest(openId string, amount int64, outTradeNo, description, attach, callbackUrl string) (*JsApiPayParams, error) {
	ctx := context.Background()

	prepayRequest := jsapi.PrepayRequest{
		Appid:       core.String(s.mgr.Config.AppID),
		Mchid:       core.String(s.mgr.Config.MerchantID),
		Description: core.String(description),
		OutTradeNo:  core.String(outTradeNo),
		Attach:      core.String(attach),
		NotifyUrl:   core.String(callbackUrl),
		Amount: &jsapi.Amount{
			Total: core.Int64(amount),
		},
		Payer: &jsapi.Payer{
			Openid: core.String(openId),
		},
	}

	reqData, _ := json.Marshal(prepayRequest)
	vlog.Infof("jsapi prepay request: %s", reqData)
	resp, _, err := s.JsApi.PrepayWithRequestPayment(ctx, prepayRequest)

	if err != nil {
		vlog.Errorf("jsapi prepay failed: %v", err)
		return nil, err
	}

	respData, _ := json.Marshal(resp)
	vlog.Infof("jsapi prepay response: %s", respData)

	return &JsApiPayParams{
		AppID:     resp.Appid,
		TimeStamp: resp.TimeStamp,
		NonceStr:  resp.NonceStr,
		Package:   resp.Package,
		SignType:  resp.SignType,
		PaySign:   resp.PaySign,
		PayNo:     &outTradeNo,
	}, nil
}

func (s *JsApiClient) JsApiNotifyParse(headerFetcher func(string) string, body []byte) (*notify.Request, map[string]interface{}, error) {
	ctx := context.Background()

	err := s.ValidateHTTPMessage(ctx, headerFetcher, body)
	if err != nil {
		vlog.Errorf("validate http message failed: %v", err)
		return nil, nil, err
	}

	return s.JsApiNotifyParseBody(body)
}

func (s *JsApiClient) JsApiNotifyParseBody(body []byte) (*notify.Request, map[string]interface{}, error) {
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
