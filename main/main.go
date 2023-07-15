package main


import (
	"net/http"
)


func main(){

	initial()
	fs:=http.FileServer(http.Dir("../pages/"))

	http.Handle("/static/",http.StripPrefix("/static/",fs))
	http.HandleFunc("/",home)	
	http.HandleFunc("/handle",handlers)
	http.HandleFunc("/echo",sock)
	print("starting server with port 8080 \n")
	http.ListenAndServe(":8080",nil)
}
