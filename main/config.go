package main

import (
	"os"
	"strings"
	"bufio"
	"log"	
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader =websocket.Upgrader{
  CheckOrigin: func(r *http.Request) bool {
    return true // Accepting all requests
  },

  ReadBufferSize:1024,
  WriteBufferSize:1024,
}

var names,backup []string

func initial(){
	file,err:=os.Open("class.txt");
	if(err!=nil){
	log.Fatal(err);
	}
	defer file.Close();

	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		line:=scanner.Text()
		lines:=strings.TrimSpace(line)
		names=append(names,lines)
	}
	backup=[]string{
		"https://images.unsplash.com/photo-1678687536549-5d18a81e66cf",
		"https://images.unsplash.com/photo-1678674259544-62cdafd7388c?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8",
		"https://images.unsplash.com/photo-1678663474154-fe251019117e",
	}
}	

func home(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w,r,"../pages/home.html")
}

func sock(w http.ResponseWriter, r *http.Request){
	conn,err:=upgrader.Upgrade(w,r,nil)
	if err !=nil{
		log.Println(err)
		return
	}

	for{
		messageType, value, err := conn.ReadMessage()
		if err !=nil{
			log.Println(err)
		}
		
		//todo :validate the code with md5 hash

	}
}
