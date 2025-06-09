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

// CityResponse 查询城市列表响应
type CityResponse struct {
	Data       []*CityInfo `json:"data"`        // 城市列表，查询返回的城市列表结果
	TotalCount int         `json:"total_count"` // 查询总条数，过滤查询到的城市数据总条数
}

// CityInfo 城市信息
type CityInfo struct {
	CityName string `json:"city_name"` // 城市名称，如"北京市"
	CityCode int    `json:"city_code"` // 城市编码，如10
}
