package login

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/WeChat-Bot-Go/lib"
	"github.com/WeChat-Bot-Go/logger"
	"github.com/parnurzeal/gorequest"
)

func (wxLogin *WxLogin) FetchUuid() string {
	params := Uuid{
		Appid:       "wx782c26e4c19acffb",
		RedirectUri: "https://wx2.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage",
		Fun:         "new",
		Lang:        "zh_CN",
		TimeStampMs: wxLogin.TimeStampMs,
	}
	resp, body, errs := gorequest.New().Get(UUID_URL).Query(params).End()

	if len(errs) > 0 || resp.StatusCode != 200 {
		logger.Fatal(resp, errs)
		return ""
	}
	return body[50:62]
}

func (wxLogin *WxLogin) FetchQRCode() bool {
	resp, _, errs := gorequest.New().Get(QR_CODE_URL + wxLogin.Uuid).Query(`{"t": "webwx"}`).End()
	if len(errs) > 0 || resp.StatusCode != 200 {
		logger.Fatal(resp, errs)
		return false
	}
	logger.Info("The QR Code Url is ", resp.Request.URL)
	lib.Open(fmt.Sprintf("%s", resp.Request.URL))
	return true
}

func (wxLogin *WxLogin) Login() {
	logger.Info("Start login")
	wxLogin.TimeStampMs = time.Now().UnixNano() / 1000
	// First get uuid
	wxLogin.Uuid = wxLogin.FetchUuid()
	// Second get the QR Code
	if !wxLogin.FetchQRCode() {
		logger.Fatal("Fetch QR code error")
	}
	// Third polling user status [qrcode weather scaned]
	userIsScanedQRCode := make(chan bool)
	go func() {
		ticker := time.NewTicker(time.Duration(SCAN_QR_CODE_POLLING_INTERVAL) * time.Second)
		for pollingTime := range ticker.C {
			msg := fmt.Sprintf("%s %s", pollingTime, "Polling user scan code status")
			params := Waitting{
				LoginIcon:   true,
				Uuid:        wxLogin.Uuid,
				Tip:         0,
				R:           time.Now().UnixNano() / 1000,
				TimeStampMs: wxLogin.TimeStampMs,
			}
			resp, body, errs := gorequest.New().Get(WAITTINE_URL).Query(params).End()
			if len(errs) > 0 || resp.StatusCode != 200 {
				logger.Error(resp, errs)
				continue
			}
			logger.Info(fmt.Sprintf("\nmsg: %s \nURL: %s \nUser Status: %s", msg, resp.Request.URL, body))
			if strings.Contains(body, "redirect_uri") && strings.Contains(body, "200") {
				wxLogin.RedirectUri = body[38 : len(body)-2]
				ticker.Stop()
				userIsScanedQRCode <- true
				break
			}
		}
	}()
	select {
	case <-userIsScanedQRCode:
		logger.Info("User Login Success")
		logger.Info(wxLogin.RedirectUri)
		resp, _, errs := gorequest.
			New().
			Get(wxLogin.RedirectUri).
			RedirectPolicy(
				func(req gorequest.Request, via []gorequest.Request) error {
					// NOTE should not redirect here.
					return http.ErrUseLastResponse
				}).
			End()
		if resp.StatusCode != 301 {
			logger.Fatal(resp, errs)
		}
		fmt.Println(resp.Status)
		fmt.Println(resp.Request.URL)
		fmt.Println(resp.Header)
	case <-time.After(time.Second * SCAN_QR_CODE_TIMEOUT):
		logger.Fatal("Scan QR Code timeout")
	}
}
