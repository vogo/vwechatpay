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

package wxpartnerpays

import (
	"github.com/vogo/vwechatpay"
)

const (
	// APIBaseURL 微信支付API基础URL
	APIBaseURL = "https://api.mch.weixin.qq.com"

	// ApplymentURL 提交申请单URL
	ApplymentURL = APIBaseURL + "/v3/applyment4sub/applyment/"

	// ApplymentQueryByBusinessCodeURL 通过业务申请编号查询申请单状态URL
	ApplymentQueryByBusinessCodeURL = APIBaseURL + "/v3/applyment4sub/applyment/business_code/%s"

	// ApplymentQueryByApplymentIDURL 通过申请单号查询申请单状态URL
	ApplymentQueryByApplymentIDURL = APIBaseURL + "/v3/applyment4sub/applyment/applyment_id/%d"
)

// PartnerPayClient 微信支付客户端
type PartnerPayClient struct {
	manager *vwechatpay.WechatPayManager
}

// NewPartnerPayClient 创建微信支付客户端
func NewPartnerPayClient(manager *vwechatpay.WechatPayManager) *PartnerPayClient {
	return &PartnerPayClient{
		manager: manager,
	}
}
