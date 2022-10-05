

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

