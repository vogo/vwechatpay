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

package vwxpartner

// 小微商户进件相关的数据模型

// ApplymentRequest 提交申请单请求
type ApplymentRequest struct {
	BusinessCode    string          `json:"business_code"`           // 业务申请编号
	ContactInfo     ContactInfo     `json:"contact_info"`            // 超级管理员信息
	SubjectInfo     SubjectInfo     `json:"subject_info"`            // 主体资料
	BusinessInfo    BusinessInfo    `json:"business_info"`           // 经营资料
	SettlementInfo  SettlementInfo  `json:"settlement_info"`         // 结算规则
	BankAccountInfo BankAccountInfo `json:"bank_account_info"`       // 结算银行账户
	AdditionInfo    *AdditionInfo   `json:"addition_info,omitempty"` // 补充材料
}

// ApplymentResponse 提交申请单响应
type ApplymentResponse struct {
	ApplymentID int64 `json:"applyment_id"` // 微信支付申请单号
}

// ContactInfo 超级管理员信息
type ContactInfo struct {
	ContactName     string `json:"contact_name"`      // 超级管理员姓名
	ContactIDNumber string `json:"contact_id_number"` // 超级管理员身份证件号码
	OpenID          string `json:"openid"`            // 超级管理员微信openid
	MobilePhone     string `json:"mobile_phone"`      // 联系手机
	ContactEmail    string `json:"contact_email"`     // 联系邮箱
}

// SubjectInfo 主体资料
type SubjectInfo struct {
	SubjectType  string       `json:"subject_type"`   // 主体类型，小微商户固定为 SUBJECT_TYPE_MICRO
	MicroBizInfo MicroBizInfo `json:"micro_biz_info"` // 小微商户经营者/法人身份证件
}

// MicroBizInfo 小微商户经营者/法人身份证件
type MicroBizInfo struct {
	MicroBizType    string           `json:"micro_biz_type"`              // 小微商户类型，MICRO_TYPE_STORE：门店场所，MICRO_TYPE_MOBILE：流动经营/便民服务，MICRO_TYPE_ONLINE：线上商品/服务交易
	MicroStoreInfo  *MicroStoreInfo  `json:"micro_store_info,omitempty"`  // 门店场所信息
	MicroMobileInfo *MicroMobileInfo `json:"micro_mobile_info,omitempty"` // 流动经营/便民服务信息
	MicroOnlineInfo *MicroOnlineInfo `json:"micro_online_info,omitempty"` // 线上商品/服务交易信息
}

// MicroStoreInfo 门店场所信息
type MicroStoreInfo struct {
	MicroName        string `json:"micro_name"`         // 门店名称
	MicroAddressCode string `json:"micro_address_code"` // 门店省市编码
	MicroAddress     string `json:"micro_address"`      // 门店地址
	StoreEntrancePic string `json:"store_entrance_pic"` // 门店门头照片
	MicroIndoorCopy  string `json:"micro_indoor_copy"`  // 店内环境照片
	StoreStreetPic   string `json:"store_street_pic"`   // 门店街景照片
}

// MicroMobileInfo 流动经营/便民服务信息
type MicroMobileInfo struct {
	MicroName        string `json:"micro_name"`         // 经营/服务名称
	MicroAddressCode string `json:"micro_address_code"` // 经营/服务所在地省市编码
	MicroAddress     string `json:"micro_address"`      // 经营/服务所在地地址
	MicroMobilePics  string `json:"micro_mobile_pics"`  // 经营/服务现场照片
}

// MicroOnlineInfo 线上商品/服务交易信息
type MicroOnlineInfo struct {
	MicroName        string `json:"micro_name"`         // 线上店铺名称
	MicroAddressCode string `json:"micro_address_code"` // 线上店铺经营地址省市编码
	MicroAddress     string `json:"micro_address"`      // 线上店铺经营地址
	MicroSite        string `json:"micro_site"`         // 线上店铺/应用截图
}

// BusinessInfo 经营资料
type BusinessInfo struct {
	MerchantShortname string `json:"merchant_shortname"` // 商户简称
	ServicePhone      string `json:"service_phone"`      // 客服电话
	SalesInfo         string `json:"sales_info"`         // 经营场景
}

// SettlementInfo 结算规则
type SettlementInfo struct {
	SettlementID        string `json:"settlement_id"`         // 入驻结算规则ID
	QualificationType   string `json:"qualification_type"`    // 所属行业
	Qualifications      string `json:"qualifications"`        // 特殊资质
	BusinessAdditionPic string `json:"business_addition_pic"` // 补充材料
}

// BankAccountInfo 结算银行账户
type BankAccountInfo struct {
	BankAccountType string `json:"bank_account_type"` // 账户类型，小微商户固定为 BANK_ACCOUNT_TYPE_PERSONAL
	AccountName     string `json:"account_name"`      // 开户名称
	AccountBank     string `json:"account_bank"`      // 开户银行
	BankAddressCode string `json:"bank_address_code"` // 开户银行省市编码
	BankBranchID    string `json:"bank_branch_id"`    // 开户银行联行号
	BankName        string `json:"bank_name"`         // 开户银行全称（含支行）
	AccountNumber   string `json:"account_number"`    // 银行账号
}

// AdditionInfo 补充材料
type AdditionInfo struct {
	LegalPersonCommitment string `json:"legal_person_commitment"` // 法人开户承诺函
	LegalPersonVideo      string `json:"legal_person_video"`      // 法人开户意愿视频
	BusinessAdditionPics  string `json:"business_addition_pics"`  // 补充材料
}

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
	BusinessCode      string `json:"business_code"`       // 业务申请编号
	ApplymentID       int64  `json:"applyment_id"`        // 微信支付申请单号
	SubMchID          string `json:"sub_mchid"`           // 特约商户号
	SignURL           string `json:"sign_url"`            // 超级管理员签约链接
	ApplymentState    string `json:"applyment_state"`     // 申请单状态
	ApplymentStateMsg string `json:"applyment_state_msg"` // 申请单状态描述
	AuditDetail       string `json:"audit_detail"`        // 驳回原因详情
}
