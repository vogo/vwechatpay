package vwxrefund

import (
	"context"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

// CreateRefund 申请退款
// 当交易发生之后一段时间内，由于买家或者卖家的原因需要退款时，卖家可以通过退款接口将支付款退还给买家，
// 微信支付将在收到退款请求并且验证成功之后，将支付款按照退款规则原路退还给买家。
func (c *RefundClient) CreateRefund(ctx context.Context, req *refunddomestic.CreateRequest) (*refunddomestic.Refund, *core.APIResult, error) {
	return c.refundApi.Create(ctx, *req)
}

// CreateRefundWithAmount 申请退款（简化版）
// 提供了一个简化版的退款接口，只需要提供必要的参数
func (c *RefundClient) CreateRefundWithAmount(ctx context.Context, outRefundNo, transactionID, outTradeNo, reason string,
	refundAmount, totalAmount int64, subMchID string) (*refunddomestic.Refund, *core.APIResult, error) {
	// 构建退款金额信息
	amountReq := &refunddomestic.AmountReq{
		Refund:   core.Int64(refundAmount),
		Total:    core.Int64(totalAmount),
		Currency: core.String("CNY"),
	}

	// 构建退款请求
	req := &refunddomestic.CreateRequest{
		OutRefundNo: core.String(outRefundNo),
		Reason:      core.String(reason),
		Amount:      amountReq,
	}

	// 设置交易单号（优先使用微信支付订单号）
	if transactionID != "" {
		req.TransactionId = core.String(transactionID)
	} else if outTradeNo != "" {
		req.OutTradeNo = core.String(outTradeNo)
	}

	// 服务商模式下需要传递子商户号
	if subMchID != "" {
		req.SubMchid = core.String(subMchID)
	}

	return c.CreateRefund(ctx, req)
}
