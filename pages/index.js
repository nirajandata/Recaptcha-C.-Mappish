var checkbox=document.getElementById("checkbox")

images=[]
for(let i=0;i<6;i++){
  let img=document.createElement("img")
  img.src=images[i%2]
  checkbox.appendChild(img)
}
