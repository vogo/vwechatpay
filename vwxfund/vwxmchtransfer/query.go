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

package vwxmchtransfer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/vogo/vogo/vlog"
)

// QueryTransferResponse 查询转账单响应参数
type QueryTransferResponse struct {
	Mchid          string `json:"mchid"`            // 商户号
	OutBillNo      string `json:"out_bill_no"`      // 商户单号
	TransferBillNo string `json:"transfer_bill_no"` // 微信转账单号
	Appid          string `json:"appid"`            // 应用ID
	State          string `json:"state"`            // 转账状态
	TransferAmount int64  `json:"transfer_amount"`  // 转账金额
	TransferRemark string `json:"transfer_remark"`  // 转账备注
	FailReason     string `json:"fail_reason"`      // 失败原因
	Openid         string `json:"openid"`           // 收款用户OpenID
	UserName       string `json:"user_name"`        // 收款用户姓名
	CreateTime     string `json:"create_time"`      // 创建时间
	UpdateTime     string `json:"update_time"`      // 更新时间
}

// QueryTransferByOutBillNo 商户单号查询转账单
// 商户可以通过该接口查询转账单据的详细信息
// 参考: https://pay.weixin.qq.com/doc/v3/merchant/4012716437
func (c *MchTransferClient) QueryTransferByOutBillNo(ctx context.Context, outBillNo string) (*QueryTransferResponse, error) {
	vlog.Infof("query transfer by outBillNo: %s", outBillNo)

	// 构建请求URL
	url := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/fund-app/mch-transfer/transfer-bills/out-bill-no/%s", outBillNo)

	// 发送HTTP请求
	result, err := c.mgr.Client.Get(ctx, url)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %w", err)
	}

	vlog.Infof("query transfer response: %s", respBody)

	var resp QueryTransferResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}
	return &resp, nil
}

// QueryTransferByTransferBillNo 微信单号查询转账单
// 商户可以通过该接口查询转账单据的详细信息
func (c *MchTransferClient) QueryTransferByTransferBillNo(ctx context.Context, transferBillNo string) (*QueryTransferResponse, error) {
	vlog.Infof("query transfer by transferBillNo: %s", transferBillNo)

	// 构建请求URL
	url := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/fund-app/mch-transfer/transfer-bills/transfer-bill-no/%s", transferBillNo)

	// 发送HTTP请求
	result, err := c.mgr.Client.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("query transfer by transferBillNo error: %w", err)
	}

	respBody, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %w", err)
	}

	vlog.Infof("query transfer response: %s", respBody)

	var resp QueryTransferResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}
	return &resp, nil
}
