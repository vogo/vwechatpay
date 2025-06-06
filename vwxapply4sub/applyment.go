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

	"github.com/vogo/vogo/vencoding/vjson"
)

// SubmitApplyment 提交申请单
func (c *PartnerClient) SubmitApplyment(ctx context.Context, req *ApplymentRequest) (*ApplymentResponse, error) {
	// 准备请求体
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request body error: %w", err)
	}

	// 发送请求
	result, err := c.mgr.Client.Post(ctx, ApplymentURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("submit applyment error: %w", err)
	}

	// 解析响应
	var resp ApplymentResponse
	if err := vjson.UnmarshalStream(result.Response.Body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}

	return &resp, nil
}

// QueryApplymentByBusinessCode 通过业务申请编号查询申请单状态
func (c *PartnerClient) QueryApplymentByBusinessCode(ctx context.Context, businessCode string) (*ApplymentStatusResponse, error) {
	// 构建URL
	url := fmt.Sprintf(ApplymentQueryByBusinessCodeURL, businessCode)

	// 发送请求
	result, err := c.mgr.Client.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("query applyment by business code error: %w", err)
	}

	// 解析响应
	var resp ApplymentStatusResponse
	if err := vjson.UnmarshalStream(result.Response.Body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}

	return &resp, nil
}

// QueryApplymentByApplymentID 通过申请单号查询申请单状态
func (c *PartnerClient) QueryApplymentByApplymentID(ctx context.Context, applymentID int64) (*ApplymentStatusResponse, error) {
	// 构建URL
	url := fmt.Sprintf(ApplymentQueryByApplymentIDURL, applymentID)

	// 发送请求
	result, err := c.mgr.Client.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("query applyment by applyment id error: %w", err)
	}

	// 解析响应
	var resp ApplymentStatusResponse
	if err := vjson.UnmarshalStream(result.Response.Body, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}

	return &resp, nil
}
