package main

import (
	"os"
	"strings"
	"bufio"
	"log"
	"math/rand"
	"encoding/json"
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
func imageCreate(w http.ResponseWriter, r *http.Request){
	n:=3
	problist:=[3][3]int{}
	lsize:=len(names)
	target:=rand.Intn(lsize)
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
		problist[i][j]=rand.Intn(lsize)
		}
	}
	//  3 target-img of each row in random column
	for i:=0;i<n;i++ {
		x:=rand.Intn(n)
		problist[i][x]=target
	}
	var api Images
	targetImg:=parser(names[target],3)
	tcount:=0
	for i:=0;i<n;i++ {
		for j:=0;j<n;j++{
		text:=names[problist[i][j]]
		if(problist[i][j]==target){
			api.urls=append(api.urls,targetImg[tcount])
			tcount+=1
		} else{
			randomImg:=parser(text,1)
			api.urls=append(api.urls,randomImg[0])
		}
		}
	}
	json.NewEncoder(w).Encode(api)
}
