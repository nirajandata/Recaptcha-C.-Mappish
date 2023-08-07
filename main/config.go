package main

import (
	"os"
	"strings"
	"bufio"
	"log"	
	"net/http"
)

var names []string
var backup map[string] string
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
	backup=make(map[string]string)
	
	// when the api fails 
 
  backup["human"]	="https://images.unsplash.com/photo-1503023345310-bd7c1de61c7d?ixid=M3wxNDM4ODV8MHwxfHNlYXJjaHwxfHxodW1hbnxlbnwwfHx8fDE2OTE0MDQzNTB8MA&ixlib=rb-4.0.3"
  backup["art"]	=	"https://images.unsplash.com/photo-1579783902614-a3fb3927b6a5?ixid=M3wxNDM4ODV8MHwxfHNlYXJjaHwxfHxhcnR8ZW58MHx8fHwxNjkxNDA0MjA5fDA&ixlib=rb-4.0.3"
	backup["rabbit"]=	"https://images.unsplash.com/photo-1518796745738-41048802f99a?ixid=M3wxNDM4ODV8MHwxfHNlYXJjaHwxfHxyYWJiaXR8ZW58MHx8fHwxNjkxNDA0NzEyfDA&ixlib=rb-4.0.3"
	backup["horse"]="https://images.unsplash.com/photo-1598974357801-cbca100e65d3?ixid=M3wxNDM4ODV8MHwxfHNlYXJjaHwxfHxob3JzZXxlbnwwfHx8fDE2OTE0MDU1NTN8MA&ixlib=rb-4.0.3"
	backup["house"]="https://images.unsplash.com/photo-1518780664697-55e3ad937233?ixid=M3wxNDM4ODV8MHwxfHNlYXJjaHwxfHxob3VzZXxlbnwwfHx8fDE2OTE0MDU1MjZ8MA&ixlib=rb-4.0.3"
	backup["mountain"]="https://images.unsplash.com/photo-1570654621852-9dd25b76b38d?ixid=M3wxNDM4ODV8MXwxfHNlYXJjaHwxfHxtb3VudGFpbnxlbnwwfHx8fDE2OTE0MDU1MjV8MA&ixlib=rb-4.0.3"
	backup["cat"]="https://images.unsplash.com/photo-1514888286974-6c03e2ca1dba?ixid=M3wxNDM4ODV8MHwxfHNlYXJjaHwxfHxjYXR8ZW58MHx8fHwxNjkxNDA2NDcyfDA&ixlib=rb-4.0.3"
			
}

func home(w http.ResponseWriter, r *http.Request){
  http.ServeFile(w,r,"../pages/home.html")
}

