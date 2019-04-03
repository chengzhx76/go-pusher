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
	"go-web/db"
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
			sendKey := wxEventMsg.EventKey
			var uid string
			// 先查询有没有该用户
			user, _ := db.LoadByOpenid(wxEventMsg.FromUserName)
			if user == nil {
				// 入库
				uid = util.RandString(16)
				db.SaveUser(uid, sendKey, wxEventMsg.FromUserName)
			} else {
				uid = user.Uid
			}
			cache.GetInstance().Put(sendKey, uid)
			fmt.Println(sendKey)
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
	if uid, ok := cache.GetInstance().Get(loginToken); ok {
		cache.GetInstance().Remove(loginToken)
		fmt.Fprint(w, uid.(string))
	} else {
		fmt.Fprint(w, "")
	}
}
func Subscription(r *http.Request, w http.ResponseWriter) {
	sendkey := r.Form["sendkey"][0]
	text := r.Form["text"][0]
	desp := r.Form["desp"][0]
	fmt.Println(sendkey, text, desp)

	fmt.Fprint(w, sendkey, text, desp)

}

func Login(r *http.Request, w http.ResponseWriter) {
	uid := r.Form["uid"][0]
	fmt.Println(uid)

	fmt.Fprint(w, uid)

}

// http://polyglot.ninja/golang-making-http-requests/
func getTicket() string {
	// {"expire_seconds": 604800, "action_name": "QR_STR_SCENE", "action_info": {"scene": {"scene_str": "test"}}}
	sendKey := util.RandString(16)
	cache.GetInstance().Put(sendKey, "")
	fmt.Println(sendKey)
	data := map[string]interface{}{
		"expire_seconds": 604800,
		"action_name":    "QR_STR_SCENE",
		"action_info": map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_str": sendKey,
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

	return string(result["ticket"].(string)) + "|" + sendKey
}
