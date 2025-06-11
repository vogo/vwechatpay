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
	"context"
	"fmt"
	"io"
	"net/http"
)

// MerchantImageUpload 上传图片
// 返回值: 媒体文件标识ID
func (s *MerchantClient) MerchantImageUpload(ctx context.Context, fileReader io.Reader, filename, contentType string) (string, error) {
	resp, result, err := s.imageUploader.Upload(ctx, fileReader, filename, contentType)
	if err != nil {
		return "", err
	}

	if result.Response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("upload image failed, status code: %d", result.Response.StatusCode)
	}

	return *resp.MediaId, nil
}
