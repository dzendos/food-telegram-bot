
url = "{{.url}}";



let confirm_btn_text = document.getElementById("button__confirm__text");
let confirm_btn = document.getElementById("button__confirm");
let restaurant_title = document.getElementById("restaurant__title");
is_validate = false;

let restaurant_name = "";
let dishes = [];
let prices = [];
let dish_names = [];
let num_dishes = 0;

let parsed_init_data = {};

function count_sum() {
    let sum = 0;
    for (let i = 0; i < num_dishes; i++){
        sum += dishes[i] * prices[i];
    }
    return sum;
}

function adding_listener() {
    //button__add
    let ind = parseInt(this.id.slice(11));
    let btnAdd = document.getElementById("button__add" + ind);
    let elem = document.getElementById("hidden" + ind);
    let countElem = document.getElementById("count" + ind);
    console.log();
    btnAdd.style.display = "none";
    elem.style.display = "flex";
    countElem.innerHTML = 1;
    dishes[ind] = 1;
    confirm_btn_text.textContent = "Подтвердить " + count_sum() + "руб";
}

function plusing_listener() {
    //button__plus
    let ind = parseInt(this.id.slice(12));
    let countElem = document.getElementById("count" + ind);
    dishes[ind] += 1;
    countElem.innerHTML = dishes[ind];
    confirm_btn_text.textContent = "Подтвердить " + count_sum() + "руб";
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
        dishes[ind] = 0;
    } else countElem.innerHTML = dishes[ind]; 
    let sum = count_sum();
    if (sum === 0) {
        confirm_btn_text.textContent = "Отменить";
    } else confirm_btn_text.textContent = "Подтвердить " + sum + "руб" 
}

function confirm_btn_listener(){
    let request = Object();
    let order = [];
    for (let i = 0; i < num_dishes; i++) {
        if (dishes[i] != 0){
            let cur_dish = Object();
            cur_dish.PositionName = dish_names[i];
            cur_dish.PositionAmount = dishes[i];
            order.push(cur_dish);
        }
    }
    console.log(parsed_init_data)
    request.UserID = parsed_init_data.user.id
    request.Order = order;
    request = JSON.stringify(request)
    var xhr = new XMLHttpRequest();
    xhr.open("POST", url + "/sendOrder", false);
    xhr.setRequestHeader('Content-type', 'application/json');
    xhr.send(request);
    Telegram.WebApp.close()

}

function parse_init_data(initData) {
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
        img[0].src = data.url;
        let dish_name_p = clone.getElementById("dish_name");
        let dish_name_price = clone.getElementById("dish_price");

        dish_name_p.textContent = data.name
        dish_name_price.textContent = data.price + " руб.";

        let hidden_div = clone.getElementById("hidden");
        let count_p = clone.getElementById("count");
        let add_btn = clone.getElementById("add");
        let minus_btn = clone.getElementById("button__minus");
        let plus_btn = clone.getElementById("button__plus");

        count_p.id = "count" + num_dishes.toString();
        add_btn.id = "button__add" + num_dishes.toString();
        minus_btn.id = "button__minus" + num_dishes.toString();
        plus_btn.id = "button__plus" + num_dishes.toString();
        hidden_div.id = "hidden" + num_dishes.toString();

        add_btn.addEventListener('click', adding_listener);
        minus_btn.addEventListener('click', minusing_listener);
        plus_btn.addEventListener('click', plusing_listener);

        hidden_div.style.display = "none";
        num_dishes++;
        dishes.push(0);
        dish_names.push(data.name);
        prices.push(data.price);
        cont.appendChild(clone)
    } else {
        alert("help")
    }
}

function set_menu(id){
    let request = new Object();
    request.UserID = id;
    request = JSON.stringify(request);
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = () => {
        if (xhr.readyState === 4) {
            let data = JSON.parse(xhr.response);
            console.log(restaurant_title);
            restaurant_title.textContent = data.Restaurant;
            for (let i = 0; i < data.Menu.length; i++){
                add_dish(data.Menu[i])
            }
        }
    }
    xhr.open("POST", url + "/getMenu", true);
    xhr.setRequestHeader('Content-type', 'application/json');
    xhr.send(request);

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
console.log("started")
confirm_btn.addEventListener('click', confirm_btn_listener)
Telegram.WebApp.ready();
let initData = Telegram.WebApp.initData || '';
let initDataUnsafe = Telegram.WebApp.initDataUnsafe || {};

fetch(url + "/validate?" + Telegram.WebApp.initData).then(function (response) {
    return response.text();
}).then(function (text) {
    is_validate = true;
    parsed_init_data = parse_init_data(initData);
    // get_user_restaurant()
    set_menu(initDataUnsafe.user.id);
}).catch(function () {
    alert("Error on validation occured");
});


// get_restaurants()
