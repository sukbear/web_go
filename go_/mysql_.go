package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:mysql@/hrdb")
	defer db.Close()
	checkErr(err)
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	res, err := stmt.Exec("astaxie", "研发", "2012-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
/*	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)*/

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
