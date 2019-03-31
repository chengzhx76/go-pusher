package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

//https://blog.csdn.net/wangshubo1989/article/details/75257614

func main() {
	db, err := sql.Open("mysql", "root:wangshubo@/test?charset=utf8")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT user_info SET id=?,name=?")
	checkErr(err)

	res, err := stmt.Exec(1, "wangshubo")
	checkErr(err)

	// update
	stmt, err = db.Prepare("update user_info set name=? where id=?")
	checkErr(err)

	res, err = stmt.Exec("wangshubo_update", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("SELECT * FROM user_info")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string

		err = rows.Scan(&uid, &username)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
	}

	// delete
	stmt, err = db.Prepare("delete from user_info where id=?")
	checkErr(err)

	res, err = stmt.Exec(1)
	checkErr(err)

	// query
	rows, err = db.Query("SELECT * FROM user_info")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string

		err = rows.Scan(&uid, &username)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
	}

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
