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
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
)

// PartnerJsApiPayParams 服务商模式 JSAPI 支付参数
type PartnerJsApiPayParams struct {
	// AppID 服务商应用ID
	AppID *string `json:"appId"`

	// TimeStamp 时间戳
	TimeStamp *string `json:"timeStamp"`

	// NonceStr 随机字符串
	NonceStr *string `json:"nonceStr"`

	// Package 订单详情扩展字符串
	Package *string `json:"package"`

	// SignType 签名方式
	SignType *string `json:"signType"`

	// PaySign 签名
	PaySign *string `json:"paySign"`

	// PayNo 商户订单号
	PayNo *string `json:"payNo"`
}

// wechatPayHeader 微信支付头部信息
type wechatPayHeader struct {
	Timestamp int64
	Nonce     string
	Serial    string
	Signature string
	RequestID string
}

// getWechatPayHeader 从 http.Header 中获取 wechatPayHeader 信息
func getWechatPayHeader(headerFetcher func(string) string) (wechatPayHeader, error) {
	requestID := strings.TrimSpace(headerFetcher(consts.RequestID))

	getHeaderString := func(key string) (string, error) {
		val := strings.TrimSpace(headerFetcher(key))
		if val == "" {
			return "", fmt.Errorf("key `%s` is empty in header, request-id=[%s]", key, requestID)
		}
		return val, nil
	}

	getHeaderInt64 := func(key string) (int64, error) {
		val, err := getHeaderString(key)
		if err != nil {
			return 0, nil
		}
		ret, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid `%s` in header, request-id=[%s], err:%w", key, requestID, err)
		}
		return ret, nil
	}

	ret := wechatPayHeader{
		RequestID: requestID,
	}
	var err error

	if ret.Serial, err = getHeaderString(consts.WechatPaySerial); err != nil {
		return ret, err
	}

	if ret.Signature, err = getHeaderString(consts.WechatPaySignature); err != nil {
		return ret, err
	}

	if ret.Timestamp, err = getHeaderInt64(consts.WechatPayTimestamp); err != nil {
		return ret, err
	}

	if ret.Nonce, err = getHeaderString(consts.WechatPayNonce); err != nil {
		return ret, err
	}

	return ret, nil
}

// checkWechatPayHeader 对 wechatPayHeader 内容进行检查，看是否符合要求
//
// 检查项：
//   - Timestamp 与当前时间之差不得超过 FiveMinute;
func checkWechatPayHeader(ctx context.Context, args wechatPayHeader) error {
	// Suppressing warnings
	_ = ctx

	if math.Abs(float64(time.Now().Unix()-args.Timestamp)) >= consts.FiveMinute {
		return fmt.Errorf("timestamp=[%d] expires, request-id=[%s]", args.Timestamp, args.RequestID)
	}
	return nil
}

// buildMessage 根据微信支付签名格式构造验签原文
func buildMessage(ctx context.Context, headerArgs wechatPayHeader, body []byte) string {
	// Suppressing warnings
	_ = ctx

	return fmt.Sprintf("%d\n%s\n%s\n", headerArgs.Timestamp, headerArgs.Nonce, string(body))
}
