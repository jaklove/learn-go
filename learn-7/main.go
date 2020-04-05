package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
	"log"
	"time"
)

func main()  {
	db, e := sql.Open("mymysql", "learn-go/root/sinowealth")
	if e != nil{
		log.Fatal(e.Error())
	}

	orm := beedb.New(db)

	//var saveone Userinfo
	//saveone.Username = "TEST ADD USER"
	//saveone.Departname = "TEST ADD Departname"
	//saveone.Created = time.Now()
	//save := orm.Save(&saveone)
	//if save != nil{
	//	log.Fatal(save.Error())
	//}
	//fmt.Println("insert info success!")


	////map插入
	//add := make(map[string]interface{})
	//add["username"] = "map user"
	//add["departname"] = "map departname"
	//add["created"] = "2020-04-04"
	//
	//id, err := orm.SetTable("userinfo").Insert(add)
	//if err != nil{
	//	log.Fatal("插入失败",err.Error())
	//}
	//fmt.Println("插入成功id:",id)


	//插入多条数据
	//addslice := make([]map[string]interface{},0)
	//add := make(map[string]interface{})
	//add2 := make(map[string]interface{})
	//
	//add["username"] = "username1"
	//add["departname"] = "departname1"
	//add["created"] = "2020-04-04"
	//
	//add2["username"] = "username2"
	//add2["departname"] = "departname2"
	//add2["created"] = "2020-04-04"
	//
	//addslice = append(addslice,add,add2)
	//int64s, err := orm.SetTable("userinfo").InsertBatch(addslice)
	//if err != nil{
	//	log.Fatal("插入失败",err.Error())
	//}
	//fmt.Println("插入成功id数组:",int64s)


	////更新操作
	//var saveone Userinfo
	//saveone.Uid = 1
	//saveone.Username = "user"
	//saveone.Departname = "departname"
	//saveone.Created = time.Now()
	//saveErr := orm.Save(&saveone)
	//if saveErr != nil{
	//	log.Fatal("更新错误")
	//}

	//更新支持map操作
	t := make(map[string]interface{})
	t["username"] = "zhourenjie"
	i, err := orm.SetTable("userinfo").SetPK("uid").Where(2).Update(t)
	if err != nil{
		log.Fatal(err.Error())
	}
	fmt.Println("影响行数:",i)

}

type Userinfo struct {
	Uid int `PK`
	Username string
	Departname string
	Created time.Time
}
