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

	"github.com/vogo/vogo/vencoding/vjson"
	"github.com/vogo/vogo/vlog"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

// QueryOrderById 根据微信支付订单号查询订单
// transactionId: 微信支付订单号
func (s *JsApiClient) QueryOrderById(ctx context.Context, transactionId string) (*payments.Transaction, error) {
	// 构建请求参数
	req := jsapi.QueryOrderByIdRequest{
		TransactionId: core.String(transactionId),
		Mchid:         core.String(s.mgr.Config.MerchantID),
	}

	vlog.Infof("jsapi query order | transaction_id: %s", transactionId)

	// 发送请求
	resp, result, err := s.jsApi.QueryOrderById(ctx, req)
	if err != nil {
		vlog.Errorf("query order by id error | err: %v", err)
		return nil, err
	}

	vlog.Infof("jsapi query order response | body: %s", vjson.EnsureMarshal(resp))

	if result.Response.StatusCode != 200 {
		return nil, fmt.Errorf("query order by id failed with status code: %d", result.Response.StatusCode)
	}

	return resp, nil
}

// QueryOrderByOutTradeNo 根据商户订单号查询订单
// outTradeNo: 商户订单号
func (s *JsApiClient) QueryOrderByOutTradeNo(ctx context.Context, outTradeNo string) (*payments.Transaction, error) {
	// 构建请求参数
	req := jsapi.QueryOrderByOutTradeNoRequest{
		OutTradeNo: core.String(outTradeNo),
		Mchid:      core.String(s.mgr.Config.MerchantID),
	}

	vlog.Infof("jsapi query request | out_trade_no: %s", outTradeNo)

	// 发送请求
	resp, result, err := s.jsApi.QueryOrderByOutTradeNo(ctx, req)
	if err != nil {
		vlog.Errorf("query order by out trade no error | err: %v", err)
		return nil, err
	}

	vlog.Infof("jsapi query order response | body: %s", vjson.EnsureMarshal(resp))

	if result.Response.StatusCode != 200 {
		return nil, fmt.Errorf("query order by out trade no failed with status code: %d", result.Response.StatusCode)
	}

	return resp, nil
}
