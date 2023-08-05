var checkbox=document.getElementById("checkbox")
var recap=document.getElementById("recap")

var url ="http://localhost:8080/handle"
var code=['o','o','o','o','o','o']
images=[]

async function imgLoader(){

  const response=await fetch(url)
  const result=await response.json()

  images=result["urls"]
  titl=document.cookie
  titl=titl.split("=")[0]
  document.getElementById("tname").textContent="Select "+titl + " from the list";


  for(let i=0;i<6;i++){
    let img=document.createElement("img")
    img.src=images[i]+"?&crop&w=100&h=160";
    img.alt=i;
    checkbox.appendChild(img)
  }

  for(let i=0;i<6;i++) {
    checkbox.children[i].addEventListener("click",(e)=>{
      if(e.target.style.border!="none"){
      e.target.style.border="none";
      code[i]='o';
      }
      else{
        e.target.style.border="5px solid red";
        code[i]='x';
        console.log(code)
        //not working here
      }
    });
  }
}
imgLoader()


window.onload=function(){
  recap.action="check?code="+code;
  }
