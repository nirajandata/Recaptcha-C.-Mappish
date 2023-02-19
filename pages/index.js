var checkbox=document.getElementById("checkbox")

//todo: to fetch json api from golang server
//don't test anything, it's in active phase of early development
images=["something.jpg","dont_test_now.jpg","underdevelopment.jpg"]
for(let i=0;i<6;i++){
  let img=document.createElement("img")
  img.src=images[i]
  img.alt=i;
  checkbox.appendChild(img)
}

var children=checkbox.children
//for testing 
console.log(children)
for(int i=0;i<6;i++){
  console.log('dfkjdkfj')
  console.log(children[i].src)
}
