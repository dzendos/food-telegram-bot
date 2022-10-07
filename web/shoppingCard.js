let itemBtnMinus = document.getElementById('item__minus');
let itemBtnPlus = document.getElementById('item__plus');
let itemCount = document.getElementById('item__count');
let intItemCount = 1;
let element1 = document.getElementById('element1');

itemCount.innerHTML = intItemCount;

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