package vwxpartnerjsapi

import (
	"context"
	"fmt"

	"github.com/vogo/vogo/vlog"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/jsapi"
)

// CloseOrder 关闭订单
// 以下情况需要调用关单接口：
// 1. 商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；
// 2. 系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。
// subMchID: 子商户号
// outTradeNo: 商户订单号
func (c *PartnerJsApiClient) CloseOrder(ctx context.Context, subMchID, outTradeNo string) error {
	// 构建请求参数
	req := jsapi.CloseOrderRequest{
		SpMchid:    core.String(c.mgr.Config.MerchantID),
		SubMchid:   core.String(subMchID),
		OutTradeNo: core.String(outTradeNo),
	}

	vlog.Infof("partner jsapi close order, subMchID: %s, outTradeNo: %s", subMchID, outTradeNo)

	// 发送请求
	result, err := c.jsapiApiService.CloseOrder(ctx, req)
	if err != nil {
		vlog.Errorf("close order error: %v", err)
		return err
	}

	vlog.Infof("partner jsapi close order response status: %d", result.Response.StatusCode)

	if result.Response.StatusCode != 204 {
		return fmt.Errorf("close order failed with status code: %d", result.Response.StatusCode)
	}

	return nil
}
