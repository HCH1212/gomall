package rpc

import (
	"github.com/HCH1212/gomall/app/frontend/conf"
	frontendUtils "github.com/HCH1212/gomall/app/frontend/utils"
	"github.com/HCH1212/gomall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	UserClient userservice.Client
	once       sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegisterAddr)
	frontendUtils.MustHandlerError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandlerError(err)
}
