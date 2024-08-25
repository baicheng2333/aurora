package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"sakura/aurora/api/internal/config"
	"sakura/aurora/rpc/user/client/authenticate"
)

type ServiceContext struct {
	Config       config.Config
	Authenticate authenticate.Authenticate
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		Authenticate: authenticate.NewAuthenticate(zrpc.MustNewClient(c.UserRpc)),
	}
}
