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

package vwxcapital

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/vogo/vogo/vlog"
)

const (
	// ProvincesURL 查询省份列表URL
	ProvincesURL = APIBaseURL + "/v3/capital/capitallhh/areas/provinces"
)

// QueryProvinces 查询省份列表
// 返回值: 省份列表响应, 错误信息
func (c *CapitalClient) QueryProvinces(ctx context.Context) (*ProvinceResponse, error) {
	// 检查缓存
	if c.provinceCache != nil {
		return c.provinceCache, nil
	}

	// 构建URL
	url := ProvincesURL

	vlog.Infof("query provinces: %s", url)

	result, err := c.mgr.Client.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("query provinces error: %w", err)
	}

	respBody, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %w", err)
	}

	vlog.Infof("query provinces response: %s", respBody)

	var resp ProvinceResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}

	// 更新缓存
	c.provinceCache = &resp

	return &resp, nil
}
