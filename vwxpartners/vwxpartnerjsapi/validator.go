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

package vwxpartnerjsapi

import (
	"context"
	"fmt"
)

// ValidateHTTPMessage 验证 HTTP 消息
func (c *PartnerJsApiClient) ValidateHTTPMessage(ctx context.Context, headerFetcher func(string) string, body []byte) error {
	headerArgs, err := getWechatPayHeader(headerFetcher)
	if err != nil {
		return err
	}

	if err = checkWechatPayHeader(ctx, headerArgs); err != nil {
		return err
	}

	message := buildMessage(ctx, headerArgs, body)

	if err = c.mgr.PlatManager.LoadVerifier().Verify(ctx, headerArgs.Serial, message, headerArgs.Signature); err != nil {
		return fmt.Errorf(
			"validate verify fail serial=[%s] request-id=[%s] err=%w",
			headerArgs.Serial, headerArgs.RequestID, err,
		)
	}

	return nil
}
