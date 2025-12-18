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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/vogo/vogo/vlog"
)

const (
	// ModifySettlementURL 修改结算账户URL
	ModifySettlementURL = APIBaseURL + "/v3/apply4sub/sub_merchants/%s/modify-settlement"

	// QueryModifySettlementURL 查询结算账户修改申请状态URL
	QueryModifySettlementURL = APIBaseURL + "/v3/apply4sub/sub_merchants/%s/application/%s"
)

// ModifySettlementRequest 修改结算账户请求
type ModifySettlementRequest struct {
	AccountType   string `json:"account_type"`             // 账户类型，ACCOUNT_TYPE_BUSINESS：对公银行账户，ACCOUNT_TYPE_PRIVATE：经营者个人银行卡
	AccountBank   string `json:"account_bank"`             // 开户银行，如"工商银行"
	BankName      string `json:"bank_name,omitempty"`      // 开户银行全称（含支行），如"中国工商银行股份有限公司北京市分行营业部"
	BankBranchID  string `json:"bank_branch_id,omitempty"` // 开户银行联行号
	AccountNumber string `json:"account_number"`           // 银行账号，数字，长度遵循系统支持的对公/对私卡号长度标准
	AccountName   string `json:"account_name,omitempty"`   // 开户名称
}

// ModifySettlementResponse 修改结算账户响应
type ModifySettlementResponse struct {
	ApplicationNo string `json:"application_no"` // 修改结算账户申请单号
}

// QueryModifySettlementResponse 查询结算账户修改申请状态响应
type QueryModifySettlementResponse struct {
	AccountName      string `json:"account_name"`                 // 开户名称，掩码显示
	AccountType      string `json:"account_type"`                 // 账户类型，ACCOUNT_TYPE_BUSINESS：对公银行账户，ACCOUNT_TYPE_PRIVATE：经营者个人银行卡
	AccountBank      string `json:"account_bank"`                 // 开户银行全称
	BankName         string `json:"bank_name,omitempty"`          // 开户银行全称（含支行）
	BankBranchID     string `json:"bank_branch_id,omitempty"`     // 开户银行联行号
	AccountNumber    string `json:"account_number"`               // 银行账号，掩码显示
	VerifyResult     string `json:"verify_result"`                // 审核状态，AUDIT_SUCCESS：审核成功，AUDITING：审核中，AUDIT_FAIL：审核驳回
	VerifyFailReason string `json:"verify_fail_reason,omitempty"` // 审核驳回原因，审核成功时为空，审核驳回时为具体原因
	VerifyFinishTime string `json:"verify_finish_time,omitempty"` // 审核结果更新时间，遵循rfc3339标准格式
}

// ModifySettlement 修改结算账户
// subMchID: 特约商户号
// req: 修改结算账户请求
// 返回值: 修改结算账户响应, 错误信息
func (c *Apply4SubClient) ModifySettlement(ctx context.Context, subMchID string, req *ModifySettlementRequest) (*ModifySettlementResponse, error) {
	// 构建URL
	url := fmt.Sprintf(ModifySettlementURL, subMchID)

	encryptAccountNumer, err := c.mgr.PlatManager.Encrypt(req.AccountNumber)
	if err != nil {
		return nil, fmt.Errorf("encrypt account number error: %w", err)
	}
	req.AccountNumber = encryptAccountNumer

	encryptAccountName, err := c.mgr.PlatManager.Encrypt(req.AccountName)
	if err != nil {
		return nil, fmt.Errorf("encrypt account name error: %w", err)
	}
	req.AccountName = encryptAccountName

	// 准备请求体
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request body error: %w", err)
	}

	vlog.Infof("modify settlement | body: %s", reqBody)

	result, err := c.mgr.Client.Post(ctx, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %w", err)
	}

	vlog.Infof("modify settlement response | body: %s", respBody)

	var resp ModifySettlementResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}

	return &resp, nil
}

// QueryModifySettlement 查询结算账户修改申请状态
// subMchID: 特约商户号
// applicationNo: 修改结算账户申请单号
// 返回值: 查询结算账户修改申请状态响应, 错误信息
func (c *Apply4SubClient) QueryModifySettlement(ctx context.Context, subMchID, applicationNo string) (*QueryModifySettlementResponse, error) {
	// 构建URL
	url := fmt.Sprintf(QueryModifySettlementURL, subMchID, applicationNo)

	vlog.Infof("query modify settlement | url: %s", url)

	result, err := c.mgr.Client.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("query modify settlement error: %w", err)
	}

	respBody, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %w", err)
	}

	vlog.Infof("query modify settlement response | body: %s", respBody)

	var resp QueryModifySettlementResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}

	return &resp, nil
}
