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
	"net/http"
	"strings"
	"time"

	"github.com/vogo/vogo/vencoding/vjson"
	"github.com/vogo/vogo/vlog"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
)

// Prepay Native 扫码支付下单，返回 code_url（二维码链接）。
// Native 模式无需 Payer（openId），由商户将 code_url 生成二维码供用户扫码支付。
func (s *NativeClient) Prepay(ctx context.Context,
	appID string, amount int64,
	outTradeNo, description, attach, callbackUrl string,
	expireTime time.Time,
) (*NativePrepayResult, error) {
	if appID == "" {
		appID = s.mgr.Config.AppID
	}

	prepayRequest := native.PrepayRequest{
		Appid:       core.String(appID),
		Mchid:       core.String(s.mgr.Config.MerchantID),
		Description: core.String(description),
		OutTradeNo:  core.String(outTradeNo),
		Attach:      core.String(attach),
		NotifyUrl:   core.String(callbackUrl),
		Amount: &native.Amount{
			Total: core.Int64(amount),
		},
		TimeExpire: core.Time(expireTime),
	}

	reqData, _ := json.Marshal(prepayRequest)

	vlog.Infof("native prepay request | body: %s", reqData)

	resp, result, err := s.native.Prepay(ctx, prepayRequest)
	if err != nil {
		errMsg := fmt.Sprintf("%v", err)

		vlog.Errorf("native prepay failed | err: %s", errMsg)

		if strings.Contains(errMsg, "ORDERPAID") &&
			strings.Contains(errMsg, "订单已支付") {
			return nil, ErrOrderPaid
		}

		return nil, err
	}

	vlog.Infof("native prepay response | body: %s", vjson.EnsureMarshal(resp))

	if result.Response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("native prepay failed, status code: %d", result.Response.StatusCode)
	}

	return &NativePrepayResult{
		CodeURL: resp.CodeUrl,
		PayNo:   &outTradeNo,
	}, nil
}
