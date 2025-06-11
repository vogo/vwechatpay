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

// 小微商户进件相关的数据模型

// 主体类型常量
const (
	SubjectTypeMicro        = "SUBJECT_TYPE_MICRO"        // （小微）无营业执照、免办理工商注册登记的实体商户
	SubjectTypeIndividual   = "SUBJECT_TYPE_INDIVIDUAL"   // （个体户）营业执照上的主体类型一般为个体户、个体工商户、个体经营
	SubjectTypeEnterprise   = "SUBJECT_TYPE_ENTERPRISE"   // （企业）营业执照上的主体类型一般为有限公司、有限责任公司
	SubjectTypeGovernment   = "SUBJECT_TYPE_GOVERNMENT"   // （政府机关）包括各级、各类政府机关，如机关党委、税务、民政、人社、工商、商务、市监等
	SubjectTypeInstitutions = "SUBJECT_TYPE_INSTITUTIONS" // （事业单位）包括国内各类事业单位，如：医疗、教育、学校等单位
	SubjectTypeOthers       = "SUBJECT_TYPE_OTHERS"       // （社会组织）包括社会团体、民办非企业、基金会、基层群众性自治组织、农村集体经济组织等组织
)

// ApplymentRequest 提交申请单请求
type ApplymentRequest struct {
	BusinessCode    string           `json:"business_code"`           // 业务申请编号
	ContactInfo     *ContactInfo     `json:"contact_info"`            // 超级管理员信息
	SubjectInfo     *SubjectInfo     `json:"subject_info"`            // 主体资料
	BusinessInfo    *BusinessInfo    `json:"business_info"`           // 经营资料
	SettlementInfo  *SettlementInfo  `json:"settlement_info"`         // 结算规则
	BankAccountInfo *BankAccountInfo `json:"bank_account_info"`       // 结算银行账户
	AdditionInfo    *AdditionInfo    `json:"addition_info,omitempty"` // 补充材料
}

// ApplymentResponse 提交申请单响应
type ApplymentResponse struct {
	ApplymentID int64 `json:"applyment_id"` // 微信支付申请单号
}

// ContactInfo 超级管理员信息
type ContactInfo struct {
	ContactType                 string `json:"contact_type"`                            // 超级管理员类型，LEGAL：法人，SUPER：经办人
	ContactName                 string `json:"contact_name"`                            // 超级管理员姓名
	ContactIDDocType            string `json:"contact_id_doc_type,omitempty"`           // 超级管理员证件类型
	ContactIDNumber             string `json:"contact_id_number"`                       // 超级管理员身份证件号码
	ContactIDDocCopy            string `json:"contact_id_doc_copy,omitempty"`           // 超级管理员证件正面照片
	ContactIDDocCopyBack        string `json:"contact_id_doc_copy_back,omitempty"`      // 超级管理员证件反面照片
	ContactPeriodBegin          string `json:"contact_period_begin,omitempty"`          // 超级管理员证件有效期开始时间
	ContactPeriodEnd            string `json:"contact_period_end,omitempty"`            // 超级管理员证件有效期结束时间
	BusinessAuthorizationLetter string `json:"business_authorization_letter,omitempty"` // 业务办理授权函
	OpenID                      string `json:"openid"`                                  // 超级管理员微信openid
	MobilePhone                 string `json:"mobile_phone"`                            // 联系手机
	ContactEmail                string `json:"contact_email"`                           // 联系邮箱
}

// SubjectInfo 主体资料
type SubjectInfo struct {
	SubjectType            string                  `json:"subject_type"`                       // 主体类型，小微商户固定为 SubjectTypeMicro
	FinanceInstitution     bool                    `json:"finance_institution,omitempty"`      // 是否金融机构
	BusinessLicenseInfo    *BusinessLicenseInfo    `json:"business_license_info,omitempty"`    // 营业执照信息
	CertificateInfo        *CertificateInfo        `json:"certificate_info,omitempty"`         // 登记证书信息
	CertificateLetterCopy  string                  `json:"certificate_letter_copy,omitempty"`  // 单位证明函照片
	FinanceInstitutionInfo *FinanceInstitutionInfo `json:"finance_institution_info,omitempty"` // 金融机构信息
	IdentityInfo           *IdentityInfo           `json:"identity_info,omitempty"`            // 经营者/法人身份证件
	UboInfoList            []*UboInfo              `json:"ubo_info_list,omitempty"`            // 最终受益人信息列表
	MicroBizInfo           *MicroBizInfo           `json:"micro_biz_info"`                     // 小微商户经营者/法人身份证件
}

// UboInfo 最终受益人信息
type UboInfo struct {
	UboIDDocType     string `json:"ubo_id_doc_type"`      // 证件类型
	UboIDDocCopy     string `json:"ubo_id_doc_copy"`      // 证件正面照片
	UboIDDocCopyBack string `json:"ubo_id_doc_copy_back"` // 证件反面照片
	UboIDDocName     string `json:"ubo_id_doc_name"`      // 受益人姓名
	UboIDDocNumber   string `json:"ubo_id_doc_number"`    // 证件号码
	UboIDDocAddress  string `json:"ubo_id_doc_address"`   // 证件地址
	UboPeriodBegin   string `json:"ubo_period_begin"`     // 证件有效期开始时间
	UboPeriodEnd     string `json:"ubo_period_end"`       // 证件有效期结束时间
}

// BusinessLicenseInfo 营业执照信息
type BusinessLicenseInfo struct {
	LicenseCopy    string `json:"license_copy"`    // 营业执照照片
	LicenseNumber  string `json:"license_number"`  // 营业执照注册号
	MerchantName   string `json:"merchant_name"`   // 商户名称
	LegalPerson    string `json:"legal_person"`    // 法人姓名
	LicenseAddress string `json:"license_address"` // 注册地址
	PeriodBegin    string `json:"period_begin"`    // 有效期开始日期
	PeriodEnd      string `json:"period_end"`      // 有效期结束日期
}

// CertificateInfo 登记证书信息
type CertificateInfo struct {
	CertCopy       string `json:"cert_copy"`       // 登记证书照片
	CertType       string `json:"cert_type"`       // 证书类型
	CertNumber     string `json:"cert_number"`     // 证书编号
	MerchantName   string `json:"merchant_name"`   // 商户名称
	CompanyAddress string `json:"company_address"` // 注册地址
	LegalPerson    string `json:"legal_person"`    // 法人姓名
	PeriodBegin    string `json:"period_begin"`    // 有效期开始日期
	PeriodEnd      string `json:"period_end"`      // 有效期结束日期
}

// FinanceInstitutionInfo 金融机构信息
type FinanceInstitutionInfo struct {
	FinanceType        string   `json:"finance_type"`         // 金融机构类型
	FinanceLicensePics []string `json:"finance_license_pics"` // 金融机构许可证图片
}

// IdentityInfo 经营者/法人身份证件
type IdentityInfo struct {
	IDHolderType        string      `json:"id_holder_type"`                  // 证件持有人类型, LEGAL: 经营者/法人 SUPER: 经办人
	IDDocType           string      `json:"id_doc_type"`                     // 证件类型, IDENTIFICATION_TYPE_IDCARD: 中国大陆居民-身份证
	AuthorizeLetterCopy string      `json:"authorize_letter_copy,omitempty"` // 授权函照片
	IDCardInfo          *IDCardInfo `json:"id_card_info,omitempty"`          // 身份证信息
	IDDocInfo           *IDDocInfo  `json:"id_doc_info,omitempty"`           // 其他类型证件信息
	Owner               bool        `json:"owner,omitempty"`                 // 是否为受益人
}

// IDCardInfo 身份证信息
type IDCardInfo struct {
	IDCardCopy      string `json:"id_card_copy"`      // 身份证人像面照片
	IDCardNational  string `json:"id_card_national"`  // 身份证国徽面照片
	IDCardName      string `json:"id_card_name"`      // 身份证姓名
	IDCardNumber    string `json:"id_card_number"`    // 身份证号码
	IDCardAddress   string `json:"id_card_address"`   // 身份证地址
	CardPeriodBegin string `json:"card_period_begin"` // 身份证有效期开始时间
	CardPeriodEnd   string `json:"card_period_end"`   // 身份证有效期结束时间
}

// IDDocInfo 其他类型证件信息
type IDDocInfo struct {
	IDDocCopy      string `json:"id_doc_copy"`      // 证件照片
	IDDocCopyBack  string `json:"id_doc_copy_back"` // 证件反面照片
	IDDocName      string `json:"id_doc_name"`      // 证件姓名
	IDDocNumber    string `json:"id_doc_number"`    // 证件号码
	IDDocAddress   string `json:"id_doc_address"`   // 证件地址
	DocPeriodBegin string `json:"doc_period_begin"` // 证件有效期开始时间
	DocPeriodEnd   string `json:"doc_period_end"`   // 证件有效期结束时间
}

// MicroBizInfo 小微商户经营者/法人身份证件
// 小微商户类型常量
const (
	MicroTypeStore  = "MICRO_TYPE_STORE"  // 门店场所
	MicroTypeMobile = "MICRO_TYPE_MOBILE" // 流动经营/便民服务
	MicroTypeOnline = "MICRO_TYPE_ONLINE" // 线上商品/服务交易
)

type MicroBizInfo struct {
	MicroBizType    string           `json:"micro_biz_type"`              // 小微商户类型，MicroTypeStore：门店场所，MicroTypeMobile：流动经营/便民服务，MicroTypeOnline：线上商品/服务交易
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
	StoreLongitude   string `json:"store_longitude"`    // 门店经度
	StoreLatitude    string `json:"store_latitude"`     // 门店纬度
}

// MicroMobileInfo 流动经营/便民服务信息
type MicroMobileInfo struct {
	MicroMobileName    string `json:"micro_mobile_name"`    // 经营/服务名称
	MicroMobileCity    string `json:"micro_mobile_city"`    // 经营/服务所在地城市编码
	MicroMobileAddress string `json:"micro_mobile_address"` // 经营/服务所在地详细地址
	MicroMobilePics    string `json:"micro_mobile_pics"`    // 经营/服务现场照片
}

// MicroOnlineInfo 线上商品/服务交易信息
type MicroOnlineInfo struct {
	MicroOnlineStore string `json:"micro_online_store"` // 线上店铺名称
	MicroEcName      string `json:"micro_ec_name"`      // 电商平台名称
	MicroQrcode      string `json:"micro_qrcode"`       // 线上店铺二维码
	MicroLink        string `json:"micro_link"`         // 线上店铺链接
}

// BusinessInfo 经营资料
type BusinessInfo struct {
	MerchantShortname string     `json:"merchant_shortname"` // 商户简称
	ServicePhone      string     `json:"service_phone"`      // 客服电话
	SalesInfo         *SalesInfo `json:"sales_info"`         // 经营场景
}

// SalesInfo 经营场景
type SalesInfo struct {
	SalesScenesType []string         `json:"sales_scenes_type"`           // 经营场景类型
	BizStoreInfo    *BizStoreInfo    `json:"biz_store_info,omitempty"`    // 线下门店场景
	MpInfo          *MpInfo          `json:"mp_info,omitempty"`           // 公众号场景
	MiniProgramInfo *MiniProgramInfo `json:"mini_program_info,omitempty"` // 小程序场景
	AppInfo         *AppInfo         `json:"app_info,omitempty"`          // APP场景
	WebInfo         *WebInfo         `json:"web_info,omitempty"`          // Web场景
	WeworkInfo      *WeworkInfo      `json:"wework_info,omitempty"`       // 企业微信场景
}

// BizStoreInfo 线下门店场景
type BizStoreInfo struct {
	BizStoreName     string   `json:"biz_store_name"`     // 门店名称
	BizAddressCode   string   `json:"biz_address_code"`   // 门店省市编码
	BizStoreAddress  string   `json:"biz_store_address"`  // 门店地址
	StoreEntrancePic []string `json:"store_entrance_pic"` // 门店门头照片
	IndoorPic        []string `json:"indoor_pic"`         // 店内环境照片
	BizSubAppid      string   `json:"biz_sub_appid"`      // 线下场所对应的商家APPID
}

// MpInfo 公众号场景
type MpInfo struct {
	MpAppid    string   `json:"mp_appid"`     // 公众号APPID
	MpSubAppid string   `json:"mp_sub_appid"` // 公众号页面截图
	MpPics     []string `json:"mp_pics"`      // 公众号页面截图
}

// MiniProgramInfo 小程序场景
type MiniProgramInfo struct {
	MiniProgramAppid    string   `json:"mini_program_appid"`     // 小程序APPID
	MiniProgramSubAppid string   `json:"mini_program_sub_appid"` // 小程序APPID
	MiniProgramPics     []string `json:"mini_program_pics"`      // 小程序页面截图
}

// AppInfo APP场景
type AppInfo struct {
	AppAppid    string   `json:"app_appid"`     // APP APPID
	AppSubAppid string   `json:"app_sub_appid"` // APP APPID
	AppPics     []string `json:"app_pics"`      // APP截图
}

// WebInfo Web场景
type WebInfo struct {
	Domain           string `json:"domain"`            // 互联网网站域名
	WebAuthorisation string `json:"web_authorisation"` // 网站授权函
	WebAppid         string `json:"web_appid"`         // 网站对应的商家APPID
}

// WeworkInfo 企业微信场景
type WeworkInfo struct {
	CorpId     string   `json:"corp_id"`     // 企业微信CorpID
	WeworkPics []string `json:"wework_pics"` // 企业微信页面截图
}

// SettlementInfo 结算规则
type SettlementInfo struct {
	SettlementID         string   `json:"settlement_id"`                    // 入驻结算规则ID
	QualificationType    string   `json:"qualification_type"`               // 所属行业
	Qualifications       []string `json:"qualifications,omitempty"`         // 特殊资质
	BusinessAdditionPic  string   `json:"business_addition_pic,omitempty"`  // 补充材料
	ActivitiesID         string   `json:"activities_id,omitempty"`          // 优惠费率活动ID
	ActivitiesRate       string   `json:"activities_rate,omitempty"`        // 优惠费率活动值
	ActivitiesAdditions  []string `json:"activities_additions,omitempty"`   // 优惠费率活动补充材料
	DebitActivitiesRate  string   `json:"debit_activities_rate,omitempty"`  // 优惠费率活动值（借记卡）
	CreditActivitiesRate string   `json:"credit_activities_rate,omitempty"` // 优惠费率活动值（贷记卡）
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
	LegalPersonCommitment string   `json:"legal_person_commitment"`          // 法人开户承诺函
	LegalPersonVideo      string   `json:"legal_person_video"`               // 法人开户意愿视频
	BusinessAdditionPics  []string `json:"business_addition_pics,omitempty"` // 补充材料
	BusinessAdditionMsg   string   `json:"business_addition_msg,omitempty"`  // 补充说明
}
