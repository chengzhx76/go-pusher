package web

import (
	"fmt"
	"go-web/cache"
	"net/url"
)

func QRCodeTicket(params url.Values, globalCache *cache.Cache) string {
	value, ok := globalCache.Get("1")
	if ok {
		fmt.Print("------web.QRCodeTicket-------")
		fmt.Print(value)
	}
	return "method-QRCodeTicket"
}
