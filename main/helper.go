package main

import (
  "encoding/json"
  "net/http"
  "io/ioutil"
  "math/rand"
  "os"
  "log"
)

type Images struct {
  Urls  string `json:"urls"`
}

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

func handlers(w http.ResponseWriter, r *http.Request) {

  var api[] Images
  var queryname string

  n:=len(names)
  target:=rand.Intn(n)

  codes:=[]rune("oooooo")
  for i:=0;i<6;i++ {
    x:=rand.Intn(3) // probability of x being target (we named target as 1) is 1/3;
    if x==1 {
      codes[i]='x';
    } 
  }

  for i:=0;i<6;i++ {
    if codes[i]=='x' {
      queryname=names[target] 
    } else{
      queryname=names[rand.Intn(n)]
    }
    tval:=parser(queryname,"1") 
    t:=Images{
      Urls:tval[0],
    }
    api=append(api,t)
  }
  err := json.NewEncoder(w).Encode(api)
  cerr(err)
}

