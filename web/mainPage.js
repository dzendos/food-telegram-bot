
url = "https://94d3-188-130-155-166.eu.ngrok.io";

let btnAdd = document.getElementById("button__add");
let btnPlus = document.getElementById("button__plus");
let btnMinus = document.getElementById("button__minus");
let countElem = document.getElementById("count");
let count = 1;
let elem = document.getElementById("hidden");

is_validate = false;

btnAdd.addEventListener('click', () => {
    btnAdd.classList.add("hide");
    elem.classList.add("show");
    countElem.innerHTML = 1;
})

btnPlus.addEventListener('click', () => {
    count += 1;
    countElem.innerHTML = count;
})

btnMinus.addEventListener('click', () => {
    count -= 1;
    if (count === 0) {
        elem.classList.remove("show");
        btnAdd.classList.remove("hide");
        count = 1;
    } else countElem.innerHTML = count; 
})

function parse_init_data(initData){
    let data = {}
    let params = initData.split("&");
    for (let i = 0; i < params.length; i++){
        tmp = params[i].split("=");
        data[tmp[0]] = tmp[1]
    }
    
    let output = "";
    for (let key in data){
        output += (key + "  :  " + data[key] + "\n");
    }
    return output
}

function create_item_html(){

}

function set_menu(menu){

}

Telegram.WebApp.ready();
let initData = Telegram.WebApp.initData || '';
let initDataUnsafe = Telegram.WebApp.initDataUnsafe || {};

fetch(url + "/validate?" + Telegram.WebApp.initData).then(function (response) {
    return response.text();
}).then(function (text) {
    is_validate = true;

}).catch(function () {
    alert("Error on validation occured");
});

