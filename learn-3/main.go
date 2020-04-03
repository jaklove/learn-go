package main

import (
	"fmt"
	"github.com/gorilla/schema"
	"html/template"
	"log"
	"net/http"
	"os"
)

var decoder = schema.NewDecoder()


type Login struct {
	Username string
	Passwd string
}

func SayHelloGame(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	var login Login
	err := decoder.Decode(&login, r.PostForm)
	if err != nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	fmt.Println("username",login.Username)
	fmt.Println("passwd",login.Username)

}

func login(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("获取请求的方法",r.Method)

	wd, err := os.Getwd()
	if err != nil{
		log.Fatal("获取路径错误:",err.Error())
	}
	fmt.Println(wd)

	if r.Method == "GET"{
		files, e := template.ParseFiles(wd+"/learn-3/login.html")
		if e != nil{
			log.Fatal("请求方法错误:",e.Error())
		}
		files.Execute(w,nil)
	}else {
	}

}

func main()  {
   http.HandleFunc("/login",SayHelloGame)
   http.HandleFunc("/detail",login)
   err := http.ListenAndServe(":8080",nil)
   if err != nil{
   	  log.Fatal("监听错误")
   }
}
