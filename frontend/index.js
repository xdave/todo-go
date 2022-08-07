
function getAddr(route){
    if (route[0] === "/"){
        route = route.slice(1, length(route))
    }
    const addr = "http://localhost:8000/"
    return addr + route
}

function titleColor() {
    document.getElementById("title").style.color = "red";
}

// pings backend to make sure it's up
const ping = async () => {
    const response = await fetch(getAddr('ping'));
    alert(response.statusText);
}
function showLoading(loading){
    let ul = document.getElementById("todo");
    let loader = document.getElementById("loader")
    if (loading){
            // toggle loading text
            ul.style.display="none"
            loader.style.display = "block";
    } else {
            // remove loading text
            ul.style.display="block"
            loader.style.display = "none";
    }

}
// populates UI with todo items
const listTodo = async () => {
    let ul = document.getElementById("todo");
    let loader = document.getElementById("loader")
    showLoading(true)

    // query backend for todo items
    let response = await fetch(getAddr('list'));
    response = await response.text()
    response = response.replace('\n', '')
    response = response.replace('"', '');
    response = response.replace('"', '');
    response = atob(response)
    if (!response || response=="null"){
        loader.textContent= "Add your first Todo!"
        return
    }
    response = JSON.parse(response)
    response.sort((a,b)=>{return a.id > b.id})

    // adds to the UI list of todo apps
    ul.textContent = "" // clears list everytime we fetch, since lists will be small no problem doing this
    for (const [_, v] of Object.entries(response)) {
        const li = createLi(v.id, v.text, v.checked)
        ul.appendChild(li)
    }
    showLoading(false)
}

const toggle = (e) => {
    let checked = e.target.classList.contains('checked')
    // console.log({is_checked: checked})
    let xhr = new XMLHttpRequest()
        let url = getAddr('check_todo')
        xhr.open("PUT", url, true);
        xhr.setRequestHeader('Content-Type', "application/json;charset=UTF-8")
        
        // Create callback to toggle checking in the UI
        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4 && xhr.status === 200) {
                if (checked) {
                    e.target.className = ""
                } else {
                    e.target.className="checked"
                }
                
            } else if (xhr.readyState === 4){
                alert("error")
            }
        }
        let myId = e.target.id
        // used !checked to toggle checking, if it is checked, I want to uncheck. and vice versa 
        let data = {checked: !checked, id: parseInt(myId)}
        console.log(data)
        xhr.send(JSON.stringify(data))
}

function newElement(){
    let inputElem = document.getElementById("myInput")
    const inputValue = inputElem.value;
    if (inputValue === '') {
        alert("You must write something, can't be empty")
    } else {
        inputElem.value = "" // reset the input value
        let xhr = new XMLHttpRequest()
        let url = getAddr('add_todo')
        xhr.open("PUT", url, true);
        xhr.setRequestHeader('Content-Type', 'text/plain')
        
        // Create callback
        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4 && xhr.status === 200) {
                let li = createLi(xhr.responseText, inputValue)
                document.getElementById("todo").appendChild(li)
            }else if (xhr.readyState === 4){
                alert("error")
            }
        }
        xhr.send(inputValue)
      }
}

function createLi(id, text, checked=false){
    let li = document.createElement('li')
    let p = document.createElement("p")
    if (checked) {
        p.className = "checked"
    }
    p.id = id
    p.appendChild(document.createTextNode(text))
    p.addEventListener("click", toggle)

    li.append(p)
    
    let span = document.createElement("span")
    let cross = document.createTextNode("\u00D7")
    span.appendChild(cross)
    span.addEventListener("click", deleteTodo)
    li.append(span)
    return li
}

function deleteTodo(e){
    const parent = e.target.parentElement
    const id = parent.getElementsByTagName("p")[0].id

    let xhr = new XMLHttpRequest()
    const url = getAddr('delete_todo')
    xhr.open("DELETE", url, true);
    xhr.setRequestHeader('Content-Type', 'text/plain')
    
    // Create callback
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            listTodo()
        } else if (xhr.readyState === 4){
            alert("error")
        }
    }
    showLoading(true)
    xhr.send(id)}

function init(){
    listTodo()
    init = function(){} // kill function so it's only used once
}

// setTimeout( listTodo, 1000)
window.onload = init()