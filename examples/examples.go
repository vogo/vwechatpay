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

package main

import (
	"context"
	"fmt"

	"github.com/vogo/vwechatpay"
	"github.com/vogo/vwechatpay/vwxfund/vwxmchbalance"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
)

func main() {
	mgr, err := vwechatpay.NewManagerFromEnv()
	if err != nil {
		panic(err)
	}
	cert := mgr.PlatManager.LoadCert()
	fmt.Printf("cert serial number: %s\n", cert.SerialNumber)

	balanceClient := vwxmchbalance.NewMchBalanceClient(mgr)
	resp, err := balanceClient.QueryBalance(context.Background(), vwxmchbalance.AccountTypeOperation)
	if err != nil {
		if wxerr, ok := err.(*core.APIError); ok {
			fmt.Printf("query balance error, code: %s, message: %s\n", wxerr.Code, wxerr.Message)
		} else {
			fmt.Printf("query balance error: %v\n", err)
		}
	} else {
		fmt.Printf("available amount: %d\n", resp.AvailableAmount)
		if resp.PendingAmount != nil {
			fmt.Printf("pending amount: %d\n", *resp.PendingAmount)
		}
	}
}
