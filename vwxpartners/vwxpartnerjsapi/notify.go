package vwxpartnerjsapi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// PartnerJsApiNotifyParse 解析服务商模式 JSAPI 支付回调通知
func (c *PartnerJsApiClient) PartnerJsApiNotifyParse(headerFetcher func(string) string, body []byte) (*notify.Request, map[string]interface{}, error) {
	ctx := context.Background()

	// 验证回调通知签名
	headerArgs, err := getWechatPayHeader(headerFetcher)
	if err != nil {
		return nil, nil, err
	}

	if err = checkWechatPayHeader(ctx, headerArgs); err != nil {
		return nil, nil, err
	}

	message := buildMessage(ctx, headerArgs, body)

	if err = c.mgr.PlatManager.LoadVerifier().Verify(ctx, headerArgs.Serial, message, headerArgs.Signature); err != nil {
		return nil, nil, fmt.Errorf(
			"validate verify fail serial=[%s] request-id=[%s] err=%w",
			headerArgs.Serial, headerArgs.RequestID, err,
		)
	}

	return c.PartnerJsApiNotifyParseBody(body)
}

// PartnerJsApiNotifyParseBody 解析服务商模式 JSAPI 支付回调通知体
func (c *PartnerJsApiClient) PartnerJsApiNotifyParseBody(body []byte) (*notify.Request, map[string]interface{}, error) {
	ret := new(notify.Request)
	if err := json.Unmarshal(body, ret); err != nil {
		return nil, nil, fmt.Errorf("parse request body error: %v", err)
	}

	plaintext, err := utils.DecryptAES256GCM(
		c.mgr.Config.MerchantAPIv3Key, ret.Resource.AssociatedData, ret.Resource.Nonce, ret.Resource.Ciphertext,
	)
	if err != nil {
		return ret, nil, fmt.Errorf("decrypt request error: %v", err)
	}

	ret.Resource.Plaintext = plaintext

	content := map[string]interface{}{}
	if err = json.Unmarshal([]byte(plaintext), &content); err != nil {
		return ret, nil, fmt.Errorf("unmarshal plaintext to content failed: %v", err)
	}

	return ret, content, nil
}
