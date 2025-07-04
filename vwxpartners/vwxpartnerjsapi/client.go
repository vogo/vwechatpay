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

package vwxpartnerjsapi

import (
	"github.com/vogo/vwechatpay"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/jsapi"
)

// PartnerJsApiClient 服务商模式 JSAPI 支付客户端
type PartnerJsApiClient struct {
	mgr      *vwechatpay.Manager
	jsapiApi *jsapi.JsapiApiService
}

// NewPartnerJsApiClient 创建服务商模式 JSAPI 支付客户端
func NewPartnerJsApiClient(mgr *vwechatpay.Manager) *PartnerJsApiClient {
	return &PartnerJsApiClient{
		mgr: mgr,
		jsapiApi: &jsapi.JsapiApiService{
			Client: mgr.Client,
		},
	}
}
