package main

import (
	"os"
	"strings"
	"bufio"
	"log"

)
var mp map[string] string

func main(){
	mp=make(map[string]string)
	file,err:=os.Open("class.txt");
	if(err!=nil){
	log.Fatal(err);
	}
	defer file.Close();

	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		lines:=scanner.Text()
		line:=strings.Split(lines,",")
		key,value:=line[0],line[1]
		key=strings.TrimSpace(key)
		value=strings.TrimSpace(value);
		mp[key]=value;
	//	print(mp[key])
	}
}
