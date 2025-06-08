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

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/vogo/vwechatpay"
)

var microApplymentRequestDemo = &ApplymentRequest{
	BusinessCode: "1900000001_10000", // 业务申请编号，服务商自定义
	ContactInfo: ContactInfo{
		ContactName:     "张三",                 // 联系人姓名
		ContactIDNumber: "110101199003070073", // 联系人身份证号
		MobilePhone:     "13900000000",        // 手机号
		ContactEmail:    "EMAIL",              // 邮箱
	},
	SubjectInfo: SubjectInfo{
		SubjectType: SubjectTypeMicro, // 小微商户固定为 SubjectTypeMicro
		MicroBizInfo: MicroBizInfo{
			MicroBizType: MicroTypeOnline, // 门店场所
			MicroOnlineInfo: &MicroOnlineInfo{
				MicroOnlineStore: "张三的小店",              // 线上店铺名称
				MicroEcName:      "XX购物平台",             // 电商平台名称
				MicroLink:        "URL_ADDRESS.qq.com", // 线上店铺链接
			},
		},
		IdentityInfo: &IdentityInfo{
			IDDocType: "IDENTIFICATION_TYPE_IDCARD", // 证件类型
			IDCardInfo: &IDCardInfo{
				IDCardCopy:      "xxx",                // 身份证人像面照片
				IDCardNational:  "xxx",                // 身份证国徽面照片
				IDCardName:      "张三",                 // 身份证姓名
				IDCardNumber:    "110101199003070073", // 身份证号码
				CardPeriodBegin: "2026-06-06",         // 身份证有效期开始时间
				CardPeriodEnd:   "2026-06-06",         // 身份证有效期结束时间
			},
		},
	},
	BusinessInfo: BusinessInfo{
		MerchantShortname: "张三的小店",       // 商户简称
		ServicePhone:      "13900000000", // 客服电话
	},
	SettlementInfo: SettlementInfo{
		SettlementID:      SettlementRuleIDMicro, // 入驻结算规则ID
		QualificationType: "家政",                  // 所属行业
	},
	BankAccountInfo: BankAccountInfo{
		BankAccountType: "BANK_ACCOUNT_TYPE_PERSONAL", // 账户类型，小微商户固定为 BANK_ACCOUNT_TYPE_PERSONAL
		AccountName:     "张三",                         // 开户名称
		AccountBank:     "工商银行",                       // 开户银行
		AccountNumber:   "6212262201023557000",        // 银行账号
	},
}

// TestSubmitApplyment 提交申请单示例
func TestSubmitApplyment(t *testing.T) {
	mgr, err := vwechatpay.NewManagerFromEnv()
	if err != nil {
		t.Skipf("初始化微信支付失败, 跳过测试: %v", err)
		return
	}

	client := NewApply4SubClient(mgr)

	// 准备请求参数
	ctx := context.Background()

	// 加密敏感信息
	contactName, err := client.mgr.PlatManager.Encrypt("张三")
	if err != nil {
		log.Fatalf("加密联系人姓名失败: %v", err)
	}

	contactIDNumber, err := client.mgr.PlatManager.Encrypt("110101199003070073")
	if err != nil {
		log.Fatalf("加密联系人身份证号失败: %v", err)
	}

	mobilePhone, err := client.mgr.PlatManager.Encrypt("13900000000")
	if err != nil {
		log.Fatalf("加密手机号失败: %v", err)
	}

	contactEmail, err := client.mgr.PlatManager.Encrypt("test@example.com")
	if err != nil {
		log.Fatalf("加密邮箱失败: %v", err)
	}

	// 构建请求
	req := &ApplymentRequest{
		BusinessCode: "1900000001_10000", // 业务申请编号，服务商自定义
		ContactInfo: ContactInfo{
			ContactName:     contactName,
			ContactIDNumber: contactIDNumber,
			OpenID:          "oUpF8uMuAJO_M2pxb1Q9zNjWeS6o", // 超级管理员微信openid
			MobilePhone:     mobilePhone,
			ContactEmail:    contactEmail,
		},
		SubjectInfo: SubjectInfo{
			SubjectType: SubjectTypeMicro, // 小微商户固定为 SubjectTypeMicro
			MicroBizInfo: MicroBizInfo{
				MicroBizType: MicroTypeStore, // 门店场所
				MicroStoreInfo: &MicroStoreInfo{
					MicroName:        "张三的小店",
					MicroAddressCode: "440305", // 深圳市南山区
					MicroAddress:     "南山区xx大厦x层xxxx室",
					StoreEntrancePic: "media_id1", // 已上传的门店门头照片媒体文件标识
					MicroIndoorCopy:  "media_id2", // 已上传的店内环境照片媒体文件标识
				},
			},
		},
		BusinessInfo: BusinessInfo{
			MerchantShortname: "张三的小店",       // 商户简称
			ServicePhone:      "13900000000", // 客服电话
			SalesInfo:         SalesInfo{},   // 经营场景
		},
		SettlementInfo: SettlementInfo{
			SettlementID:        "719",      // 入驻结算规则ID
			QualificationType:   "餐饮",       // 所属行业
			Qualifications:      []string{}, // 特殊资质
			BusinessAdditionPic: "",         // 补充材料
		},
		BankAccountInfo: BankAccountInfo{
			BankAccountType: "BANK_ACCOUNT_TYPE_PERSONAL", // 账户类型，小微商户固定为 BANK_ACCOUNT_TYPE_PERSONAL
			AccountName:     "张三",                         // 开户名称
			AccountBank:     "工商银行",                       // 开户银行
			BankAddressCode: "110000",                     // 开户银行省市编码
			BankBranchID:    "402713354941",               // 开户银行联行号
			BankName:        "中国工商银行股份有限公司北京海淀支行",         // 开户银行全称（含支行）
			AccountNumber:   "6212262201023557000",        // 银行账号
		},
	}

	// 发送请求
	resp, err := client.SubmitApplyment(ctx, req)
	if err != nil {
		log.Fatalf("提交申请单失败: %v", err)
	}

	// 处理响应
	fmt.Printf("申请单提交成功，申请单号: %d\n", resp.ApplymentID)
}

// ExampleQueryApplyment 查询申请单状态示例
func TestQueryApplyment(t *testing.T) {
	mgr, err := vwechatpay.NewManagerFromEnv()
	if err != nil {
		t.Skipf("初始化微信支付失败, 跳过测试: %v", err)
		return
	}

	client := NewApply4SubClient(mgr)

	// 准备请求参数
	ctx := context.Background()

	// 通过业务申请编号查询
	businessCode := "1900000001_10000"
	resp1, err := client.QueryApplymentByBusinessCode(ctx, businessCode)
	if err != nil {
		log.Fatalf("通过业务申请编号查询失败: %v", err)
	}

	// 处理响应
	fmt.Printf("业务申请编号: %s\n", resp1.BusinessCode)
	fmt.Printf("申请单号: %d\n", resp1.ApplymentID)
	fmt.Printf("特约商户号: %s\n", resp1.SubMchID)
	fmt.Printf("申请单状态: %s\n", resp1.ApplymentState)
	fmt.Printf("申请单状态描述: %s\n", resp1.ApplymentStateMsg)

	// 通过申请单号查询
	applymentID := int64(2000002124775691)
	resp2, err := client.QueryApplymentByApplymentID(ctx, applymentID)
	if err != nil {
		log.Fatalf("通过申请单号查询失败: %v", err)
	}

	// 处理响应
	fmt.Printf("业务申请编号: %s\n", resp2.BusinessCode)
	fmt.Printf("申请单号: %d\n", resp2.ApplymentID)
	fmt.Printf("特约商户号: %s\n", resp2.SubMchID)
	fmt.Printf("申请单状态: %s\n", resp2.ApplymentState)
	fmt.Printf("申请单状态描述: %s\n", resp2.ApplymentStateMsg)
}

// TestSubmitApplymentWithOnlineInfo 提交线上商品/服务交易申请单示例
func TestSubmitApplymentWithOnlineInfo(t *testing.T) {
	mgr, err := vwechatpay.NewManagerFromEnv()
	if err != nil {
		t.Errorf("初始化微信支付失败: %v", err)
		return
	}
	client := NewApply4SubClient(mgr)

	// 准备请求参数
	ctx := context.Background()

	// 加密敏感信息
	contactName, err := client.mgr.PlatManager.Encrypt("xxx")
	if err != nil {
		log.Fatalf("加密联系人姓名失败: %v", err)
	}

	contactIDNumber, err := client.mgr.PlatManager.Encrypt("xxx")
	if err != nil {
		log.Fatalf("加密联系人身份证号失败: %v", err)
	}

	mobilePhone, err := client.mgr.PlatManager.Encrypt("xxx")
	if err != nil {
		log.Fatalf("加密手机号失败: %v", err)
	}

	contactEmail, err := client.mgr.PlatManager.Encrypt("xxx")
	if err != nil {
		log.Fatalf("加密邮箱失败: %v", err)
	}

	idCardName, err := client.mgr.PlatManager.Encrypt("xxx")
	if err != nil {
		log.Fatalf("加密身份证姓名失败: %v", err)
	}

	idCardNumber, err := client.mgr.PlatManager.Encrypt("xxx")
	if err != nil {
		log.Fatalf("加密身份证号码失败: %v", err)
	}

	accountName, err := client.mgr.PlatManager.Encrypt("xxx")
	if err != nil {
		log.Fatalf("加密开户名称失败: %v", err)
	}

	accountNumber, err := client.mgr.PlatManager.Encrypt("xxx")
	if err != nil {
		log.Fatalf("加密银行账号失败: %v", err)
	}

	// 构建请求
	req := &ApplymentRequest{
		BusinessCode: "1900013511_10000", // 业务申请编号，服务商自定义
		ContactInfo: ContactInfo{
			ContactName:     contactName,
			ContactIDNumber: contactIDNumber,
			OpenID:          "oUpF8uMuAJO_M2pxb1Q9zNjWeS6o", // 超级管理员微信openid
			MobilePhone:     mobilePhone,
			ContactEmail:    contactEmail,
		},
		SubjectInfo: SubjectInfo{
			SubjectType: SubjectTypeMicro, // 小微商户固定为 SubjectTypeMicro
			MicroBizInfo: MicroBizInfo{
				MicroBizType: MicroTypeStore, // 门店场所
				MicroStoreInfo: &MicroStoreInfo{
					MicroName:        "大郎烧饼",
					MicroAddressCode: "440305", // 深圳市南山区
					MicroAddress:     "南山区xx大厦x层xxxx室",
					StoreEntrancePic: "xxx",        // 已上传的门店门头照片媒体文件标识
					MicroIndoorCopy:  "xxx",        // 已上传的店内环境照片媒体文件标识
					StoreLongitude:   "113.941355", // 门店经度
					StoreLatitude:    "22.546245",  // 门店纬度
				},
				MicroOnlineInfo: &MicroOnlineInfo{
					MicroOnlineStore: "李三服装店",              // 线上店铺名称
					MicroEcName:      "XX购物平台",             // 电商平台名称
					MicroQrcode:      "https://www.qq.com", // 线上店铺二维码
					MicroLink:        "https://www.qq.com", // 线上店铺链接
				},
			},
			IdentityInfo: &IdentityInfo{
				IDDocType: "IDENTIFICATION_TYPE_IDCARD", // 证件类型
				IDCardInfo: &IDCardInfo{
					IDCardCopy:      "xxx", // 身份证人像面照片
					IDCardNational:  "xxx", // 身份证国徽面照片
					IDCardName:      idCardName,
					IDCardNumber:    idCardNumber,
					CardPeriodBegin: "2026-06-06", // 身份证有效期开始时间
					CardPeriodEnd:   "2026-06-06", // 身份证有效期结束时间
				},
			},
		},
		BusinessInfo: BusinessInfo{
			MerchantShortname: "张三餐饮店",     // 商户简称
			ServicePhone:      "0758XXXXX", // 客服电话
		},
		SettlementInfo: SettlementInfo{
			SettlementID:         "703",                                    // 入驻结算规则ID
			QualificationType:    "餐饮",                                     // 所属行业
			Qualifications:       []string{"example_qualifications"},       // 特殊资质
			ActivitiesID:         "716",                                    // 优惠费率活动ID
			ActivitiesRate:       "0.6",                                    // 优惠费率活动值
			DebitActivitiesRate:  "0.6",                                    // 优惠费率活动值（借记卡）
			CreditActivitiesRate: "0.6",                                    // 优惠费率活动值（贷记卡）
			ActivitiesAdditions:  []string{"example_activities_additions"}, // 优惠费率活动补充材料
		},
		BankAccountInfo: BankAccountInfo{
			BankAccountType: "BANK_ACCOUNT_TYPE_CORPORATE", // 账户类型
			AccountName:     accountName,                   // 开户名称
			AccountBank:     "工商银行",                        // 开户银行
			BankAddressCode: "110000",                      // 开户银行省市编码
			BankBranchID:    "402713354941",                // 开户银行联行号
			BankName:        "施秉县农村信用合作联社城关信用社",            // 开户银行全称（含支行）
			AccountNumber:   accountNumber,                 // 银行账号
		},
		AdditionInfo: &AdditionInfo{
			LegalPersonCommitment: "xxx",                                      // 法人开户承诺函
			LegalPersonVideo:      "xxx",                                      // 法人开户意愿视频
			BusinessAdditionPics:  []string{"example_business_addition_pics"}, // 补充材料
			BusinessAdditionMsg:   "特殊情况，说明原因",                                // 补充说明
		},
	}

	// 发送请求
	resp, err := client.SubmitApplyment(ctx, req)
	if err != nil {
		log.Fatalf("提交申请单失败: %v", err)
	}

	// 处理响应
	fmt.Printf("申请单提交成功，申请单号: %d\n", resp.ApplymentID)
}
