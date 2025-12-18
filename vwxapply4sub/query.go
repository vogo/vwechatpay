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
	"fmt"
	"io"

	"github.com/vogo/vogo/vlog"
)

const (
	// ApplymentQueryByBusinessCodeURL 通过业务申请编号查询申请单状态URL
	ApplymentQueryByBusinessCodeURL = APIBaseURL + "/v3/applyment4sub/applyment/business_code/%s"

	// ApplymentQueryByApplymentIDURL 通过申请单号查询申请单状态URL
	ApplymentQueryByApplymentIDURL = APIBaseURL + "/v3/applyment4sub/applyment/applyment_id/%d"
)

// QueryApplymentByBusinessCode 通过业务申请编号查询申请单状态
// businessCode: 业务申请编号
// 返回值: 申请单状态响应, 错误信息
func (c *Apply4SubClient) QueryApplymentByBusinessCode(ctx context.Context, businessCode string) (*ApplymentStatusResponse, error) {
	// 构建URL
	url := fmt.Sprintf(ApplymentQueryByBusinessCodeURL, businessCode)

	vlog.Infof("query applyment by business code | business_code: %s", businessCode)

	result, err := c.mgr.Client.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("query applyment by business code error: %w", err)
	}

	respBody, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %w", err)
	}

	vlog.Infof("query applyment by business code response | body: %s", respBody)

	var resp ApplymentStatusResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}

	return &resp, nil
}

// QueryApplymentByApplymentID 通过申请单号查询申请单状态
// applymentID: 微信支付申请单号
// 返回值: 申请单状态响应, 错误信息
func (c *Apply4SubClient) QueryApplymentByApplymentID(ctx context.Context, applymentID int64) (*ApplymentStatusResponse, error) {
	// 构建URL
	url := fmt.Sprintf(ApplymentQueryByApplymentIDURL, applymentID)

	vlog.Infof("query applyment by applyment id | applyment_id: %d", applymentID)

	result, err := c.mgr.Client.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("query applyment by applyment id error: %w", err)
	}

	respBody, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %w", err)
	}

	vlog.Infof("query applyment by applyment id response | body: %s", respBody)

	var resp ApplymentStatusResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}

	return &resp, nil
}
