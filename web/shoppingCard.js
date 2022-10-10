let url = "https://c259-188-130-155-154.eu.ngrok.io";

Telegram.WebApp.ready();
let initData = Telegram.WebApp.initData || '';
let initDataUnsafe = Telegram.WebApp.initDataUnsafe || {};

let num_dishes = [];
let prices_dishes = [];
let dish_names = [];

let back_btn = document.getElementById("back__button");

function back__button(){
    Telegram.WebApp.openLink(url + "/mainPage.html")
}

back_btn.addEventListener("click", back__button);

function plus_listener(){
    //item__plus
    let ind = parseInt(this.id.slice(10));
    let dish_count = document.getElementById("item__count" + ind);
    let dish_sum = document.getElementById("dish__sum" + ind);
    num_dishes[ind] += 1;
    dish_count.textContent = num_dishes[ind];
    let tmp = prices_dishes[ind] + " X " + num_dishes[ind] + " = " + (prices_dishes[ind]*num_dishes[ind]) + " руб.";
    dish_sum.textContent = tmp
}

function minus_listener(){
    //item__minus
    let ind = parseInt(this.id.slice(11));
    let dish_count = document.getElementById("item__count" + ind);
    let dish_sum = document.getElementById("dish__sum" + ind);
    num_dishes[ind] -= 1;
    if (num_dishes[ind] === 0){
        document.getElementById("element" + ind).remove();
    }
    else{
        dish_count.textContent = num_dishes[ind];
        dish_sum.textContent = prices_dishes[ind] + " X " + num_dishes[ind] + " = " + (prices_dishes[ind]*num_dishes[ind]) + " руб.";
    }
}

function add_dish(data){
    if ('content' in document.createElement('template')) {
        const list = document.querySelector('#list__dishes');
        const template = document.querySelector('#dish__element__template');
        let dish = template.content.cloneNode(true);
        let dish_count = dish.getElementById("item__count");
        let dish_name = dish.getElementById("dish_name");
        dish_name.textContent = data.PositionName;
        dish_count.textContent = data.PositionAmount;
        dish_count.id = "item__count" + num_dishes.length;

        let dish_sum = dish.getElementById("dish__sum");
        console.log("s: " + data.PositionPrice + " : " + data.PositionAmount);
        dish_sum.textContent = data.PositionPrice + " X " + data.PositionAmount + " = " + (data.PositionPrice*data.PositionAmount) + " руб.";
        dish_sum.id = "dish__sum" + num_dishes.length;

        let plus_btn = dish.getElementById("item__plus");
        let minus_btn = dish.getElementById("item__minus");

        plus_btn.id = "item__plus" + num_dishes.length;
        minus_btn.id = "item__minus" + num_dishes.length;

        plus_btn.addEventListener("click",plus_listener);
        minus_btn.addEventListener("click",minus_listener);

        let elem = dish.getElementById("element");
        elem.id = "element" + num_dishes.length;

        prices_dishes.push(data.PositionPrice);
        num_dishes.push(data.PositionAmount);
        dish_names.push(data.PositionName);
        list.appendChild(dish);
    }
}

function get_order(id){
    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = () => {
        if (xhr.readyState === 4) {
            let data = JSON.parse(xhr.response);
            data = data.Order;
            console.log("got answer");
            for (let i = 0; i < data.length; i++){
                console.log(data[i].toString());
                add_dish(data[i]);
            }
        }
    }
    xhr.open("POST", url + "/getOrder", true);
    xhr.setRequestHeader('Content-type', 'application/json');
    let request = new Object();
    request.UserID = id;
    request = JSON.stringify(request);  
    xhr.send(request);
}

function order_to_json(){
    if (is_validate){
        let order = Object();
        order.UserID = initDataUnsafe.user.id;
        let dishes = [];
        for (let i  = 0; i < num_dishes.length; i++){
            let tmp = new Object();
            tmp.PositionName = dish_names[i];
            tmp.PositionAmount = num_dishes[i];
            dishes.push(tmp);
        }
        order.Order = dishes;
        order = JSON.stringify(order);
        return order;
    }
    return '{}';
}

function set_order(){
    var xhr = new XMLHttpRequest();
    let request = order_to_json();
    xhr.open("POST", url + "/sendOrder", false);
    xhr.setRequestHeader('Content-type', 'application/json');
    xhr.send(request);
    Telegram.WebApp.close()
}

let confirm_btn = document.getElementById("confirm__button");
confirm_btn.addEventListener("click", set_order);

fetch(url + "/validate?" + Telegram.WebApp.initData).then(function (response) {
    return response.text();
}).then(function (text) {
    is_validate = true;
    get_order(initDataUnsafe.user.id);
})
