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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/vogo/vogo/vlog"
)

// CancelTransferRequest 撤销转账请求参数
type CancelTransferRequest struct {
	OutBillNo string `json:"out_bill_no"`
}

// CancelTransferResponse 撤销转账响应参数
type CancelTransferResponse struct {
	OutBillNo      string `json:"out_bill_no"`      // 商户单号
	TransferBillNo string `json:"transfer_bill_no"` // 微信转账单号
	State          string `json:"state"`            // 转账状态
	UpdateTime     string `json:"update_time"`      // 更新时间
}

// CancelTransfer 撤销转账
// 商户可以通过该接口撤销转账，仅能撤销处于"WAIT_PAY"状态的转账单
// 参考: https://pay.weixin.qq.com/doc/v3/merchant/4012716458
func (c *MchTransferClient) CancelTransfer(ctx context.Context, outBillNo string) (*CancelTransferResponse, error) {
	// 构建请求参数
	req := CancelTransferRequest{
		OutBillNo: outBillNo,
	}

	vlog.Infof("cancel transfer | out_bill_no: %s", outBillNo)

	// 序列化请求体
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request error: %v", err)
	}

	vlog.Infof("cancel transfer request | body: %s", reqBody)

	// 构建请求URL
	url := fmt.Sprintf("https://api.mch.weixin.qq.com/v3/fund-app/mch-transfer/transfer-bills/out-bill-no/%s/cancel", outBillNo)

	// 发送HTTP请求
	result, err := c.mgr.Client.Post(ctx, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("cancel transfer error: %w", err)
	}

	respBody, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %w", err)
	}

	vlog.Infof("cancel transfer response | body: %s", respBody)

	var resp CancelTransferResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}
	return &resp, nil
}
