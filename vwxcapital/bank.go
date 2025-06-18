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
	"strings"

	"github.com/vogo/vogo/vlog"
)

const (
	// PersonalBankingURL 查询支持个人业务的银行列表URL
	PersonalBankingURL = APIBaseURL + "/v3/capital/capitallhh/banks/personal-banking"
	// BranchBankingURL 查询支行列表URL，需要替换 {bank_alias_code}
	BranchBankingURL = APIBaseURL + "/v3/capital/capitallhh/banks/%s/branches"
)

// QueryPersonalBanks 查询支持个人业务的银行列表
// req: 查询请求，包含偏移量和查询条数
// 返回值: 银行列表响应, 错误信息
func (c *CapitalClient) QueryPersonalBanks(ctx context.Context, req *PersonalBankRequest) (*PersonalBankResponse, error) {
	// 构建URL和查询参数
	url := PersonalBankingURL

	// 添加查询参数到URL
	queryParams := make([]string, 0, 2)
	if req.Offset > 0 {
		queryParams = append(queryParams, fmt.Sprintf("offset=%d", req.Offset))
	}
	queryParams = append(queryParams, fmt.Sprintf("limit=%d", req.Limit))

	// 将查询参数添加到URL
	if len(queryParams) > 0 {
		url = url + "?" + strings.Join(queryParams, "&")
	}

	vlog.Debugf("query personal banks: %s", url)

	result, err := c.mgr.Client.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("query personal banks error: %w", err)
	}

	respBody, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %w", err)
	}

	vlog.Debugf("query personal banks response: %s", respBody)

	var resp PersonalBankResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}

	return &resp, nil
}

// QueryBranchBanks 查询支行列表
// req: 查询请求，包含银行别名编码、城市编码、偏移量和查询条数
// 返回值: 支行列表响应, 错误信息
func (c *CapitalClient) QueryBranchBanks(ctx context.Context, req *BranchBankRequest) (*BranchBankResponse, error) {
	// 构建URL和查询参数
	url := fmt.Sprintf(BranchBankingURL, req.BankAliasCode)

	// 添加查询参数到URL
	queryParams := make([]string, 0, 3)
	queryParams = append(queryParams, fmt.Sprintf("city_code=%d", req.CityCode))
	if req.Offset > 0 {
		queryParams = append(queryParams, fmt.Sprintf("offset=%d", req.Offset))
	}
	queryParams = append(queryParams, fmt.Sprintf("limit=%d", req.Limit))

	// 将查询参数添加到URL
	if len(queryParams) > 0 {
		url = url + "?" + strings.Join(queryParams, "&")
	}

	vlog.Infof("query branch banks: %s", url)

	result, err := c.mgr.Client.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("query branch banks error: %w", err)
	}

	respBody, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %w", err)
	}

	vlog.Infof("query branch banks response: %s", respBody)

	var resp BranchBankResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}

	return &resp, nil
}

func (c *CapitalClient) UpdateBankCache(ctx context.Context) error {
	cache := make(map[string]*BankInfo)

	vlog.Infof("update bank cache start")

	offset := 0
	limit := 100
	for {
		req := &PersonalBankRequest{
			Offset: offset,
			Limit:  limit,
		}
		resp, err := c.QueryPersonalBanks(ctx, req)
		if err != nil {
			return fmt.Errorf("query personal banks error: %w", err)
		}
		for _, bank := range resp.Data {
			cache[bank.BankAlias] = bank
		}

		if resp.Count < limit {
			break
		}

		offset = offset + limit
	}

	c.bankCache = cache

	vlog.Infof("update bank cache end, count: %d", len(cache))

	return nil
}

func (c *CapitalClient) GetBankInfo(bankAlias string) []*BankInfo {
	if c.bankCache == nil {
		if err := c.UpdateBankCache(context.Background()); err != nil {
			vlog.Errorf("update bank cache error: %v", err)
			return nil
		}
	}

	if bank, ok := c.bankCache[bankAlias]; ok {
		return []*BankInfo{bank}
	}

	var banks []*BankInfo
	for _, bank := range c.bankCache {
		if strings.Contains(bank.BankAlias, bankAlias) {
			banks = append(banks, bank)
		}
	}
	return banks
}
