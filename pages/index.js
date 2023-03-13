var checkbox=document.getElementById("checkbox")
//todo: to fetch json api from golang server
//don't test anything, it's in active phase of early development
var url ="http://localhost:8080/getimg"

//images=["something.jpg","dont_test_now.jpg","underdevelopment.jpg"]

images=[]
async function imgLoader(){
  const response=await fetch(url)
  const result=await response.json()
  images=result.urls
}
imgLoader()
for(let i=0;i<6;i++){
  let img=document.createElement("img")
  img.src=images[i]
  img.alt=i;
  img.onclick=checker(i);
  checkbox.appendChild(img)
}

//var children=checkbox.children
//
var code="oooooo"; //x for checked , o for unchecked


function checker(index){
  var item=checkbox.children[index];
  if(code[index]='o'){
    item.style.border="5 px solid #834";
    code[index]='x';
  }
  else {
    item.style.border="none";
    code[index]='o';
  }
}

//add a websocket here for sending code to check
