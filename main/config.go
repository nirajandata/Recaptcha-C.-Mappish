package main

import (
	"os"
	"strings"
	"bufio"
	"log"
	"net/http"
)

var names []string

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
}	

func home(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"static/home.html")
}
