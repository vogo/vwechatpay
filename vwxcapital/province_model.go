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

// ProvinceResponse 查询省份列表响应
type ProvinceResponse struct {
	Data       []*ProvinceInfo `json:"data"`        // 省份列表，查询到的省份列表数组
	TotalCount int             `json:"total_count"` // 查询总条数，查询到的省份数据总条数
}

// ProvinceInfo 省份信息
type ProvinceInfo struct {
	ProvinceName string `json:"province_name"` // 省份名称，如"广东省"
	ProvinceCode int    `json:"province_code"` // 省份编码，如22
}
