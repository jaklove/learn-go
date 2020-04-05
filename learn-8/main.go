package main

import (
	"fmt"
	"github.com/astaxie/goredis"
)

func main()  {
	var client goredis.Client


	////字符串操作
	//client.Set("a",[]byte("hello"))
	//bytes, e := client.Get("a")
	//if e != nil{
	//	log.Fatal(e.Error())
	//}
	//fmt.Println(string(bytes))

	//list操作
	vals := []string{"a","b","c","d","e"}
	for i,v := range vals{
		fmt.Println(i)
		client.Rpush("l",[]byte(v))
	}

	dbvals,_ := client.Lrange("l",0,4)
	for i,v := range dbvals{
		fmt.Println(i,":",string(v))
	}
	client.Del("l")
}
