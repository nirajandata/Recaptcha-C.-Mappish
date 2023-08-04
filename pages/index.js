
let socket = new WebSocket("ws://localhost:8080/echo");

var checkbox=document.getElementById("checkbox")

var url ="http://localhost:8080/handle"
export var code="oooooo"; //x for checked , o for unchecked

images=[]
async function imgLoader(){
  const response=await fetch(url)
  const result=await response.json()
  //  let resp=JSON.stringify(response)
  images=result["urls"]
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
      }
      else{
        e.target.style.border="5px solid red";
      }
    });
  }
}
imgLoader()


function sender(){
  socket.send(code);
}




