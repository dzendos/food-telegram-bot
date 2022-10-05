

let btnAdd = document.getElementById("button__add");
let btnPlus = document.getElementById("button__plus");
let btnMinus = document.getElementById("button__minus");
let countElem = document.getElementById("count");
let count = 1;
let elem = document.getElementById("hidden");

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
let aboba = document.getElementById("aboba");
// aboba.innerHTML = "1";
Telegram.WebApp.ready();
// aboba.innerHTML = "2";
let initData = Telegram.WebApp.initData || '';
// aboba.innerHTML = "3";
let initDataUnsafe = Telegram.WebApp.initDataUnsafe || {};
// aboba.innerHTML = "4";
console.log("{{ .WebAppURL }}/validate?");
// alert("your name is: " + Telegram.WebApp.initDataUnsafe.user.first_name)
fetch("https://f865-188-130-155-154.eu.ngrok.io/validate?" + Telegram.WebApp.initData).then(function (response) {
    return response.text();
}).then(function (text) {
    aboba.innerHTML = text;
    // alert("result: " + text);
}).catch(function () {
    aboba.innerHTML = "cringe";
});
