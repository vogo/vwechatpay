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

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/vogo/vwechatpay"
)

// TestSubmitApplyment 提交申请单示例
func TestSubmitApplyment(t *testing.T) {
	mgr, err := vwechatpay.NewManagerFromEnv()
	if err != nil {
		t.Errorf("初始化微信支付失败: %v", err)
		return
	}
	client := NewPartnerClient(mgr)

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
			SubjectType: "SUBJECT_TYPE_MICRO", // 小微商户固定为 SUBJECT_TYPE_MICRO
			MicroBizInfo: MicroBizInfo{
				MicroBizType: "MICRO_TYPE_STORE", // 门店场所
				MicroStoreInfo: &MicroStoreInfo{
					MicroName:        "张三的小店",
					MicroAddressCode: "440305", // 深圳市南山区
					MicroAddress:     "南山区xx大厦x层xxxx室",
					StoreEntrancePic: "media_id1", // 已上传的门店门头照片媒体文件标识
					MicroIndoorCopy:  "media_id2", // 已上传的店内环境照片媒体文件标识
					StoreStreetPic:   "media_id3", // 已上传的门店街景照片媒体文件标识
				},
			},
		},
		BusinessInfo: BusinessInfo{
			MerchantShortname: "张三的小店",              // 商户简称
			ServicePhone:      "13900000000",        // 客服电话
			SalesInfo:         "sales_scenes_store", // 经营场景
		},
		SettlementInfo: SettlementInfo{
			SettlementID:        "719", // 入驻结算规则ID
			QualificationType:   "餐饮",  // 所属行业
			Qualifications:      "",    // 特殊资质
			BusinessAdditionPic: "",    // 补充材料
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
		t.Errorf("初始化微信支付失败: %v", err)
		return
	}
	client := NewPartnerClient(mgr)

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
