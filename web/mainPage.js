
url = "https://7936-188-130-155-154.eu.ngrok.io";

// let btnAdd = document.getElementById("button__add");
// let btnPlus = document.getElementById("button__plus");
// let btnMinus = document.getElementById("button__minus");
// let countElem = document.getElementById("count");
// let elem = document.getElementById("hidden");

is_validate = false;

let dishes = [];
let num_dishes = 0;

let parsed_init_data = {};

function count_sum(){
    let sum = 0;
    for (let i = 0; i < )
}

function adding_listener() {
    //button__add
    let ind = parseInt(this.id.slice(11));
    let btnAdd = document.getElementById("button__add"+ind);
    let elem = document.getElementById("hidden" + ind);
    let countElem = document.getElementById("count" + ind);
    console.log()
    btnAdd.style.display = "none";
    elem.style.display = "flex";
    countElem.innerHTML = 1;
}

function plusing_listener(){
    //button__plus
    let ind = parseInt(this.id.slice(12));
    let countElem = document.getElementById("count" + ind);
    dishes[ind] += 1;
    countElem.innerHTML = dishes[ind];
}

function minusing_listener(){
    //button__minus
    let ind = parseInt(this.id.slice(13));
    let btnAdd = document.getElementById("button__add"+ind);
    let elem = document.getElementById("hidden" + ind);
    let countElem = document.getElementById("count" + ind);
    dishes[ind] -= 1;
    if (dishes[ind] === 0) {
        elem.style.display = "none";
        btnAdd.style.display = "block";
        dishes[ind] = 1;
    } else countElem.innerHTML = dishes[ind]; 
}

function parse_init_data(initData){
    let data = {}
    let params = initData.split("&");
    for (let i = 0; i < params.length; i++){
        tmp = params[i].split("=");
        data[tmp[0]] = tmp[1]
    }
    let user_data = decodeURI(data.user);
    user_data = user_data.replaceAll("%3A",":");
    user_data = user_data.replaceAll("%2C",",");
    user_data = JSON.parse(user_data);
    data.user = user_data;
    return data
}

function add_dish(data){
    if ('content' in document.createElement('template')) {
        const template = document.querySelector('#dish_template');
        const cont = document.querySelector("#dish_container")
        const clone = template.content.cloneNode(true);
        let img = clone.querySelectorAll("img");
        img[0].src = "https://media.istockphoto.com/photos/hamburger-with-cheese-and-french-fries-picture-id1188412964?k=20&m=1188412964&s=612x612&w=0&h=Ow-uMeygg90_1sxoCz-vh60SQDssmjP06uGXcZ2MzPY=";
        let dish_name_p = clone.getElementById("dish_name");
        let dish_name_price = clone.getElementById("dish_price");

        dish_name_p.textContent = data.name
        dish_name_price.textContent = data.price + " руб.";

        let hiden_div = clone.getElementById("hidden");
        let count_p = clone.getElementById("count");
        let add_btn = clone.getElementById("add");
        let minus_btn = clone.getElementById("button__minus");
        let plus_btn = clone.getElementById("button__plus");

        count_p.id = "count" + num_dishes.toString();
        add_btn.id = "button__add" + num_dishes.toString();
        minus_btn.id = "button__minus" + num_dishes.toString();
        plus_btn.id = "button__plus" + num_dishes.toString();
        hiden_div.id = "hidden" + num_dishes.toString();

        add_btn.addEventListener('click', adding_listener);
        minus_btn.addEventListener('click', minusing_listener);
        plus_btn.addEventListener('click', plusing_listener);

        hiden_div.style.display = "none";
        num_dishes++;
        dishes.push(1)
        cont.appendChild(clone)
    }
    else{
        alert("help")
    }
}

function set_menu(restaurant){
    fetch(url + "/getMenu?restaurant="+restaurant).then(response=>{
        response.json().then(data=>{
            for (let i = 0; i < data.length; i++){
                add_dish(data[i]);
            }
        })
    })
}

function get_restaurants(){
    fetch(url + "/getRestaurant").then(function (response) {
        response.json().then(data=>{
            let s = "";
            for (let i = 0; i < data["Restaurants"].length; i++){
                s += data["Restaurants"][i] + " ";
            }
            alert("kavo\n" + s);
        });
    })
}

function get_user_restaurant(){
    fetch(url + "/get_user_restaurant").then(function (response) {
        response.json().then(data=>{
            let s = "";
            for (let i = 0; i < data["Restaurants"].length; i++){
                s += data["Restaurants"][i] + " ";
            }
            alert("kavo\n" + s);
        });
    }).catch(function(){
        alert("Error during geting user restaurant")
    })
}

Telegram.WebApp.ready();
let initData = Telegram.WebApp.initData || '';
let initDataUnsafe = Telegram.WebApp.initDataUnsafe || {};

fetch(url + "/validate?" + Telegram.WebApp.initData).then(function (response) {
    return response.text();
}).then(function (text) {
    is_validate = true;
    parsed_init_data = parse_init_data(initData);
    // get_user_restaurant()
    set_menu();
}).catch(function () {
    alert("Error on validation occured");
});

// get_restaurants()
