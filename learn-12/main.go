package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
)

type Server struct {
	ServerName string
	ServerIP string
}

type Serverslice struct {
	Servers []Server
}

func main()  {
	//var s Serverslice
	//str := `{"servers":[{"ServerName":"shanghai","ServerIP":"127.0.0.1"},{"ServerName":"beijing","ServerIP":"127.0.0.2"}]}`
	//err := json.Unmarshal([]byte(str), &s)
	//if err != nil{
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(s)

	//TestSimpleJson()

	CreateJson()
}

func TestSimpleJson()  {
	js ,err := simplejson.NewJson([]byte(`{
       "test":{
           "array":[1,"2",3],
           "int" :10,
			"float":5.150,
           "bignum":92332312131313132,
           "string":"simplejson", 
           "bool":true
        }
        }`))


	if err != nil{
		fmt.Println(err.Error())
	}

	arr,_ := js.Get("test").Get("array").Array()
	fmt.Println("array:",arr)
	i,_ := js.Get("test").Get("int").Int()
	fmt.Println("i:",i)
	m := js.Get("test").Get("string").MustString()
	fmt.Println("m:",m)
}

func CreateJson(){
	var s Serverslice
	s.Servers = append(s.Servers,Server{ServerName:"shanghai",ServerIP:"127.0.0.1"})
	s.Servers = append(s.Servers,Server{ServerName:"beijing",ServerIP:"127.0.0.2"})
	bytes, e := json.Marshal(s)
	if e != nil{
		fmt.Println(e.Error())
	}
	fmt.Println(string(bytes))
}
