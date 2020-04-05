package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Recurlyservers struct {
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`
	Description string `xml:",innerxml"`
}

type server struct {
	XMLName xml.Name `xml:"server"`
	ServerName string `xml:"ServerName"`
	ServerIP string `xml:"ServerIP"`
}

func main()  {
	wd, err := os.Getwd()
	if err != nil{
		log.Fatal("获取路径错误:",err.Error())
	}

	file, e := os.Open(wd+"/learn-10/server.xml")
	if e != nil{
		fmt.Println("error :",e.Error())
		return
	}
	defer file.Close()

	bytes, e := ioutil.ReadAll(file)
	if e != nil{
		fmt.Println("read file err:",e.Error)
		return
	}

	v := Recurlyservers{}
	xmlError := xml.Unmarshal(bytes, &v)
	if xmlError != nil{
		fmt.Println("error:",xmlError.Error)
	}
	fmt.Println(v)
}


