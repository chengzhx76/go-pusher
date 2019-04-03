package gotest

import (
	"testing"
	"go-web/db"
	"database/sql"
	"fmt"
)

func TestSaveUser(t *testing.T) {
	db.SaveUser("123", "456", "789")
}
func TestDb(t *testing.T) {
	db, _ := sql.Open("mysql", "pusher:pusher#cheng@tcp(139.196.35.134:3306)/pusher?charset=utf8")
	stmt, _ := db.Prepare("INSERT INTO user(uid, sendKey, openid) VALUES (?,?,?)")

	res, _ := stmt.Exec("123", "456", "789")
	fmt.Println(res.LastInsertId())
}