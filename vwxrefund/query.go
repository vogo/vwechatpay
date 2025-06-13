package vwxrefund

import (
	"context"
	"fmt"
	"net/http"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

// QueryByOutRefundNo 通过商户退款单号查询单笔退款
// 商户可以通过商户退款单号查询单笔退款，仅能查询自有退款单号的退款单，支持商户使用API、商户平台、微信支付小程序等多种方式发起的退款单查询。
func (c *RefundClient) QueryByOutRefundNo(ctx context.Context, subMchID, outRefundNo string) (*refunddomestic.Refund, error) {
	req := refunddomestic.QueryByOutRefundNoRequest{
		OutRefundNo: core.String(outRefundNo),
	}

	// 服务商模式下需要传递子商户号
	if subMchID != "" {
		req.SubMchid = core.String(subMchID)
	}

	resp, result, err := c.refundApi.QueryByOutRefundNo(ctx, req)
	if err != nil {
		return nil, err
	}

	if result.Response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("query refund failed, status code: %d", result.Response.StatusCode)
	}

	return resp, nil
}
