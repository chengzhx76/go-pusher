package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"go-web/mode"
	"go-web/cache"
)

// https://github.com/silenceper/wechat/blob/master/context/access_token.go
func requestWxAccessToken() {
	accessToken := func() {
		resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wx52ddb78878fa6d98&secret=44af2777f136af01accabc96bc78d9cc")
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))

		wxAccessToken := &mode.WxAccessToken{}
		if err = json.Unmarshal(body, wxAccessToken); err != nil {
			accessToken = wxAccessToken.AccessToken
			cache0 := cache.New(10)
			cache0.Remove()
			cache1 := &cache.Cache{}
			cache1.
			cache2 := new(cache.Cache)
			cache3 := make(cache.Cache)

		}
	}
	go accessToken()
}

func accessTokenTask() {
	task := time.NewTimer(5 * time.Second)
	for {
		select {
		case <-task.C:
			requestWxAccessToken()
			task.Reset(5 * time.Second)
			fmt.Println("======================================================")
		}
	}
}

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
