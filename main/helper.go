package main

import (
  "encoding/json"
  "net/http"
  "io/ioutil"
  "math/rand"
  "os"
  "log"
)

func cerr(err error){
  if(err!=nil){
    log.Fatal(err)
    os.Exit(0)
  }
}
func parser(query,pgg string) []string{
  cid:=string(os.Getenv("CID"))
  url:="https://api.unsplash.com/search/photos?query="+query+"&per_page="+pgg+"&client_id="+cid
  jsondatas,err:=http.Get(url)
  cerr(err)
  jsondata,errs:=ioutil.ReadAll(jsondatas.Body)
  cerr(errs)
//  log.Printf("json data: %s\n", jsondata)
  var val map[string]interface{}
  err2:=json.Unmarshal([]byte(jsondata),&val)
  cerr(err2)
  var imglink []string
  result:=val["results"].([]interface{})
  n:=len(result)
  var test map[string]interface{}
  for i:=0;i<n ;i++ {
    test=result[i].(map[string]interface{})
    urls:=test["urls"].(map[string]interface{})
    imglink=append(imglink,urls["raw"].(string))
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
		}
	}
	//  3 target-img of each row in random column
	for i:=0;i<n;i++ {
		x:=rand.Intn(n)
		problist[i][x]=target
	}
	var api []Images
	targetImg:=parser(names[target],"3")
	tcount:=0
	for i:=0;i<n;i++ {
		for j:=0;j<n;j++{
		var test Images 
		text:=names[problist[i][j]]
		print(text)
		if problist[i][j]==target {
		  log.Println("count is",tcount)
		  test=Images{
        urls: targetImg[tcount%3],
		  }
			api=append(api,test)
			tcount+=1
		} else{
     test=Images{
       urls:targetImg[2],
     } 
			randomImg:=parser(text,"1")
			test=Images{
			  urls:randomImg[0],
			}
			log.Println(randomImg[0])
			api=append(api,test)
		}
		log.Println("this is",api)
		}
	}
//	log.Println(api.urls)
	json.NewEncoder(w).Encode(api)
}

