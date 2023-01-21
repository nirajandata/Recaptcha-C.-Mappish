package main


import (
	"net/http"
)
type Images struct{
	urls []string
}

func main(){
	initial()
	fs:=http.FileServer(http.Dir("../pages/"))
	http.Handle("/static/",http.StripPrefix("/static/",fs))
	http.HandleFunc("/",home)	
	http.HandleFunc("/json",imageCreate)
	http.ListenAndServe(":8080",nil)
}
