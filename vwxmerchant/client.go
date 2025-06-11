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

package vwxmerchant

import (
	"github.com/vogo/vwechatpay"
	"github.com/wechatpay-apiv3/wechatpay-go/services/fileuploader"
)

// MerchantClient 微信商户客户端
type MerchantClient struct {
	mgr           *vwechatpay.Manager
	imageUploader *fileuploader.ImageUploader
	videoUploader *fileuploader.VideoUploader
}

// NewMerchantClient 创建商户客户端
func NewMerchantClient(mgr *vwechatpay.Manager) *MerchantClient {
	return &MerchantClient{
		mgr: mgr,
		imageUploader: &fileuploader.ImageUploader{
			Client: mgr.Client,
		},
		videoUploader: &fileuploader.VideoUploader{
			Client: mgr.Client,
		},
	}
}
