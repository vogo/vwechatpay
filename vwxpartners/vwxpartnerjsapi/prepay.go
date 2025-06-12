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
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// Prepay 服务商模式 JSAPI 支付下单请求
// subMchID: 子商户号
// userOpenID: 用户唯一标识
// amount: 金额，单位为分
// outTradeNo: 商户订单号
// description: 商品描述
// attach: 附加数据
// callbackUrl: 回调通知地址
// subAppID: 子商户 appid，可选
// profitSharing: 是否分账，可选
func (c *PartnerJsApiClient) Prepay(ctx context.Context,
	subMchID, userOpenID string,
	amount int64,
	outTradeNo, description, attach, callbackUrl string,
	subAppID string,
	profitSharing bool,
) (*PartnerJsApiPayParams, error) {
	req := jsapi.PrepayRequest{
		SpAppid:     core.String(c.mgr.Config.AppID),
		SpMchid:     core.String(c.mgr.Config.MerchantID),
		SubMchid:    core.String(subMchID),
		Description: core.String(description),
		OutTradeNo:  core.String(outTradeNo),
		NotifyUrl:   core.String(callbackUrl),
		Amount: &jsapi.Amount{
			Total: core.Int64(amount),
		},
	}

	// 设置子商户 appid（可选）
	if subAppID != "" {
		req.SubAppid = core.String(subAppID)
	}

	// 设置附加数据（可选）
	if attach != "" {
		req.Attach = core.String(attach)
	}

	// 设置是否分账（可选）
	if profitSharing {
		req.SettleInfo = &jsapi.SettleInfo{
			ProfitSharing: core.Bool(true),
		}
	}

	// 设置支付者信息
	if subAppID != "" {
		// 如果有子商户 appid，使用 sub_openid
		req.Payer = &jsapi.Payer{
			SubOpenid: core.String(userOpenID),
		}
	} else {
		// 如果没有子商户 appid，使用服务商 appid 下的 openid
		req.Payer = &jsapi.Payer{
			SpOpenid: core.String(userOpenID),
		}
	}

	reqData, _ := json.Marshal(req)
	vlog.Infof("partner jsapi prepay request: %s", reqData)

	resp, result, err := c.jsapiApiService.Prepay(ctx, req)
	if err != nil {
		vlog.Errorf("partner jsapi prepay error: %v", err)
		return nil, err
	}

	vlog.Infof("partner jsapi prepay response: %s", vjson.EnsureMarshal(resp))

	if result.Response.StatusCode != 200 {
		return nil, fmt.Errorf("prepay failed with status code: %d", result.Response.StatusCode)
	}

	if resp.PrepayId == nil || *resp.PrepayId == "" {
		return nil, fmt.Errorf("prepay_id is empty")
	}

	// 获取预支付交易会话标识
	prepayID := *resp.PrepayId

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
