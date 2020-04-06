package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

func getuser(w http.ResponseWriter,r *http.Request)  {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w,"you are get user %s",uid)
}

func modifyuser(w http.ResponseWriter,r *http.Request)  {
	params := r.URL.Query()
	uid := params.Get(":id")
	fmt.Fprintf(w,"you are get user %s",uid)
}

func deleteuser(w http.ResponseWriter,r *http.Request)  {
	params := r.URL.Query()
	uid := params.Get(":id")
	fmt.Fprintf(w,"you are get user %s",uid)
}

func adduser(w http.ResponseWriter,r *http.Request)  {
	params := r.URL.Query()
	uid := params.Get(":id")
	fmt.Fprintf(w,"you are get user %s",uid)
}

func main(){
	mux := routes.New()
	mux.Get("/user/:uid",getuser)
	mux.Post("/user/:id",modifyuser)
	mux.Del("/user/:id",deleteuser)
	mux.Del("/user/:id",adduser)
	http.Handle("/",mux)
	http.ListenAndServe(":8080",nil)
}
