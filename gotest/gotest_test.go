package gotest

import "testing"

func TestString(t *testing.T) {
	path := "/cheng"
	path = string([]byte(path)[1:len(path)])
	t.Log(path)
}
