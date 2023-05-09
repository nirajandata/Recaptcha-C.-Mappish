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


