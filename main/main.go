package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

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
	http.HandleFunc("/", Dipatcher)

	requestErr := http.ListenAndServe(":9090", nil)
	if requestErr != nil {
		log.Fatal("error-->", err)

	}

}
