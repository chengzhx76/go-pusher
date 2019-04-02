package web

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"go-web/cache"
	"go-web/mode"
	"go-web/util"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func WeChatEvent(r *http.Request, w http.ResponseWriter) {

	fmt.Println("------------BODY--------------")
	wxMsg := mode.CommonWxMsg{}
	result, _ := ioutil.ReadAll(r.Body)
	xmlData := bytes.NewBuffer(result).String()
	fmt.Println(xmlData)
	xml.Unmarshal(result, &wxMsg)
	if wxMsg.MsgType == "event" {
		wxEventMsg := mode.WxScanEvent{}
		xml.Unmarshal(result, &wxEventMsg)
		if wxEventMsg.Event == "SCAN" {
			randomStr := wxEventMsg.EventKey
			cache.GetInstance().Put(randomStr, 1)

			fmt.Println(randomStr)
		}
	} else if wxMsg.MsgType == "text" {
		fmt.Println("text")
	}

	fmt.Println("------------GET--------------")
	for k, v := range r.Form {
		fmt.Print(k + ":" + strings.Join(v, ""))
		fmt.Println()
	}

	fmt.Println("-----------POST---------------")
	for k, v := range r.PostForm {
		fmt.Print(k + ":" + strings.Join(v, ""))
		fmt.Println()
	}

	fmt.Fprint(w, "SUCCESS")
}

func QRCodeTicket(r *http.Request, w http.ResponseWriter) {
	fmt.Fprint(w, getTicket())
}
func CheckLoginState(r *http.Request, w http.ResponseWriter) {
	loginToken := r.Form["loginToken"][0]
	if loginState, ok := cache.GetInstance().Get(loginToken); ok {

		fmt.Fprint(w, loginState.(int))
	} else {
		fmt.Fprint(w, 0)
	}
}

// http://polyglot.ninja/golang-making-http-requests/
func getTicket() string {
	// {"expire_seconds": 604800, "action_name": "QR_STR_SCENE", "action_info": {"scene": {"scene_str": "test"}}}
	randomStr := util.RandString(16)
	cache.GetInstance().Put(randomStr, 0)
	fmt.Println(randomStr)
	data := map[string]interface{}{
		"expire_seconds": 604800,
		"action_name":    "QR_STR_SCENE",
		"action_info": map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_str": randomStr,
			},
		},
	}
	accessToken, _ := cache.GetInstance().Get("ACCESS_TOKEN")
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

	return string(result["ticket"].(string)) + "|" + randomStr
}
