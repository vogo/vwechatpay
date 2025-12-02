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
	// ApplymentURL 提交申请单URL
	ApplymentURL = APIBaseURL + "/v3/applyment4sub/applyment/"
)

// SubmitApplyment 提交申请单
func (c *Apply4SubClient) SubmitApplyment(ctx context.Context, req *ApplymentRequest) (*ApplymentResponse, error) {
	// 准备请求体
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request body error: %w", err)
	}

	vlog.Infof("submit applyment: %s", reqBody)

	result, err := c.mgr.Client.Post(ctx, ApplymentURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %w", err)
	}

	vlog.Infof("submit applyment response: %s", respBody)

	var resp ApplymentResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}

	return &resp, nil
}
