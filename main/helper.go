package main

import (
  "encoding/json"
)

func test(query string){
  url="https://api.unsplash.com/search/photos?query="+query+"&per_page=6&client_id=gK52De2Tm_dL5o1IXKa9FROBAJ-LIYqR41xBdlg3X2k"
  jsondata,err=http.Get(url)
  if(err!=nil){
    print("error")
    return
  }
  var data []mp[string] inferface{}
  json.Unmarshal([]byte(jsondata),&data)
  var imglink []string
  for _,value:=range data{
    results:=data["results"].(map[string]interface{})
  }
}
