package main

import (
	"os"
	"strings"
	"bufio"
	"log"
	"math/rand"
	"encoding/json"
)
names:=[]int{}
func init(){
	file,err:=os.Open("class.txt");
	if(err!=nil){
	log.Fatal(err);
	}
	defer file.Close();

	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		lines:=scanner.Text()
		lines=strings.TrimSpace(lines)
		names=append(names,lines)
	}
	//this is the 2d list for 3x3 target-img probability filling 

}	
func imageSend(w http.ResponseWriter, r *http.Request){
	n:=3
	problist:=[n][n]int{}
	lsize:=len(names)
	target:=rand.Intn(lsize)
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
		problist[i][j]=rand.Intn(lsize)
		}
	}
	//  3 target-img of each row in random column
	for i:=0;i<n;i++){
		x:=rand.Intn(n)
		problist[i][x]=target
	}
	var api Images
	link:="https://something.com/"
	for( i:=0;i<n;i++){
		for(j:=0;j<n;j++){
		text:=link+names[problist[i][j]]
		api.urls=append(api.urls,text)
		}
	}
	json.NewEncoder(w).Encode(api)
}
