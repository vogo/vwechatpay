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
	"encoding/json"
	"fmt"
	"time"

	"github.com/vogo/vogo/vencoding/vjson"
	"github.com/vogo/vogo/vlog"
	"github.com/vogo/vwechatpay"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// PartnerJsApiClient 服务商模式 JSAPI 支付客户端
type PartnerJsApiClient struct {
	mgr *vwechatpay.Manager
}

// NewPartnerJsApiClient 创建服务商模式 JSAPI 支付客户端
func NewPartnerJsApiClient(mgr *vwechatpay.Manager) *PartnerJsApiClient {
	return &PartnerJsApiClient{
		mgr: mgr,
	}
}

// PartnerJsApiPayRequest 服务商模式 JSAPI 支付下单请求
// subMchID: 子商户号
// subOpenID: 用户在子商户 appid 下的唯一标识
// amount: 金额，单位为分
// outTradeNo: 商户订单号
// description: 商品描述
// attach: 附加数据
// callbackUrl: 回调通知地址
// subAppID: 子商户 appid，可选
// profitSharing: 是否分账，可选
func (c *PartnerJsApiClient) PartnerJsApiPayRequest(
	subMchID, subOpenID string,
	amount int64,
	outTradeNo, description, attach, callbackUrl string,
	subAppID string,
	profitSharing bool,
) (*PartnerJsApiPayParams, error) {
	ctx := context.Background()

	// 构建请求参数
	request := map[string]interface{}{
		"sp_appid":     c.mgr.Config.AppID,
		"sp_mchid":     c.mgr.Config.MerchantID,
		"sub_mchid":    subMchID,
		"description":  description,
		"out_trade_no": outTradeNo,
		"notify_url":   callbackUrl,
		"amount": map[string]interface{}{
			"total": amount,
		},
	}

	// 设置子商户 appid（可选）
	if subAppID != "" {
		request["sub_appid"] = subAppID
	}

	// 设置附加数据（可选）
	if attach != "" {
		request["attach"] = attach
	}

	// 设置是否分账（可选）
	if profitSharing {
		request["profit_sharing"] = true
	}

	// 设置支付者信息
	if subAppID != "" {
		// 如果有子商户 appid，使用 sub_openid
		request["payer"] = map[string]interface{}{
			"sub_openid": subOpenID,
		}
	} else {
		// 如果没有子商户 appid，使用服务商 appid 下的 openid
		request["payer"] = map[string]interface{}{
			"openid": subOpenID,
		}
	}

	reqData, _ := json.Marshal(request)
	vlog.Infof("partner jsapi prepay request: %s", reqData)

	// 发送请求
	result, err := c.mgr.Client.Post(ctx, "/v3/pay/partner/transactions/jsapi", request)
	if err != nil {
		vlog.Errorf("partner jsapi prepay failed: %v", err)
		return nil, err
	}

	// 解析响应
	responseMap := make(map[string]interface{})
	if err = vjson.UnmarshalStream(result.Response.Body, &responseMap); err != nil {
		vlog.Errorf("unmarshal response failed: %v", err)
		return nil, err
	}

	vlog.Infof("partner jsapi prepay response: %v", responseMap)

	// 获取预支付交易会话标识
	prepayID, ok := responseMap["prepay_id"].(string)
	if !ok || prepayID == "" {
		return nil, fmt.Errorf("prepay_id is empty")
	}

	// 构建调起支付参数
	timeStamp := fmt.Sprintf("%d", time.Now().Unix())
	nonceStr, err := utils.GenerateNonce()
	if err != nil {
		vlog.Errorf("generate nonce error: %v", err)
		return nil, err
	}
	packageStr := fmt.Sprintf("prepay_id=%s", prepayID)

	// 构建签名参数
	message := fmt.Sprintf("%s\n%s\n%s\n%s\n",
		c.mgr.Config.AppID, timeStamp, nonceStr, packageStr)

	// 计算签名
	signature, err := c.mgr.Sign(message)
	if err != nil {
		vlog.Errorf("sign error: %v", err)
		return nil, err
	}

	// 返回调起支付参数
	return &PartnerJsApiPayParams{
		AppID:     core.String(c.mgr.Config.AppID),
		TimeStamp: core.String(timeStamp),
		NonceStr:  core.String(nonceStr),
		Package:   core.String(packageStr),
		SignType:  core.String("RSA"),
		PaySign:   core.String(signature),
		PayNo:     core.String(outTradeNo),
	}, nil
}

// PartnerJsApiNotifyParse 解析服务商模式 JSAPI 支付回调通知
func (c *PartnerJsApiClient) PartnerJsApiNotifyParse(headerFetcher func(string) string, body []byte) (*notify.Request, map[string]interface{}, error) {
	ctx := context.Background()

	// 验证回调通知签名
	headerArgs, err := getWechatPayHeader(headerFetcher)
	if err != nil {
		return nil, nil, err
	}

	if err = checkWechatPayHeader(ctx, headerArgs); err != nil {
		return nil, nil, err
	}

	message := buildMessage(ctx, headerArgs, body)

	if err = c.mgr.PlatManager.LoadVerifier().Verify(ctx, headerArgs.Serial, message, headerArgs.Signature); err != nil {
		return nil, nil, fmt.Errorf(
			"validate verify fail serial=[%s] request-id=[%s] err=%w",
			headerArgs.Serial, headerArgs.RequestID, err,
		)
	}

	return c.PartnerJsApiNotifyParseBody(body)
}

// PartnerJsApiNotifyParseBody 解析服务商模式 JSAPI 支付回调通知体
func (c *PartnerJsApiClient) PartnerJsApiNotifyParseBody(body []byte) (*notify.Request, map[string]interface{}, error) {
	ret := new(notify.Request)
	if err := json.Unmarshal(body, ret); err != nil {
		return nil, nil, fmt.Errorf("parse request body error: %v", err)
	}

	plaintext, err := utils.DecryptAES256GCM(
		c.mgr.Config.MerchantAPIv3Key, ret.Resource.AssociatedData, ret.Resource.Nonce, ret.Resource.Ciphertext,
	)
	if err != nil {
		return ret, nil, fmt.Errorf("decrypt request error: %v", err)
	}

	ret.Resource.Plaintext = plaintext

	content := map[string]interface{}{}
	if err = json.Unmarshal([]byte(plaintext), &content); err != nil {
		return ret, nil, fmt.Errorf("unmarshal plaintext to content failed: %v", err)
	}

	return ret, content, nil
}
