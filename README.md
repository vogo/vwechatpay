# vwechatpay

本项目包装微信支付 [go sdk](https://github.com/wechatpay-apiv3/wechatpay-go) 的一些支付能力，降低使用门槛。

## 功能特点

- 简化微信支付API的调用流程，提供更友好的接口封装
- 支持普通商户和服务商模式的支付处理
- 支持JSAPI支付、APP支付等多种支付方式
- 提供退款、订单查询、关闭订单等完整功能
- 支持支付回调通知的处理和验证
- 提供商户进件、资金账户等扩展功能

## 项目结构

```
├── vwxpayments     # 支付相关功能
│   ├── vwxjsapi    # JSAPI支付（公众号、小程序支付）
│   └── vwxapp      # APP支付
├── vwxpartners     # 服务商模式相关功能
│   ├── vwxpartnerjsapi  # 服务商JSAPI支付
│   └── vwxpartnerapp    # 服务商APP支付
├── vwxrefund       # 退款相关功能
├── vwxapply4sub    # 商户进件相关功能
├── vwxcapital      # 资金账户相关功能
├── vwxmerchant     # 商户相关功能
├── vwxplat         # 微信支付平台相关功能
└── vwxutils        # 工具函数
```

## 安装

```bash
go get github.com/vogo/vwechatpay
```

## 基本用法

### 配置初始化

可以通过环境变量或直接设置配置参数来初始化微信支付管理器：

```go
// 方式1：通过环境变量初始化
// 需要设置以下环境变量：
// WECHAT_PAY_MERCHANT_ID - 商户号
// WECHAT_PAY_MERCHANT_CERT_SERIAL_NO - 商户证书序列号
// WECHAT_PAY_MERCHANT_APIV3_KEY - 商户APIv3密钥
// WECHAT_PAY_APP_ID - 应用ID
// WECHAT_PAY_PRIVATE_KEY_PATH 或 WECHAT_PAY_PRIVATE_KEY_CONTENT - 私钥路径或内容
// WECHAT_PAY_CERT_PATH 或 WECHAT_PAY_CERT_CONTENT - 证书路径或内容
mgr, err := vwechatpay.NewManagerFromEnv()
if err != nil {
    // 处理错误
}

// 方式2：直接设置配置参数
cfg := &vwechatpay.Config{
    MerchantID:           "商户号",
    MerchantCertSerialNO: "商户证书序列号",
    MerchantAPIv3Key:     "商户APIv3密钥",
    PrivateKeyPath:       "私钥文件路径",  // 或使用 PrivateKeyContent
    CertPath:             "证书文件路径",  // 或使用 CertContent
    AppID:                "应用ID",
}

mgr, err := vwechatpay.NewManager(cfg)
if err != nil {
    // 处理错误
}
```

### JSAPI支付（公众号/小程序支付）

```go
// 创建JSAPI支付客户端
jsapiClient := vwxjsapi.NewJsApiClient(mgr)

// 发起预支付
ctx := context.Background()
payParams, err := jsapiClient.Prepay(
    ctx,
    "用户的OpenID",
    100,  // 金额，单位：分
    "商户订单号",
    "商品描述",
    "附加数据",
    "回调通知URL",
    time.Now().Add(30 * time.Minute),  // 订单过期时间
)

// 处理支付结果
if err != nil {
    // 处理错误
}

// payParams 包含了前端调起支付所需的参数
// 返回给前端，用于调起微信支付
```

### 查询订单

```go
// 通过微信支付订单号查询
transaction, err := jsapiClient.QueryOrderById(ctx, "微信支付订单号")

// 或通过商户订单号查询
transaction, err := jsapiClient.QueryOrderByOutTradeNo(ctx, "商户订单号")
```

### 关闭订单

```go
// 关闭订单
err := jsapiClient.CloseOrder(ctx, "商户订单号")
```

### 申请退款

```go
// 创建退款客户端
refundClient := vwxrefund.NewRefundClient(mgr)

// 简化版退款申请
refund, err := refundClient.CreateRefundWithAmount(
    ctx,
    "商户退款单号",
    "微信支付订单号",  // 与商户订单号二选一
    "商户订单号",     // 与微信支付订单号二选一
    "退款原因",
    100,  // 退款金额，单位：分
    100,  // 订单总金额，单位：分
    "",   // 子商户号，服务商模式下使用
)
```

### 处理支付回调通知

```go
// 解析支付通知
notifyReq, notifyContent, err := jsapiClient.JsApiNotifyParse(
    func(key string) string {
        // 从HTTP请求头中获取对应的值
        return r.Header.Get(key)
    },
    requestBody,  // HTTP请求体
)

// 处理通知内容
if err != nil {
    // 处理错误
}

// notifyContent 包含了支付结果信息
// 根据业务需求处理支付结果
```

## 服务商模式

### 服务商JSAPI支付

```go
// 创建服务商JSAPI支付客户端
partnerJsapiClient := vwxpartnerjsapi.NewJsApiClient(mgr)

// 发起服务商模式预支付
ctx := context.Background()
payParams, err := partnerJsapiClient.Prepay(
    ctx,
    "用户的OpenID",
    100,  // 金额，单位：分
    "商户订单号",
    "商品描述",
    "附加数据",
    "回调通知URL",
    time.Now().Add(30 * time.Minute),  // 订单过期时间
    "服务商商户号",
    "子商户号",
)

// 查询服务商模式订单
transaction, err := partnerJsapiClient.QueryOrderById(ctx, "微信支付订单号", "服务商商户号", "子商户号")
```

## 高级功能

### 商户进件

```go
// 创建商户进件客户端
apply4subClient := vwxapply4sub.NewApply4SubClient(mgr)

// 提交商户进件申请
resp, err := apply4subClient.Submit(ctx, applyRequest)
```

## 最佳实践

### 日志记录

建议在生产环境中实现完善的日志记录，包括：

- 请求参数和响应结果（注意脱敏敏感信息）
- 错误信息和堆栈跟踪
- 关键业务流程的执行时间

### 安全建议

- 妥善保管商户私钥和APIv3密钥
- 定期更新证书和密钥
- 实现IP白名单限制回调通知
- 对敏感数据进行加密存储

## 贡献代码

欢迎提交 issue 和 pull request，一起完善本项目。

### 提交 Issue

- 使用清晰的标题描述问题
- 详细描述问题的复现步骤
- 提供相关的日志和错误信息

### 提交 Pull Request

- 确保代码符合 Go 的代码规范
- 提供详细的描述说明修改的内容和原因
- 确保所有测试通过

## 问题反馈

如果您在使用过程中遇到任何问题，可以通过以下方式获取帮助：

- [GitHub Issues](https://github.com/vogo/vwechatpay/issues)
- [微信支付官方文档](https://pay.weixin.qq.com/wiki/doc/apiv3/index.shtml)

## 许可证

本项目采用 MIT 许可证，详情请参阅 LICENSE 文件。