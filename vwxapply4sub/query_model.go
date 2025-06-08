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
	BusinessCode      string        `json:"business_code"`       // 业务申请编号
	ApplymentID       int64         `json:"applyment_id"`        // 微信支付申请单号
	SubMchID          string        `json:"sub_mchid"`           // 特约商户号
	SignURL           string        `json:"sign_url"`            // 超级管理员签约链接
	ApplymentState    string        `json:"applyment_state"`     // 申请单状态
	ApplymentStateMsg string        `json:"applyment_state_msg"` // 申请单状态描述
	AuditDetail       []AuditDetail `json:"audit_detail"`        // 驳回原因详情
}

// AuditDetail 驳回原因详情
type AuditDetail struct {
	Field        string `json:"field"`         // 字段名
	FieldName    string `json:"field_name"`    // 字段中文名
	RejectReason string `json:"reject_reason"` // 驳回原因
}
