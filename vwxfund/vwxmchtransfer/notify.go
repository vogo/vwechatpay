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

package vwxmchtransfer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/vogo/vogo/vlog"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// ParseTransferNotify 解析转账回调通知
// 商户需要验证签名，确保回调通知的真实性
func (c *MchTransferClient) ParseTransferNotify(headerFetcher func(string) string, body []byte) (*notify.Request, *TransferNotify, error) {
	ctx := context.Background()

	// 验证回调通知签名
	err := c.mgr.PlatManager.VerifyRequestMessage(ctx, headerFetcher, body)
	if err != nil {
		vlog.Errorf("validate http message failed | err: %v", err)
		return nil, nil, err
	}

	return c.ParseTransferNotifyBody(body)
}

// ParseTransferNotifyBody 解析转账回调通知体
func (c *MchTransferClient) ParseTransferNotifyBody(body []byte) (*notify.Request, *TransferNotify, error) {
	ret := new(notify.Request)
	if err := json.Unmarshal(body, ret); err != nil {
		return nil, nil, fmt.Errorf("parse request body error: %v", err)
	}

	// 解密通知内容
	plaintext, err := utils.DecryptAES256GCM(
		c.mgr.Config.MerchantAPIv3Key, ret.Resource.AssociatedData, ret.Resource.Nonce, ret.Resource.Ciphertext,
	)
	if err != nil {
		return ret, nil, fmt.Errorf("decrypt request error: %v", err)
	}

	ret.Resource.Plaintext = plaintext

	vlog.Infof("received transfer notify | plaintext: %s", plaintext)

	// 解析转账通知内容
	var transferNotify TransferNotify
	if err := json.Unmarshal([]byte(plaintext), &transferNotify); err != nil {
		return ret, nil, fmt.Errorf("unmarshal transfer notify error: %v", err)
	}

	return ret, &transferNotify, nil
}

// TransferNotify 转账回调通知数据结构
type TransferNotify struct {
	MchId          string `json:"mch_id"`           // 商户号
	OutBillNo      string `json:"out_bill_no"`      // 商户单号
	State          string `json:"state"`            // 转账状态
	FailReason     string `json:"fail_reason"`      // 失败原因
	Openid         string `json:"openid"`           // 用户openid
	UserName       string `json:"user_name"`        // 收款用户姓名
	TransferBillNo string `json:"transfer_bill_no"` // 微信转账单号
	TransferAmount int64  `json:"transfer_amount"`  // 转账金额
	TransferRemark string `json:"transfer_remark"`  // 转账备注
	TransferTime   string `json:"transfer_time"`    // 转账时间
	CreateTime     string `json:"create_time"`      // 创建时间
	UpdateTime     string `json:"update_time"`      // 更新时间
}
