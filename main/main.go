package main


import (
	"net/http"
)
//todo urlparam crop & apply enlarger in frontend for blurry effect


func main(){

	fs:=http.FileServer(http.Dir("pages/"))
	http.Handle("/static/",http.StripPrefix("/static/",fs))
	http.HandleFunc("/",home)
	http.ListenAndServe(":8080",nil)
}
