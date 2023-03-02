var checkbox=document.getElementById("checkbox")
//todo: to fetch json api from golang server
//don't test anything, it's in active phase of early development
var url ="http://localhost:8080/getimg"

//images=["something.jpg","dont_test_now.jpg","underdevelopment.jpg"]

images=[]
async funtion imgLoader(){
  const response=await fetch(url)
  const result=await response.json()
  images=result.img
}
imgLoader()
for(let i=0;i<6;i++){
  let img=document.createElement("img")
  img.src=images[i]
  img.alt=i;
  checkbox.appendChild(img)
}

//var children=checkbox.children
