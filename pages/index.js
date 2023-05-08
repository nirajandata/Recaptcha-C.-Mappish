var checkbox=document.getElementById("checkbox")
//todo: to fetch json api from golang server
//don't test anything, it's in active phase of early development
var url ="http://localhost:8080/handle"
var code="oooooo"; //x for checked , o for unchecked
//images=["something.jpg","dont_test_now.jpg","underdevelopment.jpg"]

images=[]
async function imgLoader(){
  const response=await fetch(url)
  const result=await response.json()
  //  let resp=JSON.stringify(response)
  images=result["urls"]
  for(let i=0;i<6;i++){
    let img=document.createElement("img")
    img.src=images[i];
    img.alt=i;
    checkbox.appendChild(img)
  }

  for(let i=0;i<6;i++) {
    checkbox.children[i].onclick=checker[i];
  }
}
imgLoader()

//var children=checkbox.children
//

function checker(ind){
  var item=checkbox.children[ind];
  console.log(code)
  if(code[ind]='o'){
    item.style.border="5 px solid #834";
    code[ind]='x';
  }
  else {
    item.style.border="none";
    code[ind]='o';
  }
}

//add a websocket here for sending code to check
