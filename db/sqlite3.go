package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//打开数据库
	db, err := sql.Open("sqlite3", "E:\\sqlite\\test.db")
	checkErr(err)

	//插入记录
	stmt, err := db.Prepare("insert into userinfo(username,departname,created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("shine", "研发一部", "2017-10-13")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("插入的最后一条记录Id为：", id)

	//更新记录
	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err := stmt.Exec("shine.w", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("受影响的记录：", affect)

	//查询记录
	rows, err := db.Exec("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created time.Time

		err = rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)
		fmt.Print("uid:%d - username:%s - departname:%s created:%s", uid, username, departname, created)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
