package main

import (
  "encoding/json"
  "net/http"
  "io/ioutil"
  "math/rand"
  "os"
  "log"
  "crypto/md5"
  "encoding/hex"
)

type Images struct {
  Urls  []string `json:"urls"`
}

func cerr(err error){
  if(err!=nil){
    log.Fatal(err)
    os.Exit(0)
  }
}

//initially, it was for returning lots of image, but later i realized that it's not that fruitful 
//so, now it only returns 1 image url
func parser(query string) string{
  cid:=string(os.Getenv("CID"))
  url:="https://api.unsplash.com/search/photos?query="+query+"&per_page=1&client_id="+cid
  
  jsondatas,err:=http.Get(url)
  cerr(err)
  jsondata,errs:=ioutil.ReadAll(jsondatas.Body)
  cerr(errs)
  //  log.Printf("json data: %s\n", jsondata)
  var val map[string]interface{}
  err2:=json.Unmarshal([]byte(jsondata),&val)
  cerr(err2)
  result:=val["results"].([]interface{})
  var imglink string
  var test map[string]interface{}

  // TODO :
  // need to handle this issue when the code is X and image fails to get data 

  if(len(result)==0){
   return backup[rand.Intn(len(backup))]
  }
  test=result[0].(map[string]interface{})
  urls:=test["urls"].(map[string]interface{})
  imglink=urls["raw"].(string)
  return imglink
}

func handlers(w http.ResponseWriter, r *http.Request) {

  var api Images
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
    result:=parser(queryname) 
    api.Urls=append(api.Urls,result)
  }

  hasher:=md5.New()
  hasher.Write([]byte(string(codes)))
  hcode:=hex.EncodeToString(hasher.Sum(nil))


  cookie:= http.Cookie{
    Name:"code",
    Value:hcode, 
  }
  http.SetCookie(w,&cookie)
  err := json.NewEncoder(w).Encode(api)
  cerr(err)
}

