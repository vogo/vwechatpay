# vwxmchbalance - 商户账户余额查询

本包提供微信支付商户账户实时余额查询功能。

## 功能特点

- 查询基本账户、运营账户、手续费账户的实时余额
- 支持查看可用余额和冻结余额
- 仅支持普通服务商使用

## 使用示例

### 初始化客户端

```go
import (
    "context"
    "github.com/vogo/vwechatpay"
    "github.com/vogo/vwechatpay/vwxfund/vwxmchbalance"
)

// 创建微信支付管理器
mgr, err := vwechatpay.NewManagerFromEnv()
if err != nil {
    // 处理错误
}

// 创建账户余额查询客户端
balanceClient := vwxmchbalance.NewMchBalanceClient(mgr)
```

### 查询账户余额

```go
ctx := context.Background()

// 查询基本账户余额
balance, err := balanceClient.QueryBalance(ctx, vwxmchbalance.AccountTypeBasic)
if err != nil {
    // 处理错误
}

// 可用余额(单位:分)
fmt.Printf("可用余额: %d 分\n", balance.AvailableAmount)

// 冻结余额(单位:分)
if balance.PendingAmount != nil {
    fmt.Printf("冻结余额: %d 分\n", *balance.PendingAmount)
}
```

### 账户类型

支持以下三种账户类型:

```go
// 基本账户
vwxmchbalance.AccountTypeBasic

// 运营账户
vwxmchbalance.AccountTypeOperation

// 手续费账户
vwxmchbalance.AccountTypeFees
```

## API 响应

### BalanceResponse

| 字段 | 类型 | 说明 |
|------|------|------|
| AvailableAmount | int64 | 可用余额(单位:分),可用于提现等操作 |
| PendingAmount | *int64 | 不可用余额(单位:分),冻结金额不可进行提现等操作 |

## 错误处理

常见错误码:

- `PARAM_ERROR` (400): 参数校验失败
- `SIGN_ERROR` (401): 签名验证失败
- `NO_AUTH` (403): 商户无接口权限
- `INVALID_REQUEST` (400): 账户类型未开通
- `SYSTEM_ERROR` (500): 系统暂时不可用

## 注意事项

1. 此接口仅支持普通服务商使用
2. 所有金额单位均为分(1元 = 100分)
3. 冻结余额(PendingAmount)可能为空,需要检查是否为 nil
4. 需要确保商户已开通相应的账户类型

## 参考文档

- [查询账户实时余额API文档](https://pay.weixin.qq.com/doc/v3/partner/4012720926)
