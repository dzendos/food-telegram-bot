
data = {
    number: 4,
    titles: ["Fuck you", "fuck you and you", "fuck this shit", "fuck fuck fuck"],
    images: ["https://pbs.twimg.com/profile_images/832146364423364609/QolnTZgP_400x400.jpg",
        "https://media-cdn.tripadvisor.com/media/photo-s/1b/99/44/8e/kfc-faxafeni.jpg",
        "https://pbs.twimg.com/profile_images/832146364423364609/QolnTZgP_400x400.jpg",
        "https://media-cdn.tripadvisor.com/media/photo-s/1b/99/44/8e/kfc-faxafeni.jpg"]
}

const btnListener = () => {
    alert("fuck you");
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
    } else {
        alert("fuuuuuuuuck you")
    }
}

for (let i = 0; i < data.number; i++){
    addRestaurant(data.titles[i], data.images[i], i);
}
