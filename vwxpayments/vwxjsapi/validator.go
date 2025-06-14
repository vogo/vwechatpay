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
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
)

func (s *JsApiClient) ValidateHTTPMessage(ctx context.Context, headerFetcher func(string) string, body []byte) error {
	headerArgs, err := getWechatPayHeader(headerFetcher)
	if err != nil {
		return err
	}

	if err = checkWechatPayHeader(ctx, headerArgs); err != nil {
		return err
	}

	message := buildMessage(ctx, headerArgs, body)

	if err = s.mgr.PlatManager.LoadVerifier().Verify(ctx, headerArgs.Serial, message, headerArgs.Signature); err != nil {
		return fmt.Errorf(
			"validate verify fail serial=[%s] request-id=[%s] err=%w",
			headerArgs.Serial, headerArgs.RequestID, err,
		)
	}

	return nil
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
