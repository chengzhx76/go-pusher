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
	err := r.ParseForm()
	if err != nil {
		_, wErr := fmt.Fprintf(w, "server error")
		if wErr == nil {
			fmt.Println("error")
		}
		fmt.Println("error")
		return
	}

	var response string
	path := r.URL.Path
	if strings.EqualFold(path, "/QRCodeTicket") {
		response = web.QRCodeTicket(r.Form, globalCache)
	}
	res, _ := json.Marshal(response)
	fmt.Fprintf(w, string(res))

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
