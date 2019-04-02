package db

import (
	"database/sql"
	"sync"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var once sync.Once

func GetDataSource() *sql.DB {
	var err error
	once.Do(func() {
		db, err = sql.Open("mysql", "root:wangshubo@/test?charset=utf8")
	})
	if err != nil {
		fmt.Println("get db error")
	}
	defer db.Close()
	return db
}
