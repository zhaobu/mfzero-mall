package svc

import (
	"mfzero-mall/app/pay/api/internal/config"
	"mfzero-mall/app/pay/rpc/payclient"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	PayRpc payclient.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: payclient.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}
