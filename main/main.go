package main

import (
	"encoding/json"
	"fmt"
	"go-web/cache"
	"go-web/web"
	"log"
	"net/http"
	"strings"
)

var globalCache *cache.Cache

func dispatcher(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		if _, err = fmt.Fprintf(w, "server error"); err == nil {
			fmt.Println("error")
		}
		fmt.Println("error")
		return
	}

	var response string
	path := r.URL.Path
	path = string([]byte(path)[1:len(path)])
	if strings.EqualFold(path, "QRCodeTicket") {
		response = web.QRCodeTicket(r.Form, globalCache)
	}
	var responseByte = []byte("")
	var jsonErr error

	if response != "" {
		responseByte, jsonErr = json.Marshal(response)
		if jsonErr != nil {
			fmt.Println("json marshal error")
		}
	}
	fmt.Fprintf(w, string(responseByte))

	/*fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["name"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")*/
}

func main() {
	globalCache = cache.New(10)

	//go task.StartAccessTokenTask(globalCache)

	globalCache.Put("1", "1-1-2")

	//time.Sleep(time.Second * 1)
	http.HandleFunc("/", dispatcher)
	requestErr := http.ListenAndServe(":9090", nil)
	if requestErr != nil {
		log.Fatal("error-->", requestErr)
	}

}
