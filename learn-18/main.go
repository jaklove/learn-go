package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func main()  {

	var salt = "username"
	bytes, e := scrypt.Key([]byte("some password"),[]byte(salt), 32768, 8, 1, 32)
	if e != nil{
		fmt.Println(e.Error())
	}
	fmt.Println(string(bytes))

	//base64加解密
	hello := "你好，世界! hello word"
	debyte := base64Encode([]byte(hello))
	fmt.Println(debyte)
	//decode
	decode, err := base64Decode(debyte)
	if err != nil{
		fmt.Println(err.Error())
	}

	if hello != string(decode){
		fmt.Println("hello is not equal to enbyte")
	}

	fmt.Println(string(decode))

}
func base64Encode(src []byte)[]byte  {
	return  []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte)([]byte ,error){
	return  base64.StdEncoding.DecodeString(string(src))
}


