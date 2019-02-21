package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:nopass.123@tcp(192.168.145.5:3306)/crowdfunding")
	checkErr(err)
	defer db.Close()

	err = db.Ping()
	checkErr(err)

	stmtOut, err := db.Prepare("select name from fleet where id = ?")
	checkErr(err)
	defer stmtOut.Close()

	var name string
	err = stmtOut.QueryRow(1).Scan(&name)
	checkErr(err)
	fmt.Printf("id为1的车队名字叫：%v \n", name)
	fmt.Println("获取多个值")

	rows, _ := db.Query("select id,name from fleet")
	id := 0
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}

	fmt.Println("获取全部数值")
	rows2, _ := db.Query("select * from fleet")
	cols, _ := rows2.Columns()              //所有列
	vals := make([][]byte, len(cols))       // 列的值
	scans := make([]interface{}, len(cols)) // 一行记录用interface{} 记录
	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	i := 0
	result := make(map[int]map[string]string)
	for rows2.Next() {
		rows2.Scan(scans...)

		row := make(map[string]string)
		for k, v := range vals {
			key := cols[k]
			row[key] = string(v)
		}
		result[i] = row
		i++
	}
	fmt.Println(result)

	fmt.Println("事务处理")
	tx, _ := db.Begin()
	fmt.Println("当前时间：", time.Now().Add(8*time.Hour))
	// ret1, _ := tx.Exec("update fleet set modify_time=? where id=1", time.Now())//这样会是0时区时间，需要加8小时
	ret1, _ := tx.Exec("update fleet set modify_time=now() where id=1")
	if num, _ := ret1.RowsAffected(); num > 0 {
		fmt.Println("事务提交")
		tx.Commit()
	} else {
		fmt.Println("事务回滚")
		tx.Rollback()
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
