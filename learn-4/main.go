package main

import (
	"crypto/md5"
	"fmt"
	"github.com/gorilla/schema"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
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
		 crutime := time.Now().Unix()
		 hash := md5.New() //new出一个hash对象
		 _, err := io.WriteString(hash, strconv.FormatInt(crutime, 10))
		 if err != nil{
		 	log.Fatal("io.WriteString err:",err.Error())
		 }

		wd, err := os.Getwd()
		if err != nil{
			log.Fatal("获取路径错误:",err.Error())
		}

		files, err := template.ParseFiles(wd + "/learn-4/login.html")
		if err != nil{
			log.Fatal("template err:",err.Error())
		}

		token := fmt.Sprintf("%x", hash.Sum(nil))

		fmt.Println("crate token :",token)
		files.Execute(w,token)
	}else {
		//登陆
		r.ParseForm()
		token := r.PostForm.Get("token")
		if token != ""{
			fmt.Println("验证token",token)
		}else{
			fmt.Fprintf(w,"token不存在")
		}


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
