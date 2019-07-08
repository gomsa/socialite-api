package main

import (
	"github.com/gomsa/tools/config"
)

// Conf 配置
var Conf config.Config = config.Config{
	Service: "socialite-api",
	Version: "latest",
	Permissions: []config.Permission{
		{Service: "socialite-api", Method: "Miniprogram.Auth", Auth: false, Policy: false, Name: "小程序授权", Description: "小程序授权通过 code 换取 token"},
		{Service: "socialite-api", Method: "Users.Build", Auth: true, Policy: false, Name: "查询绑定", Description: "查询第三方登录绑定"},
		{Service: "socialite-api", Method: "Users.Delete", Auth: true, Policy: false, Name: "解除绑定", Description: "解除第三方登录绑定"},
	},
}
