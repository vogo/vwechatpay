package vwxrefund

import (
	"github.com/vogo/vwechatpay"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

// RefundClient 微信退款客户端
type RefundClient struct {
	mgr       *vwechatpay.Manager
	refundApi *refunddomestic.RefundsApiService
}

func NewRefundClient(mgr *vwechatpay.Manager) *RefundClient {
	return &RefundClient{
		mgr: mgr,
		refundApi: &refunddomestic.RefundsApiService{
			Client: mgr.Client,
		},
	}
}
