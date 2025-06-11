部分微信支付业务指定商户需要使用图片上传 API来上报图片、视频信息，从而获得必传参数的值：MediaID 。

- 图片上传: https://pay.weixin.qq.com/doc/v3/partner/4012760490
    - https://github.com/wechatpay-apiv3/wechatpay-go/blob/main/services/fileuploader/image_uploader.go
- 视频上传: https://pay.weixin.qq.com/doc/v3/partner/4012761084
    - https://github.com/wechatpay-apiv3/wechatpay-go/blob/main/services/fileuploader/video_uploader.go