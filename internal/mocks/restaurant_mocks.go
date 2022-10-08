package mocks

import (
	"github.com/dzendos/dubna/internal/model/menu"
	"github.com/dzendos/dubna/internal/model/position"
	"github.com/dzendos/dubna/internal/model/restaurant"
	"github.com/dzendos/dubna/internal/model/state"
)

func FillMockRestaurant() {
	position1 := &position.Position{
		Name:     "Пицца из половинок",
		ImageUrl: "https://cdn.dodostatic.net/static/Img/Products/9a468e7d8f5149d89464b4e174599b65_366x366.png",
		Price:    690,
		Type:     "Pizza",
	}

	position2 := &position.Position{
		Name:     "Миу-пицца с ветчиной и сюрприз",
		ImageUrl: "https://cdn.dodostatic.net/static/Img/Products/ae7892e0a7b44a9ab6bfc0a9d3c8eb0d_366x366.png",
		Price:    569,
		Type:     "Pizza",
	}

	position3 := &position.Position{
		Name:     "Миу-пицца с пепперони и сюрприз",
		ImageUrl: "https://cdn.dodostatic.net/static/Img/Products/24d44448b8b4475bba6693f45eb959d0_366x366.png",
		Price:    579,
		Type:     "Pizza",
	}

	state.ServerState.Restaurants = append(state.ServerState.Restaurants, &restaurant.Restaurant{
		Name:          "Dodo",
		Reference:     "https://dodopizza.ru/abakan/kirova101",
		ImageUrl:      "https://sun7-15.userapi.com/impg/Ocu8UXi770N3V3E2b-K2xuUNmpgKiEvp_Ezhhw/8qUjm-hrswI.jpg?size=2560x2560&quality=95&sign=f45821a68464bb2105d35791262485c8&type=album",
		PaymentMethod: restaurant.ViaMessenger,
		Menu: &menu.Menu{
			Positions: []*position.Position{position1, position2, position3},
		},
	})
}
