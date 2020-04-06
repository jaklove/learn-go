package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	//userFile := "learn-13/hello.txt"
	//file, e := os.Create(userFile)
	//if e != nil{
	//	fmt.Println(file)
	//	return
	//}
	//
	//defer file.Close()
	//
	//for i := 0;i < 10;i++ {
	//	file.WriteString("just a test!\r\n")
	//	file.Write([]byte("just a test!\r\n"))
	//}

	//读文件
	//ReadFile()

	//字符串
	ContainString()
}

func ReadFile()  {
	userFile := "learn-13/hello.txt"
	file, e := os.Open(userFile)
	if e != nil{
		fmt.Println(e.Error())
		return
	}
	defer file.Close()

	buf := make([]byte,1024)
	for  {
		n, err := file.Read(buf)

		if err != nil{
			fmt.Println(err.Error())
			return
		}
		if n == 0{
			break
		}
		os.Stdout.WriteString(string(buf[:n]))
	}
}

func ContainString()  {
	fmt.Println(strings.Contains("seafood","foo"))
	fmt.Println(strings.Contains("seafood","bar"))
	fmt.Println(strings.Contains("seafood",""))
	fmt.Println(strings.Contains("",""))

	s := []string{"foo","bar","baz"}
	fmt.Println(strings.Join(s,","))

	fmt.Println(strings.Index("chicken","ken"))
	fmt.Println(strings.Index("chicken","dmr"))

	fmt.Printf("%q\n",strings.Split("a,b,c",","))
	fmt.Printf("%q\n",strings.Split("a man a plan canal panama","a "))
	fmt.Printf("%q\n",strings.Split(" xyz ",""))
	fmt.Printf("%q\n",strings.Split("","Bernardo 0'Higgins"))
	fmt.Printf("[%q]",strings.Trim(" !!! Achtung !!! ","! "))
	fmt.Printf("Fields are: %q\n",strings.Fields(" foo bar baz "))

}