package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Auth(request *http.Request, writer http.ResponseWriter) {
	decoder := json.NewDecoder(request.Body)
	var params map[string]interface{}
	decoder.Decode(&params)

	nickName := params["nickName"]
	watermark := params["watermark"]
	watermarkMap := watermark.(map[string]string)
	timestamp := watermarkMap["timestamp"]
	fmt.Printf("POST json: username=%s, password=%s\n", nickName, timestamp)

	fmt.Fprintf(writer, `{"code":0}`)
}
