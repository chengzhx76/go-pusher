package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var accessToken string

// https://blog.csdn.net/u012210379/article/details/52795296
func Dipatcher(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["name"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello astaxie!")

}

// https://github.com/silenceper/wechat/blob/master/context/access_token.go
func requestWxAccessToken() {
	accessToken := func() {
		resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx52ddb78878fa6d98&secret=44af2777f136af01accabc96bc78d9cc")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))

		wxAccessToken := &mode.WxAccessToken{}

	}
	go accessToken()
}

func accessTokenTask() {
	task := time.NewTimer(2 * time.Second)
	for {
		select {
		case <-task.C:
			requestWxAccessToken()
			task.Reset(5 * time.Second)
		default:
			//fmt.Println("-----default---")
		}
	}
}

type user struct {
	Name string
	Sex  string
	Age  int8
}

func main() {

	zhangsan := &user{
		Name: "zhansan",
		Sex:  "1",
		Age:  12,
	}

	rs, err := json.Marshal(zhangsan)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rs)
	fmt.Println(string(rs))
	fmt.Println([]byte("{\"Name\":\"zhansan\",\"Age\":12}"))
	zhangsan1 := &user{}

	json.Unmarshal(rs, zhangsan1)

	fmt.Println(zhangsan1.Name)

	fmt.Println("-------------------------------")

	go accessTokenTask()

	http.HandleFunc("/", Dipatcher)

	requestErr := http.ListenAndServe(":9090", nil)
	if requestErr != nil {
		log.Fatal("error-->", err)

	}

}
