package main

import (
	"github.com/vogo/vwechatpay"
)

func main() {
	mgr, err := vwechatpay.NewManagerFromEnv()
	if err != nil {
		panic(err)
	}
	cert := mgr.PlatManager.LoadCert()
	println(cert)
}
