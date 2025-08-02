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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/vogo/vogo/vlog"
)

// TransferSceneReportInfo 转账场景报备信息, 参考 https://pay.weixin.qq.com/doc/v3/merchant/4013774588
type TransferSceneReportInfo struct {
	InfoType    string `json:"info_type,omitempty"`    // 【信息类型】 不能超过15个字符，商户所属转账场景下的信息类型，此字段内容为固定值, 信息类型，两条明细中必须分别填写以下两个取值：活动名称,奖励说明
	InfoContent string `json:"info_content,omitempty"` // 【信息内容】 不能超过32个字符，商户所属转账场景下的信息内容，商户可按实际业务场景自定义传参
}

// TransferRequest 发起转账请求
type TransferRequest struct {
	Appid                    string                     `json:"appid"`                          // 商户AppID
	OutBillNo                string                     `json:"out_bill_no"`                    // 商户单号
	TransferSceneId          string                     `json:"transfer_scene_id"`              // 转账场景ID, 可前往“商户平台-产品中心-商家转账”中申请。如：1000（现金营销），1006（企业报销）等
	Openid                   string                     `json:"openid"`                         // 收款用户OpenID
	UserName                 string                     `json:"user_name,omitempty"`            // 收款用户姓名
	TransferAmount           int64                      `json:"transfer_amount"`                // 转账金额
	TransferRemark           string                     `json:"transfer_remark"`                // 转账备注
	NotifyUrl                string                     `json:"notify_url,omitempty"`           // 通知地址
	UserRecvPerception       string                     `json:"user_recv_perception,omitempty"` // 用户收款感知, 参考 https://pay.weixin.qq.com/doc/v3/merchant/4012711988#2.3-%E5%8F%91%E8%B5%B7%E8%BD%AC%E8%B4%A6
	TransferSceneReportInfos []*TransferSceneReportInfo `json:"transfer_scene_report_infos"`    // 转账场景报备信息
}

// TransferBillsResponse 转账单信息
type TransferBillsResponse struct {
	OutBillNo      string `json:"out_bill_no"`      // 商户单号
	TransferBillNo string `json:"transfer_bill_no"` // 微信转账单号
	CreateTime     string `json:"create_time"`      // 创建时间
	State          string `json:"state"`            // 转账状态
	FailReason     string `json:"fail_reason"`      // 失败原因
	PackageInfo    string `json:"package_info"`     // 转账场景包信息
}

// Transfer 发起转账
// outBillNo: 商户单号
// transferSceneId: 转账场景ID
// openid: 收款用户OpenID
// transferAmount: 转账金额（单位：分）
// transferRemark: 转账备注
// userName: 收款用户姓名（可选，转账金额>=2000元时必填）
// notifyUrl: 通知地址（可选）
// userRecvPerception: 用户收款感知（可选）
// transferSceneReportInfos: 转账场景报备信息
func (c *MchTransferClient) Transfer(ctx context.Context,
	outBillNo, transferSceneId, openid string,
	transferAmount int64,
	transferRemark string,
	userName, notifyUrl, userRecvPerception string,
	transferSceneReportInfos []*TransferSceneReportInfo,
) (*TransferBillsResponse, error) {
	// 构建请求参数
	req := TransferRequest{
		Appid:                    c.mgr.Config.AppID,
		OutBillNo:                outBillNo,
		TransferSceneId:          transferSceneId,
		Openid:                   openid,
		TransferAmount:           transferAmount,
		TransferRemark:           transferRemark,
		TransferSceneReportInfos: transferSceneReportInfos,
	}

	// 设置可选参数
	if userName != "" {
		req.UserName = userName
	}
	if notifyUrl != "" {
		req.NotifyUrl = notifyUrl
	}
	if userRecvPerception != "" {
		req.UserRecvPerception = userRecvPerception
	}

	return c.DoTransfer(ctx, &req)
}

func (c *MchTransferClient) DoTransfer(ctx context.Context, req *TransferRequest) (*TransferBillsResponse, error) {
	// 序列化请求数据
	reqBody, err := json.Marshal(req)
	if err != nil {
		vlog.Errorf("marshal transfer request error: %v", err)
		return nil, err
	}

	vlog.Infof("mch transfer request: %s", reqBody)

	// 发送HTTP请求
	url := "https://api.mch.weixin.qq.com/v3/fund-app/mch-transfer/transfer-bills"
	result, err := c.mgr.Client.Post(ctx, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("transfer error: %w", err)
	}

	respBody, err := io.ReadAll(result.Response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %w", err)
	}

	vlog.Infof("transfer response: %s", respBody)

	var resp TransferBillsResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response error: %w", err)
	}
	return &resp, nil
}
