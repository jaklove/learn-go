package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {

}

func (m *MyMux) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	fmt.Println(r.URL.Path)
	if r.URL.Path == "/user"{
		sayHelloGame(w,r)
		return
	}
	http.NotFound(w,r)
	return
}

func sayHelloGame(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"hello word!")
}

func main()  {
	mux := &MyMux{}
	http.ListenAndServe(":8080",mux)


}
