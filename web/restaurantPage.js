
data = {
    number: 4,
    titles: ["Fuck you", "fuck you and you", "fuck this shit", "fuck fuck fuck"],
    images: ["https://pbs.twimg.com/profile_images/832146364423364609/QolnTZgP_400x400.jpg",
        "https://media-cdn.tripadvisor.com/media/photo-s/1b/99/44/8e/kfc-faxafeni.jpg",
        "https://pbs.twimg.com/profile_images/832146364423364609/QolnTZgP_400x400.jpg",
        "https://media-cdn.tripadvisor.com/media/photo-s/1b/99/44/8e/kfc-faxafeni.jpg"]
}
let url = "https://8f82-188-130-155-154.eu.ngrok.io";
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
        Telegram.WebApp.close()
    }
}

const btnListener = () => {
    //button__restaurant
    let id = parseInt(this.id.slice(18));
    restaurant_names[id];
    Telegram.WebApp.close();
}



const addRestaurant = (name, ImageUrl, id) => {
    if ('content' in document.createElement('template')) {
        let restaurants = document.querySelector("#restaurants");
        let template = document.querySelector('#restaurant');
        let clone = template.content.cloneNode(true);
        let title = clone.getElementById("title");
        let imageUrl = clone.getElementById("image");
        let button = clone.getElementById("button");
        title.id = "title" + id;
        title.textContent = name;
        imageUrl.id = "image" + id;
        imageUrl.src = ImageUrl;
        button.id = "button" + id;
        button.addEventListener("click", btnListener)
        restaurants.appendChild(clone);
        restaurant_names.push(name);
    }
    else{
        alert("Your browser is not supported by this website");
    }
}

function get_restaurants(){
    fetch(url + "/getRestaurant").then(function (response) {
        response.json().then(data=>{
            for (let i = 0; i < data["Restaurants"].length; i++){
                let rest = data["Restaurants"];
                addRestaurant(rest["Name"],rest["Url"], i);
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