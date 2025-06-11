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

package vwxapply4sub

// ApplymentStatusRequest 查询申请单状态请求
type ApplymentStatusRequest struct {
	BusinessCode string `json:"business_code"` // 业务申请编号
}

// ApplymentStatusByIDRequest 通过申请单号查询申请单状态请求
type ApplymentStatusByIDRequest struct {
	ApplymentID int64 `json:"applyment_id"` // 微信支付申请单号
}

// ApplymentStatusResponse 查询申请单状态响应
type ApplymentStatusResponse struct {
	BusinessCode      string         `json:"business_code"`       // 业务申请编号
	ApplymentID       int64          `json:"applyment_id"`        // 微信支付申请单号
	SubMchID          string         `json:"sub_mchid"`           // 特约商户号
	SignURL           string         `json:"sign_url"`            // 超级管理员签约链接
	ApplymentState    string         `json:"applyment_state"`     // 申请单状态
	ApplymentStateMsg string         `json:"applyment_state_msg"` // 申请单状态描述
	AuditDetail       []*AuditDetail `json:"audit_detail"`        // 驳回原因详情
}

// AuditDetail 驳回原因详情
type AuditDetail struct {
	Field        string `json:"field"`         // 字段名
	FieldName    string `json:"field_name"`    // 字段中文名
	RejectReason string `json:"reject_reason"` // 驳回原因
}

var applymentStateMap = map[string]string{
	"APPLYMENT_STATE_EDITTING":        "编辑中",
	"APPLYMENT_STATE_AUDITING":        "审核中",
	"APPLYMENT_STATE_REJECTED":        "已驳回",
	"APPLYMENT_STATE_TO_BE_CONFIRMED": "待账户验证",
	"APPLYMENT_STATE_TO_BE_SIGNED":    "待签约",
	"APPLYMENT_STATE_FINISHED":        "已完成",
	"APPLYMENT_STATE_CANCELED":        "已作废",
	"APPLYMENT_STATE_SIGNING":         "开通权限中",
}

var applymentStateDetailMap = map[string]string{
	"APPLYMENT_STATE_EDITTING":        "提交申请发生错误导致，请尝试重新提交",
	"APPLYMENT_STATE_AUDITING":        "申请单正在审核中，可让超级管理员用微信打开“签约链接”，完成绑定微信号后，申请单进度将通过微信公众号通知超级管理员，引导完成后续步骤",
	"APPLYMENT_STATE_REJECTED":        "请按照驳回原因修改申请资料，可让超级管理员用微信打开“签约链接”，完成绑定微信号，后续申请单进度将通过微信公众号通知超级管理员",
	"APPLYMENT_STATE_TO_BE_CONFIRMED": "请超级管理员使用微信打开返回的“签约链接”，根据页面指引完成账户验证",
	"APPLYMENT_STATE_TO_BE_SIGNED":    "请超级管理员使用微信打开返回的“签约链接”，根据页面指引完成签约",
	"APPLYMENT_STATE_FINISHED":        "商户入驻申请已完成，可发起交易",
	"APPLYMENT_STATE_CANCELED":        "申请单已被撤销",
	"APPLYMENT_STATE_SIGNING":         "系统开通相关权限中，请耐心等待",
}

func (a *ApplymentStatusResponse) StateDesc() string {
	desc, ok := applymentStateMap[a.ApplymentState]
	if ok {
		return desc
	}
	return a.ApplymentStateMsg
}

func (a *ApplymentStatusResponse) StateDetail() string {
	detail, ok := applymentStateDetailMap[a.ApplymentState]
	if ok {
		return detail
	}
	return a.ApplymentStateMsg
}
