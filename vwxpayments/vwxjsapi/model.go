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

package vwxjsapi

import (
	"errors"
)

var (
	ErrOrderPaid = errors.New("订单已支付")
)

type JsApiPayParams struct {
	AppID     *string `json:"appId"`
	TimeStamp *string `json:"timeStamp"`
	NonceStr  *string `json:"nonceStr"`
	Package   *string `json:"package"`
	SignType  *string `json:"signType"`
	PaySign   *string `json:"paySign"`
	PayNo     *string `json:"payNo"`
}

type wechatPayHeader struct {
	RequestID string
	Serial    string
	Signature string
	Nonce     string
	Timestamp int64
}
