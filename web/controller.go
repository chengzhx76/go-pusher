package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-web/cache"
	"log"
	"net/http"
	"strings"
	"io/ioutil"
)

func WeChatEvent(r *http.Request, w http.ResponseWriter) {

	fmt.Println("------------BODY--------------")
	result, _ := ioutil.ReadAll(r.Body)
	fmt.Println(bytes.NewBuffer(result).String())

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

	fmt.Fprint(w, "")
}

func QRCodeTicket(r *http.Request, w http.ResponseWriter) {
	fmt.Fprint(w, getTicket())
}

// http://polyglot.ninja/golang-making-http-requests/
func getTicket() string {
	// {"expire_seconds": 604800, "action_name": "QR_STR_SCENE", "action_info": {"scene": {"scene_str": "test"}}}
	data := map[string]interface{}{
		"expire_seconds": 604800,
		"action_name":    "QR_STR_SCENE",
		"action_info": map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_str": "cheng",
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

	return string(result["ticket"].(string))
}
