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
	"net/http"
	"strings"
	"time"

	"github.com/vogo/vogo/vencoding/vjson"
	"github.com/vogo/vogo/vlog"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

func (s *JsApiClient) Prepay(ctx context.Context,
	appID, openId string, amount int64,
	outTradeNo, description, attach, callbackUrl string,
	expireTime time.Time,
) (*JsApiPayParams, error) {
	if appID == "" {
		appID = s.mgr.Config.AppID
	}

	prepayRequest := jsapi.PrepayRequest{
		Appid:       core.String(appID),
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
		TimeExpire: core.Time(expireTime),
	}

	reqData, _ := json.Marshal(prepayRequest)

	vlog.Infof("jsapi prepay request | body: %s", reqData)

	resp, result, err := s.jsApi.PrepayWithRequestPayment(ctx, prepayRequest)
	if err != nil {
		errMsg := fmt.Sprintf("%v", err)

		vlog.Errorf("jsapi prepay failed | err: %s", errMsg)

		if strings.Contains(errMsg, "ORDERPAID") &&
			strings.Contains(errMsg, "订单已支付") {
			return nil, ErrOrderPaid
		}

		return nil, err
	}

	vlog.Infof("jsapi prepay response | body: %s", vjson.EnsureMarshal(resp))

	if result.Response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("jsapi prepay failed, status code: %d", result.Response.StatusCode)
	}

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
