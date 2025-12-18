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

package vwxapply4sub

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/vogo/vogo/vlog"
)

const (
	// QuerySettlementURL 查询结算账户URL
	QuerySettlementURL = APIBaseURL + "/v3/apply4sub/sub_merchants/%s/settlement"
)

// SettlementInfoResponse 查询结算账户响应
type SettlementInfoResponse struct {
	AccountType      string `json:"account_type"`                 // 账户类型，ACCOUNT_TYPE_BUSINESS：对公银行账户，ACCOUNT_TYPE_PRIVATE：经营者个人银行卡
	AccountBank      string `json:"account_bank"`                 // 开户银行
	BankName         string `json:"bank_name,omitempty"`          // 开户银行全称（含支行）
	BankBranchID     string `json:"bank_branch_id,omitempty"`     // 开户银行联行号
	AccountNumber    string `json:"account_number"`               // 银行账号，掩码显示
	VerifyResult     string `json:"verify_result"`                // 验证结果，VERIFY_SUCCESS：验证成功，VERIFY_FAIL：验证失败，VERIFYING：验证中
	VerifyFailReason string `json:"verify_fail_reason,omitempty"` // 验证失败原因
}

func (s *SettlementInfoResponse) Error() error {
	if s.VerifyResult == "VERIFY_SUCCESS" {
		return nil
	}

	if s.VerifyResult == "VERIFYING" {
		return errors.New("验证中")
	}

	return errors.New(s.VerifyFailReason)
}

// QuerySettlement 查询结算账户
// subMchID: 特约商户号
// 返回值: 结算账户信息响应, 错误信息
func (c *Apply4SubClient) QuerySettlement(ctx context.Context, subMchID string) (*SettlementInfoResponse, error) {
	// 构建URL
	url := fmt.Sprintf(QuerySettlementURL, subMchID)

	vlog.Infof("query settlement | sub_mch_id: %s", subMchID)

	result, err := c.mgr.Client.Get(ctx, url)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %w", err)
	}

	vlog.Infof("query settlement response | body: %s", respBody)

	var resp SettlementInfoResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}

	return &resp, nil
}
