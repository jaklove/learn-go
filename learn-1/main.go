package learn_1

import (
	"fmt"
	"log"
	"net/http"
)

func sayhelloMame(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_log"])
	for k,v := range r.Form{
		fmt.Println("key:",k)
		fmt.Println("value:",v)
	}
	fmt.Fprintf(w,"hello word")

}

func main()  {
	http.HandleFunc("/",sayhelloMame)
	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("监听错误",err.Error())
	}


}



