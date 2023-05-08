const checkbox = document.createElement("div");
let imgur="https://plus.unsplash.com/premium_photo-1664035152461-128d4a940b91?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1364&q=80"

for(let i=0;i<6;i++){
  let img=document.createElement("img")
  img.alt=i;
  img.src=imgur
  img.onclick=checker(i);
  checkbox.appendChild(img)
}

//var children=checkbox.children
//


function checker(ind){
  var item=checkbox.children[ind];
  var code="oooooo"; //x for checked , o for unchecked
  console.log(item[0])
  if(code[ind]='o'){
    item.style.border="5 px solid #834";
    code[ind]='x';
  }
  else {
    item.style.border="none";
    code[ind]='o';
  }
}

