package main

// http://www.cnblogs.com/txw1958/p/weixin-qrcode-with-parameters.html

import (
	"fmt"
	"go-web/task"
	"go-web/web"
	"log"
	"net/http"
	"strings"
)

func dispatcher(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		if _, err = fmt.Fprintf(w, "server error"); err == nil {
			fmt.Println("error")
		}
		fmt.Println("error")
		return
	}

	//var response interface{}
	path := r.URL.Path
	path = string([]byte(path)[1:len(path)])
	if strings.EqualFold(path, "WeChatEvent") {
		web.WeChatEvent(r, w)
	} else if strings.EqualFold(path, "QRCodeTicket") {
		web.QRCodeTicket(r, w)
	} else if strings.EqualFold(path, "CheckLoginState") {
		web.CheckLoginState(r, w)
	} else if strings.EqualFold(path, "Login") {
		web.Login(r, w)
	} else if strings.EqualFold(path, "Sub") {
		web.Subscription(r, w)
	}

	// TODO 返回Josn

	/*switch resp := response.(type) {
	case nil:
		fmt.Print("response is nil")
	case string:
		fmt.Fprintf(w, resp)
	default:
		responseByte, jsonErr := json.Marshal(response)
		if jsonErr != nil {
			fmt.Println("json marshal error")
		}
		fmt.Fprintf(w, string(responseByte))
	}*/

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

	go task.StartAccessTokenTask()

	//time.Sleep(time.Second * 1)
	http.HandleFunc("/", dispatcher)
	requestErr := http.ListenAndServe(":9090", nil)
	if requestErr != nil {
		log.Fatal("error-->", requestErr)
	}

}
