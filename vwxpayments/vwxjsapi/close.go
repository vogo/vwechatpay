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
	"fmt"

	"github.com/vogo/vogo/vlog"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

// CloseOrder 关闭订单
// 以下情况需要调用关单接口：
// 1. 商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；
// 2. 系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。
// outTradeNo: 商户订单号
func (s *JsApiClient) CloseOrder(ctx context.Context, outTradeNo string) error {
	// 构建请求参数
	req := jsapi.CloseOrderRequest{
		Mchid:      core.String(s.mgr.Config.MerchantID),
		OutTradeNo: core.String(outTradeNo),
	}

	vlog.Infof("jsapi close order, outTradeNo: %s", outTradeNo)

	// 发送请求
	result, err := s.jsApi.CloseOrder(ctx, req)
	if err != nil {
		vlog.Errorf("close order error: %v", err)
		return err
	}

	vlog.Infof("jsapi close order response status: %d", result.Response.StatusCode)

	// 关单成功返回204状态码
	if result.Response.StatusCode != 204 {
		return fmt.Errorf("close order failed with status code: %d", result.Response.StatusCode)
	}

	return nil
}
