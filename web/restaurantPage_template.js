
let url = "{{.url}}";
Telegram.WebApp.ready();
let initData = Telegram.WebApp.initData || '';
let initDataUnsafe = Telegram.WebApp.initDataUnsafe || {};

let is_validate = false;
let restaurant_names = [];

const sendRestaurant = (restaurant) =>{
    if (is_validate){
        let request = new Object();
        request.UserID = initDataUnsafe.user.id;
        request.Restaurant = restaurant;
        request = JSON.stringify(request);
        var xhr = new XMLHttpRequest();
        xhr.open("POST", url + "/sendRestaurant", false);
        xhr.setRequestHeader('Content-type', 'application/json');
        xhr.send(request);
    }
}

function btnListener(){
    //button__restaurant
    console.log(this.id)
    let id = parseInt(this.id.slice(18));
    sendRestaurant(restaurant_names[id])
    // window.open(url + "/mainPage.html","_self")
    Telegram.WebApp.close()
}



const addRestaurant = (name, ImageUrl, id) => {
    if ('content' in document.createElement('template')) {
        let restaurants = document.querySelector("#restaurants");
        let template = document.querySelector('#restaurant');
        let clone = template.content.cloneNode(true);
        let title = clone.getElementById("title");
        let imageUrl = clone.getElementById("image");
        let button = clone.getElementById("button__restaurant");
        title.id = "title" + id;
        title.textContent = name;
        imageUrl.id = "image" + id;
        imageUrl.src = ImageUrl;
        button.id = "button__restaurant" + id;
        button.addEventListener("click", btnListener)
        restaurants.appendChild(clone);
        restaurant_names.push(name);
    }
    else{
        alert("Your browser is not supported by this website");
    }
}

function get_restaurants(){
    fetch(url + "/getRestaurants").then(function (response) {
        response.json().then(data=>{
            let rest = data["Restaurants"];
            for (let i = 0; i < rest.length; i++){
                console.log(rest[i]["Name"] + " : " + rest[i]["Url"]);
                addRestaurant(rest[i]["Name"],rest[i]["Url"], i);
            }
        });
    })
}

fetch(url + "/validate?" + Telegram.WebApp.initData).then(function (response) {
    return response.text();
}).then(function (text) {
    is_validate = true;
}).catch(function () {
    alert("Error on validation occured");
});

get_restaurants()