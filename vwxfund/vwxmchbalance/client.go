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

package vwxmchbalance

import (
	"github.com/vogo/vwechatpay"
)

// MchBalanceClient 商户账户余额查询客户端
type MchBalanceClient struct {
	mgr *vwechatpay.Manager
}

// NewMchBalanceClient 创建商户账户余额查询客户端
func NewMchBalanceClient(mgr *vwechatpay.Manager) *MchBalanceClient {
	return &MchBalanceClient{
		mgr: mgr,
	}
}
