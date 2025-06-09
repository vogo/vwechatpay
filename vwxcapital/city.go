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
	// CitiesURL 查询城市列表URL，需要替换 {province_code}
	CitiesURL = APIBaseURL + "/v3/capital/capitallhh/areas/provinces/%d/cities"
)

// QueryCities 查询城市列表
// provinceCode: 省份编码，唯一标识一个省份
// 返回值: 城市列表响应, 错误信息
func (c *CapitalClient) QueryCities(ctx context.Context, provinceCode int) (*CityResponse, error) {
	// 检查缓存
	if c.cityCache != nil {
		if cityResp, ok := c.cityCache[provinceCode]; ok {
			vlog.Infof("using city cache for province code: %d", provinceCode)
			return cityResp, nil
		}
	}

	// 构建URL
	url := fmt.Sprintf(CitiesURL, provinceCode)

	vlog.Infof("query cities for province code %d: %s", provinceCode, url)

	result, err := c.mgr.Client.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("query cities error: %w", err)
	}

	respBody, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %w", err)
	}

	vlog.Infof("query cities response: %s", respBody)

	var resp CityResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}

	// 更新缓存
	if c.cityCache == nil {
		c.cityCache = make(map[int]*CityResponse)
	}
	c.cityCache[provinceCode] = &resp

	return &resp, nil
}

// ClearCityCache 清除城市列表缓存
// provinceCode: 省份编码，如果为0则清除所有省份的城市缓存
func (c *CapitalClient) ClearCityCache(provinceCode int) {
	if provinceCode == 0 {
		c.cityCache = nil
		vlog.Info("all city cache cleared")
	} else if c.cityCache != nil {
		delete(c.cityCache, provinceCode)
		vlog.Infof("city cache for province code %d cleared", provinceCode)
	}
}
