package web

import (
	"fmt"
	"go-web/cache"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
	"log"
	"bytes"
)

func weChatEvent(params url.Values, globalCache *cache.Cache) string {
	return getTicket(globalCache)
}
func QRCodeTicket(params url.Values, globalCache *cache.Cache) string {
	return getTicket(globalCache)
}

// http://polyglot.ninja/golang-making-http-requests/
func getTicket(globalCache *cache.Cache) string {
	// {"expire_seconds": 604800, "action_name": "QR_STR_SCENE", "action_info": {"scene": {"scene_str": "test"}}}
	data := map[string]interface{}{
		"expire_seconds": 604800,
		"action_name": "QR_STR_SCENE",
		"action_info": map[string]interface{}{
			"scene":  map[string]interface{}{
				"scene_id": 123,
			},
		},
	}
	accessToken, _ := globalCache.Get("ACCESS_TOKEN")
	var wxUrl = "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + accessToken.(string)
	byteData, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(wxUrl, "application/json", bytes.NewBuffer(byteData))
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)
	log.Println(result["data"])

	return string(result["ticket"])
}
