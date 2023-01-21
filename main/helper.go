package main

import (
  "encoding/json"
  "net/http"
  "net/http/httputil"
  "math/rand"
  "os"
  "log"
)
//todo : to add environmental variable

func parser(query string, pg int) []string{
  cid:=os.Getenv("CID")
  url:="https://api.unsplash.com/search/photos?query="+query+"&per_page="+string(pg)+"&client_id="+cid
  jsondatas,err:=http.Get(url)
  if(err!=nil){
    log.Fatal(err)
  }
  jsondata,errs:=httputil.DumpResponse(jsondatas,true)
  if(errs!=nil){
    log.Fatal(err)
  }
	var data []map[string]interface{}
  json.Unmarshal([]byte(jsondata),&data)
  var imglink []string
  for _,value:=range data{
    results:=value["results"].(map[string]interface{})
    urls:=results["urls"].(map[string]string)
    imglink=append(imglink,urls["raw"])
  }
  return imglink
}
func imageCreate(w http.ResponseWriter, r *http.Request){
	n:=3
	problist:=[3][3]int{}
	lsize:=len(names)
	target:=rand.Intn(lsize)
	for i:=0;i<n;i++{
		for j:=0;j<n;j++{
		problist[i][j]=rand.Intn(lsize)
		print(problist[i][j])
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

