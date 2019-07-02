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
	},
}
