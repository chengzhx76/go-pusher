package web

import (
	"fmt"
	"go-web/cache"
	"io/ioutil"
	"net/http"
	"net/url"
)

func QRCodeTicket(params url.Values, globalCache *cache.Cache) string {
	accessToken, ok := globalCache.Get("ACCESS_TOKEN")

	resp, err := http.Post("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + accessToken)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	if ok {
		fmt.Print("------web.QRCodeTicket-------")
		fmt.Print(accessToken)
	}
	return "method-QRCodeTicket"
}
