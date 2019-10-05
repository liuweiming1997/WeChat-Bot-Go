package wechat

func (wxService *WxService) ServerForever() {
	wxService.LoginService.Login()
	// wxService.InitService.Init()
}
