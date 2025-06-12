package vwxpartnerjsapi

import (
	"context"
	"fmt"

	"github.com/vogo/vogo/vencoding/vjson"
	"github.com/vogo/vogo/vlog"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnerpayments/jsapi"
)

// QueryOrderById 根据微信支付订单号查询订单
// subMchID: 子商户号
// transactionId: 微信支付订单号
func (c *PartnerJsApiClient) QueryOrderById(ctx context.Context, subMchID, transactionId string) (*partnerpayments.Transaction, error) {
	// 构建请求参数
	req := jsapi.QueryOrderByIdRequest{
		TransactionId: core.String(transactionId),
		SpMchid:       core.String(c.mgr.Config.MerchantID),
		SubMchid:      core.String(subMchID),
	}

	vlog.Infof("partner jsapi query order, subMchID: %s, transactionId: %s", subMchID, transactionId)

	// 发送请求
	resp, result, err := c.jsapiApiService.QueryOrderById(ctx, req)
	if err != nil {
		vlog.Errorf("query order by id error: %v", err)
		return nil, err
	}

	vlog.Infof("partner jsapi query order response: %s", vjson.EnsureMarshal(resp))

	if result.Response.StatusCode != 200 {
		return nil, fmt.Errorf("query order by id failed with status code: %d", result.Response.StatusCode)
	}

	return resp, nil
}

// QueryOrderByOutTradeNo 根据商户订单号查询订单
// subMchID: 子商户号
// outTradeNo: 商户订单号
func (c *PartnerJsApiClient) QueryOrderByOutTradeNo(ctx context.Context, subMchID, outTradeNo string) (*partnerpayments.Transaction, error) {
	// 构建请求参数
	req := jsapi.QueryOrderByOutTradeNoRequest{
		OutTradeNo: core.String(outTradeNo),
		SpMchid:    core.String(c.mgr.Config.MerchantID),
		SubMchid:   core.String(subMchID),
	}

	vlog.Infof("partner jsapi query request, subMchID: %s, outTradeNo: %s", subMchID, outTradeNo)

	// 发送请求
	resp, result, err := c.jsapiApiService.QueryOrderByOutTradeNo(ctx, req)
	if err != nil {
		vlog.Errorf("query order by out trade no error: %v", err)
		return nil, err
	}

	vlog.Infof("partner jsapi query order response: %s", vjson.EnsureMarshal(resp))

	if result.Response.StatusCode != 200 {
		return nil, fmt.Errorf("query order by out trade no failed with status code: %d", result.Response.StatusCode)
	}

	return resp, nil
}
