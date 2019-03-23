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

var accessToken = ""

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

// http://xiaorui.cc/2016/03/06/%E5%85%B3%E4%BA%8Egolang-timer%E5%AE%9A%E6%97%B6%E5%99%A8%E7%9A%84%E8%AF%A6%E7%BB%86%E7%94%A8%E6%B3%95/
func requestWxAccessToken() {
	go func() {
		resp, err := http.Get("https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=ACCESS_TOKEN")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		s, err := ioutil.ReadAll(resp.Body)
		fmt.Printf(string(s))

		accessToken = string(s)
	}()
}

func accessTokenTask() {
	task := time.NewTicker(2 * time.Second)
	select {
	case <-task.C:
		requestWxAccessToken()
	}
}

type user struct {
	Name string
	sex  string
	Age  int8
}

func main() {

	zhangsan := user{
		Name: "zhansan",
		sex:  "1",
		Age:  12,
	}

	rs, err := json.Marshal(zhangsan)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rs)
	fmt.Println(string(rs))

	fmt.Println("-------------------------------")

	go accessTokenTask()

	http.HandleFunc("/", Dipatcher)

	requestErr := http.ListenAndServe(":9090", nil)
	if requestErr != nil {
		log.Fatal("error-->", err)

	}

}
