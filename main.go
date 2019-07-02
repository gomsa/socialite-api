package main

import (
	// 公共引入

	"github.com/micro/go-micro/util/log"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"


	"github.com/gomsa/user/client"
	m "github.com/gomsa/user/middleware"

	"github.com/gomsa/socialite-api/hander"
	userPB "github.com/gomsa/socialite-api/proto/user"
	mpPB "github.com/gomsa/socialite-api/proto/miniprogram"

)

func main() {
	// 设置权限
	h := m.Handler{
		Permissions: Conf.Permissions,
	}
	srv := k8s.NewService(
		micro.Name(Conf.Service),
		micro.Version(Conf.Version),
		micro.WrapHandler(h.Wrapper), //验证权限
	)
	srv.Init()


	// 小程序服务实现
	mpPB.RegisterMiniprogramHandler(srv.Server(), &hander.Miniprogram{})
	userPB.RegisterUsersHandler(srv.Server(), &hander.User{})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	// 同步权限
	if err := client.SyncPermission(Conf.Permissions); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
