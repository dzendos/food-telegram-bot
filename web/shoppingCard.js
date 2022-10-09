let url = "https://8f82-188-130-155-154.eu.ngrok.io";

// let itemBtnMinus = document.getElementById('item__minus');
// let itemBtnPlus = document.getElementById('item__plus');
// let itemCount = document.getElementById('item__count');
// let intItemCount = 1;
// let element1 = document.getElementById('element1');
// itemCount.innerHTML = intItemCount;


Telegram.WebApp.ready();
let initData = Telegram.WebApp.initData || '';
let initDataUnsafe = Telegram.WebApp.initDataUnsafe || {};

let num_dishes = [];
let prices_dishes = [];

itemBtnPlus.addEventListener('click', () => {
    intItemCount += 1;
    itemCount.innerHTML = intItemCount;
})

itemBtnMinus.addEventListener('click', () => {
    intItemCount -= 1;
    if (intItemCount > 0) {
        itemCount.innerHTML = intItemCount;
    } else {
        element1.style.display = "none";
        intItemCount = 0;
    }
})

function plus_listener(){
    //item__plus
    let ind = parseInt(this.id.slice(10));
    let dish_count = document.getElementById("item__plus" + ind);
    let dish_sum = document.getElementById("item__count" + ind);
    num_dishes[ind] += 1;
    dish_count.textContent = num_dishes[ind];
    dish_sum.textContent = prices_dishes[ind] + " X " + dish_names[ind] + " = " + (prices_dishes[ind]*dish_names[ind]) + " руб.";
}

function minus_listener(){
    //item__minus
    let ind = parseInt(this.id.slice(11));
    let dish_count = document.getElementById("item__plus" + ind);
    let dish_sum = document.getElementById("item__count" + ind);
    num_dishes[ind] -= 1;
    if (num_dishes[ind] === 0){
        document.getElementById("element" + ind).remove();
    }
    else{
        dish_count.textContent = num_dishes[ind];
        dish_sum.textContent = prices_dishes[ind] + " X " + dish_names[ind] + " = " + (prices_dishes[ind]*dish_names[ind]) + " руб.";
    }
}

function add_dish(dish){
    if ('content' in document.createElement('template')) {
        const list = document.querySelector('#list__dishes');
        const template = document.querySelector('#dish__element__template');
        let dish = template.cloneNode(true);

        let dish_count = dish.getElementById("item__count");
        dish_count.textContent = dish.PositionName;
        dish_count.id = "item__count" + num_dishes.length;

        let dish_sum = dish.getElementById("dish__sum")
        dish_sum.textContent = dish.Price + " X " + dish.PositionAmount + " = " + (dish.Price*dish.PositionAmount) + " руб.";
        dish_sum.id = "dish__sum" + num_dishes.length;

        let plus_btn = dish.getElementById("item__plus");
        let minus_btn = dish.getElementById("item__minus");

        plus_btn.id = "item__plus" + num_dishes.length;
        minus_btn.id = "item__minus" + num_dishes.length;

        plus_btn.addEventListener(plus_listener);
        minus_btn.addEventListener(minus_listener);

        let elem = dish.getElementById("element");
        elem.id = "element" + num_dishes.length;

        prices_dishes.push(dish.Price);
        num_dishes.push(dish.PositionAmount);
    }
}

function get_order(id){
    fetch(url + "/getOrder?UserId=" + id).then(response => {
        response.json().then(data => {
            data = data.order;
            for (let i = 0; i < data.length; i++){
                add_dish(data[i]);
            }
        })
    })
}

fetch(url + "/validate?" + Telegram.WebApp.initData).then(function (response) {
    return response.text();
}).then(function (text) {
    alert(text);
    is_validate = true;
    get_order(initDataUnsafe.user.id);
}).catch(function () {
    alert("Error on validation occured");
});

