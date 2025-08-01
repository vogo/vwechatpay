package vwxmchtransfer

const (
	StateAccepted        = "ACCEPTED"          // 转账已受理
	StateProcessing      = "PROCESSING"        // 转账锁定资金中。如果一直停留在该状态，建议检查账户余额是否足够，如余额不足，可充值后再原单重试。
	StateWaitUserConfirm = "WAIT_USER_CONFIRM" // 待收款用户确认，可拉起微信收款确认页面进行收款确认
	StateTransfering     = "TRANSFERING"       // 转账中，可拉起微信收款确认页面再次重试确认收款
	StateSuccess         = "SUCCESS"           // 转账成功
	StateFail            = "FAIL"              // 转账失败
	StateCanceling       = "CANCELING"         // 商户撤销请求受理成功，该笔转账正在撤销中
	StateCancelled       = "CANCELLED"         // 转账撤销完成
)

func StateText(state string) string {
	switch state {
	case StateAccepted:
		return "转账已受理"
	case StateProcessing:
		return "转账锁定资金中"
	case StateWaitUserConfirm:
		return "待收款用户确认"
	case StateTransfering:
		return "转账中"
	case StateSuccess:
		return "转账成功"
	case StateFail:
		return "转账失败"
	case StateCanceling:
		return "商户撤销请求受理成功"
	case StateCancelled:
		return "转账撤销完成"
	default:
		return "未知状态"
	}
}

func IsStateSuccess(state string) bool {
	return state == StateSuccess
}

func IsStateProcessing(state string) bool {
	return state == StateAccepted ||
		state == StateProcessing ||
		state == StateWaitUserConfirm ||
		state == StateTransfering
}

func IsStateWaitUserConfirm(state string) bool {
	return state == StateWaitUserConfirm
}

func IsStateCancel(state string) bool {
	return state == StateCanceling || state == StateCancelled
}

func IsStateFail(state string) bool {
	return state == StateFail
}
