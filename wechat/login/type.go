package login

const (
	/*
		https://login.wx2.qq.com/jslogin?appid=wx782c26e4c19acffb&redirect_uri=https%3A%2F%2Fwx2.qq.com%2Fcgi-bin%2Fmmwebwx-bin%2Fwebwxnewloginpage&fun=new&lang=zh_CN&_=1570239590655
		method: [GET]
		   appid: wx782c26e4c19acffb
		   redirect_uri: https://wx2.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage
		   fun: new
		   lang: zh_CN
		   _: 1570239590655
	*/
	UUID_URL = `https://login.wx2.qq.com/jslogin`
	/*
		https://login.weixin.qq.com/qrcode/4cyZTXOC9w==
		method: [GET]
	*/
	QR_CODE_URL = `https://login.weixin.qq.com/qrcode/`
	/*
		https://login.wx2.qq.com/cgi-bin/mmwebwx-bin/login?loginicon=true&uuid=4cyZTXOC9w==&tip=0&r=1716389794&_=1570241589766
		method: [GET]
			loginicon: true
			uuid: 4cyZTXOC9w==
			tip: 0
			r: 1716389794
			_: 1570241589766
	*/
	WAITTINE_URL = `https://login.wx2.qq.com/cgi-bin/mmwebwx-bin/login`

	SCAN_QR_CODE_TIMEOUT          = 60 * 5 // in second
	SCAN_QR_CODE_POLLING_INTERVAL = 2      // in second
)

type WxLogin struct {
	Uuid        string
	TimeStampMs int64
	redirectUri string
}

type Uuid struct {
	Appid       string `json:"appid"`
	RedirectUri string `json:"redirect_uri"`
	Fun         string `json:"fun"`
	Lang        string `json:"lang"`
	TimeStampMs int64  `json:"_"`
}

type Waitting struct {
	LoginIcon   bool   `json:"loginicon"`
	Uuid        string `json:"uuid"`
	Tip         int    `json:"tip"`
	R           int    `json:"r"`
	TimeStampMs int64  `json:"_"`
}
