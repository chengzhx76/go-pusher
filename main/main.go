package main

import (
	"fmt"
	"go-web/cache"
	"go-web/task"
	"log"
	"net/http"
	"strings"
)

func Dispatcher(w http.ResponseWriter, r *http.Request) {
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

func main() {
	task.StartAccessTokenTask()
	//time.Sleep(time.Second * 1)
	http.HandleFunc("/", Dispatcher)
	requestErr := http.ListenAndServe(":9090", nil)
	if requestErr != nil {
		log.Fatal("error-->", requestErr)

	}

	c1 := cache.New(10)
	c1.Put("1", "2")
	c2 := cache.New(10)
	if value, ok := c2.Get("1"); ok {
		fmt.Printf(value)
	}

}
