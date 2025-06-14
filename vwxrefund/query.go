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

package vwxrefund

import (
	"context"
	"fmt"
	"net/http"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

// QueryByOutRefundNo 通过商户退款单号查询单笔退款
// 商户可以通过商户退款单号查询单笔退款，仅能查询自有退款单号的退款单，支持商户使用API、商户平台、微信支付小程序等多种方式发起的退款单查询。
func (c *RefundClient) QueryByOutRefundNo(ctx context.Context, subMchID, outRefundNo string) (*refunddomestic.Refund, error) {
	req := refunddomestic.QueryByOutRefundNoRequest{
		OutRefundNo: core.String(outRefundNo),
	}

	// 服务商模式下需要传递子商户号
	if subMchID != "" {
		req.SubMchid = core.String(subMchID)
	}

	resp, result, err := c.refundApi.QueryByOutRefundNo(ctx, req)
	if err != nil {
		return nil, err
	}

	if result.Response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("query refund failed, status code: %d", result.Response.StatusCode)
	}

	return resp, nil
}
