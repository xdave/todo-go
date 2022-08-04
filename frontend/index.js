
function titleColor() {
    document.getElementById("title").style.color = "red";
}

// pings backend to make sure it's up
const ping = async () => {
    const response = await fetch('http://localhost:8000/ping');
    alert(response.statusText);
}

// populates UI with todo items
const listTodo = async () => {
    let ul = document.getElementById("todo");
    let loader = document.getElementById("loader")

    // toggle loading text
    ul.style.display="none"
    loader.style.display = "block";

    // query backend for todo items
    let response = await fetch('http://localhost:8000/list');
    response = await response.text()
    response = response.replace('\n', '')
    response = response.replace('"', '');
    response = response.replace('"', '');
    response = atob(response)
    response = JSON.parse(response)
    console.log(response);

    // adds to the UI list of todo apps
    ul.textContent = "" // clears list everytime we fetch, since lists will be small no problem doing this
    for (const [k, v] of Object.entries(response)) {
        let li = document.createElement("li")
        li.appendChild(document.createTextNode(v.Text))
        if (k.Done) {
            li.className = "checked"
        }
        ul.addEventListener("click", toggle)
        ul.appendChild(li)
    }

    // remove loading text
    ul.style.display="block"
    loader.style.display = "none";
}

const toggle = (e) => {
    if (e.target.classList.contains("checked")) {
        e.target.className = ""
    } else {
        e.target.className="checked"
    }
}

function newElement(){

    const inputValue = document.getElementById("myInput").value;
    if (inputValue === '') {
        alert("You must write something, can't be empty")
    } else {
        let xhr = new XMLHttpRequest()
        let url = 'http://localhost:8000/add_todo'
        xhr.open("PUT", url, true);
        xhr.setRequestHeader('Content-Type', 'text/plain')
        
        // Create callback
        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4 && xhr.status === 200) {
                li = document.createElement('li')
                li.textContent = inputValue
                document.getElementById("todo").appendChild(li)
            }
        }
        xhr.send(inputValue)
      }
}

function init(){
    listTodo()
    init = function(){} // kill function so it's only used once
}

// setTimeout( listTodo, 1000)
window.onload = init()