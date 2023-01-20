package main

import (
  "encoding/json"
  "net/http"
  "net/http/httputil"
)

func parser(query string, pg int) []string{
  url:="https://api.unsplash.com/search/photos?query="+query+"&per_page="+string(pg)+"&client_id=gK52De2Tm_dL5o1IXKa9FROBAJ-LIYqR41xBdlg3X2k"
  jsondatas,err:=http.Get(url)
  jsondata,errs:=httputil.DumpResponse(jsondatas,true)
  if(err!=nil || errs!=nil){
    print("error")
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
