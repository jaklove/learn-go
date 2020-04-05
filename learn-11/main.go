package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`
}

type server struct {
	ServerName string `xml:"ServerName"`
	ServerIP string `xml:"ServerIP"`
}

func main()  {
	v := &Servers{Version:"1"}
	v.Svs = append(v.Svs,server{"shanghai","127.0.0.1"})
	v.Svs = append(v.Svs,server{"beijing","127.0.0.2"})
	fmt.Println(v.Svs)
	bytes, e := xml.MarshalIndent(v, " ", "  ")
	if e != nil{
		fmt.Println(e.Error)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(bytes)
}