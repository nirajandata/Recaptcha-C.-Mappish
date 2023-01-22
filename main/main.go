package main


import (
	"net/http"
)
type Images struct{
	urls string `json:urls`
}

func main(){
	initial()
	fs:=http.FileServer(http.Dir("../pages/"))
	http.Handle("/static/",http.StripPrefix("/static/",fs))
	http.HandleFunc("/",home)	
	http.HandleFunc("/jsonthrow",imageCreate)
	print("starting server with port 8080 \n")
	http.ListenAndServe(":8080",nil)
}
