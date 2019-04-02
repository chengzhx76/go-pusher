package gotest

import (
	"fmt"
	"go-web/cache"
	"go-web/util"
	"testing"
)

func TestString(t *testing.T) {
	path := "/cheng"
	path = string([]byte(path)[1:len(path)])
	t.Log(path)
}
func TestCache(t *testing.T) {
	cache.GetInstance().Put("fpllngzieyoh54e1", 1)
	value, _ := cache.GetInstance().Get("fpllngzieyoh54e1")
	fmt.Println(value)
}
func TestRandom(t *testing.T) {

	fmt.Println(util.RandString(16))
}
