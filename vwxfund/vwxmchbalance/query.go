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

package vwxmchbalance

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/vogo/vogo/vlog"
)

// BalanceQueryResponse 账户余额响应
type BalanceQueryResponse struct {
	// AvailableAmount 可用余额(单位:分)
	// 可用余额可用于提现等操作
	AvailableAmount int64 `json:"available_amount"`

	// PendingAmount 不可用余额(单位:分)
	// 不可用余额为冻结金额,不可进行提现等操作
	PendingAmount *int64 `json:"pending_amount,omitempty"`
}

// QueryBalance 查询账户实时余额
// 商户可以通过该接口查询账户的实时余额
// 参考: https://pay.weixin.qq.com/doc/v3/partner/4012720926
func (c *MchBalanceClient) QueryBalance(ctx context.Context, accountType AccountType) (*BalanceQueryResponse, error) {
	vlog.Infof("query balance for account type: %s", accountType)

	// 构建请求URL
	url := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/merchant/fund/balance/%s", accountType)

	// 发送HTTP请求
	result, err := c.mgr.Client.Get(ctx, url)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %w", err)
	}

	vlog.Infof("query balance response: %s", respBody)

	var resp BalanceQueryResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}

	return &resp, nil
}
