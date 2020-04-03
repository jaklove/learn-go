package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func upload(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method:",r.Method)
	if r.Method == "GET"{
		crutime := time.Now().Unix()
		h := md5.New()
		_, err := io.WriteString(h, strconv.FormatInt(crutime, 10))
		if err != nil{
			log.Fatal("io.WriteString err:",err.Error())
		}

		wd, err := os.Getwd()
		if err != nil{
			log.Fatal("获取路径错误:",err.Error())
		}

		files, err := template.ParseFiles(wd + "/learn-5/upload.html")
		if err != nil{
			log.Fatal("template err:",err.Error())
		}

		token := fmt.Sprintf("%x", h.Sum(nil))
		fmt.Println("crate token :",token)
		files.Execute(w,token)
	}else {
		r.ParseMultipartForm(128)
		file, header, e := r.FormFile("uploadfile")
		if e != nil{
			fmt.Println(e.Error())
			return
		}

		defer file.Close()
		fmt.Fprintf(w,"%v",header.Header)

		wd, err := os.Getwd()
		if err != nil{
			log.Fatal("获取路径错误:",err.Error())
		}

		openFile, e := os.OpenFile(wd+"/learn-5/"+header.Filename, os.O_WRONLY|os.O_CREATE,0666)
		if e != nil{
			log.Fatal(e.Error())
			return
		}
		defer openFile.Close()
		io.Copy(openFile,file)
	}


}

func main()  {
	http.HandleFunc("/upload",upload)
	err := http.ListenAndServe(":8080",nil)
	if err != nil{
		log.Fatal("监听错误")
	}

}
