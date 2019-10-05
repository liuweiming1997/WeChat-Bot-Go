package main

import (
	"github.com/WeChat-Bot-Go/wechat"
	"github.com/WeChat-Bot-Go/wechat/login"
)

func main() {
	wxService := &wechat.WxService{
		LoginService: &login.WxLogin{},
	}
	wxService.ServerForever()
}
