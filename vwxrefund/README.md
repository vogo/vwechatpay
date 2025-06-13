微信订单退款: 
- 商户版: https://pay.weixin.qq.com/doc/v3/merchant/4013071001
- 服务商: https://pay.weixin.qq.com/doc/v3/partner/4013080622

NOTE: 商户版、服务商版都是同一套退款API.

商户版退款API:
- 申请退款: https://pay.weixin.qq.com/doc/v3/merchant/4013071036
- 查询单笔退款（通过商户退款单号）: https://pay.weixin.qq.com/doc/v3/merchant/4013071041
- 发起异常退款: https://pay.weixin.qq.com/doc/v3/merchant/4013071193
- 退款结果通知: https://pay.weixin.qq.com/doc/v3/merchant/4013071196

服务商版退款API:
- 申请退款: https://pay.weixin.qq.com/doc/v3/partner/4013080625
- 查询单笔退款（通过商户退款单号）: https://pay.weixin.qq.com/doc/v3/partner/4013080626
- 发起异常退款: https://pay.weixin.qq.com/doc/v3/partner/4013080627
- 退款结果通知: https://pay.weixin.qq.com/doc/v3/partner/4013080628


调用官方SDK接口： https://github.com/wechatpay-apiv3/wechatpay-go/blob/main/services/refunddomestic/api_refunds.go