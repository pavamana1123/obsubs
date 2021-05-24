class List{
    constructor(list,container,centered){
        this.list=list
        this.centered=!!centered
        this.index=0
        this.centered=false
        this.container=container
    }

    move(inc){
        inc = inc || 1
        if(this.index + inc < 0){
            this.index = this.list.length-Math.abs(inc)
        }else{
            this.index = (this.index + inc) % this.list.length
        }
        this.container.children[this.index].click()
    }

    center(yes){
        if(yes){
            var maxLen = 0
            this.list.map((li)=>{
                if(li.length>maxLen){
                    maxLen = li.length
                }
            })

            this.list = this.list.map((li)=>{
                var pads = maxLen - li.length
                for(var i=0; i<(pads/2)+1; i++){
                    li = " " + li
                    li = li + " "
                }
                return li
            })
        }else{
            this.list = this.list.map((li)=>{
               return li.trim()
            })
        }
    }

    render(){
        this.container.style.display="block"
        this.container.innerHTML=""
        this.list.map((li,i)=>{
          var listItemBox = document.createElement("div")
          listItemBox.innerHTML = `
          <span class="listnum"></span>
          <span class="listItem buttons"></span>
          `
          listItemBox.setAttribute("class","listItemBox")
          listItemBox.children[0].textContent=(i+"").padStart(4,0)
          listItemBox.children[1].textContent=li
          this.container.appendChild(listItemBox)
          listItemBox.addEventListener('click',setOBSText, true)
        })
    }
}


var list
var listElement = document.getElementById("list")
var controlsElement = document.getElementById("controls")
var lastItemClicked = null

if(!list){
    listElement.style.display="none"
    controlsElement.style.display="none"
}

document.getElementById("fileopener").addEventListener('change',(e)=>{
    const reader = new FileReader();
    reader.addEventListener('load', (event) => {
      list = new List(event.target.result.trim().split("\n"), listElement)
      controlsElement.style.display="inline-flex"
      list.render()
      handleKeys()
      listElement.focus()
    });
    reader.readAsText(e.target.files[0]);
})

document.getElementById("centerer").addEventListener('click',()=>{
    list.centered=!list.centered
    list.center(list.centered)
    list.render()
})

document.getElementById("next").addEventListener("click",()=>{
    list.move(1)
})

document.getElementById("prev").addEventListener("click",()=>{
    list.move(-1)
})

function handleKeys(){
    document.addEventListener("keydown",(e)=>{
        switch(e.code){
            case "Enter":
            case "NumpadAdd":
                list.move(1)
                break
            case "NumpadSubtract":
                list.move(-1)
                break;
        }
    })
}

function setOBSText(e){
    var target = e.target

    if(e.target.localName=="span"){
        target = e.target.parentElement
    }

    lastItemClicked && lastItemClicked.children[1].setAttribute("class","listItem buttons")
    target.children[1].setAttribute("class","listItem listSelected buttons")
    target.scrollIntoView({behavior: "smooth", block: "center", inline: "nearest"})
    lastItemClicked = target

    list.index = parseInt(target.children[0].textContent)

    var request = new XMLHttpRequest()
    request.open('POST', "/obs/text", true)
    request.onload = function () {
        if(this.status!=200){
           alert(this.status)
        }
    }
    request.send(JSON.stringify({
        itemName: "LNnames",
        text: target.children[1].textContent
    }))
}