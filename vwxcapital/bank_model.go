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

// PersonalBankRequest 查询支持个人业务的银行列表请求
type PersonalBankRequest struct {
	Offset int `json:"offset,omitempty"` // 本次查询偏移量，非负整数，表示该次请求资源的起始位置，从0开始计数
	Limit  int `json:"limit"`            // 本次请求最大查询条数，非0非负的整数，该次请求可返回的最大资源条数
}

// PersonalBankResponse 查询支持个人业务的银行列表响应
type PersonalBankResponse struct {
	TotalCount int          `json:"total_count"` // 查询数据总条数，银行列表数据的总条数
	Count      int          `json:"count"`       // 本次查询数据条数，本次查询银行列表返回的数据条数
	Data       []*BankInfo  `json:"data"`        // 银行列表，本次查询到的银行列表数据
	Offset     int          `json:"offset"`      // 本次查询偏移量
	Links      PaginateLink `json:"links"`       // 分页链接
}

// BankInfo 银行信息
type BankInfo struct {
	BankAlias       string `json:"bank_alias"`        // 银行别名，如"招商银行"
	BankAliasCode   string `json:"bank_alias_code"`   // 银行别名编码，如"1000009561"
	AccountBank     string `json:"account_bank"`      // 开户银行，如"招商银行"
	AccountBankCode int    `json:"account_bank_code"` // 开户银行编码，如1001
	NeedBankBranch  bool   `json:"need_bank_branch"`  // 是否需要填写支行
}

// PaginateLink 分页链接
type PaginateLink struct {
	Next string `json:"next"` // 下一页链接
	Prev string `json:"prev"` // 上一页链接
	Self string `json:"self"` // 当前页链接
}

// BranchBankRequest 查询支行列表请求
type BranchBankRequest struct {
	BankAliasCode string `json:"bank_alias_code"`  // 银行别名编码，如"1000009561"
	CityCode      int    `json:"city_code"`        // 城市编码，唯一标识一座城市，用于结合银行别名编码查询支行列表
	Offset        int    `json:"offset,omitempty"` // 本次查询偏移量，非负整数，表示该次请求资源的起始位置，从0开始计数
	Limit         int    `json:"limit"`            // 本次请求最大查询条数，非0非负的整数，该次请求可返回的最大资源条数
}

// BranchBankResponse 查询支行列表响应
type BranchBankResponse struct {
	TotalCount      int           `json:"total_count"`       // 查询数据总条数，经过条件筛选，查询到的支行总数
	Count           int           `json:"count"`             // 本次查询条数，本次查询到的支行数据条数
	Data            []*BranchInfo `json:"data"`              // 支行列表，单次查询返回的支行列表结果数组
	Offset          int           `json:"offset"`            // 本次查询偏移量
	Links           PaginateLink  `json:"links"`             // 分页链接
	AccountBank     string        `json:"account_bank"`      // 开户银行，开户银行名称
	AccountBankCode int           `json:"account_bank_code"` // 开户银行编码
	BankAlias       string        `json:"bank_alias"`        // 银行别名，查询到的支行所属银行的银行别名
	BankAliasCode   string        `json:"bank_alias_code"`   // 银行别名编码，查询到的支行所属银行的银行别名编码
}

// BranchInfo 支行信息
type BranchInfo struct {
	BankBranchName string `json:"bank_branch_name"` // 支行名称，如"招商银行股份有限公司北京分行"
	BankBranchID   string `json:"bank_branch_id"`   // 支行联行号，如"308100005019"
}
