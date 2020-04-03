package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var dbInstacne *sql.DB

func init()  {
	db, e := sql.Open("mysql", "root:sinowealth@(127.0.0.1:3306)/learn-go?charset=utf8mb4")
	checkErr(e)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	dbInstacne = db
}

func main()  {
	//db, e := sql.Open("mysql", "root:sinowealth@(127.0.0.1:3306)/learn-go?charset=utf8mb4")
	//checkErr(e)
	//stmt, e := db.Prepare("insert userinfo set username = ?,departname = ?,created = ?")
	//checkErr(e)
	//
	//result, e := stmt.Exec("zhourenjie", "研发中心", "2020-04-03")
	//checkErr(e)
	//
	//i, e := result.LastInsertId()
	//checkErr(e)
	//fmt.Println("insert id:",i)
	//
	////更新数据
	//pre, e := db.Prepare("update userinfo set username=? where uid = ?")
	//checkErr(e)
	//
	//exec, e := pre.Exec("faker", i)
	//checkErr(e)
	//
	//affected, e := exec.RowsAffected()
	//checkErr(e)
	//fmt.Println(affected)

	//查询数据
	//rows, e := dbInstacne.Query("select * from userinfo")
	//checkErr(e)
	//for rows.Next() {
	//	var (
	//		uid int
	//		username string
	//		departname string
	//		created string
	//	)
	//	e := rows.Scan(&uid, &username, &departname, &created)
	//	checkErr(e)
	//	fmt.Println(uid)
	//	fmt.Println(username)
	//	fmt.Println(departname)
	//	fmt.Println(created)
	//}

	////删除数据
	//stmt, e := dbInstacne.Prepare("delete from userinfo where uid =?")
	//checkErr(e)
	//result, e := stmt.Exec(1)
	//checkErr(e)
	//fmt.Println("删除的行",result)

	//压测
	for i := 0; i < 100; i++ {
		go func(i int) {
			mSql := "select * from userinfo"
			rows, _ := dbInstacne.Query(mSql)
			rows.Close() //这里如果不释放连接到池里，执行5次后其他并发就会阻塞
			fmt.Println("第 ", i)
		}(i)
	}

	for {
		time.Sleep(time.Second)
	}

	defer dbInstacne.Close()
}

func checkErr(err error)  {
	if err != nil{
		panic(err)
	}
}