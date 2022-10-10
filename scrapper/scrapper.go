package scrapper

import (
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dzendos/dubna/internal/model/menu"
	"github.com/dzendos/dubna/internal/model/position"
	"github.com/dzendos/dubna/internal/model/restaurant"
	"github.com/dzendos/dubna/internal/model/state"
	"github.com/dzendos/dubna/scripts"
)

func InitializeServerState() {
	state.ServerState.Restaurants = append(state.ServerState.Restaurants, parseDubna())
	state.ServerState.Restaurants = append(state.ServerState.Restaurants, parseMrLim())
}

func parseDubna() *restaurant.Restaurant {
	restaurant := restaurant.Restaurant{
		Name:      "Dubna china",
		Reference: "https://dubna-china.ru/",
		ImageUrl:  "https://sun7-15.userapi.com/impg/Ocu8UXi770N3V3E2b-K2xuUNmpgKiEvp_Ezhhw/8qUjm-hrswI.jpg?size=2560x2560&quality=95&sign=f45821a68464bb2105d35791262485c8&type=album",
	}

	html := scripts.GetDubnaChinaHtml()
	myReader := strings.NewReader(html)

	doc, err := goquery.NewDocumentFromReader(myReader)
	if err != nil {
		log.Println(err)
	}

	curMenu := restaurant.Menu
	curMenu = &menu.Menu{}

	doc.Find(".sc-bczRLJ.hJYFVB").Each(func(i int, s *goquery.Selection) {
		if s.ChildrenFiltered(".sc-bczRLJ.dqChRk").Text() != "" {
			// type
			itemType := s.ChildrenFiltered(".sc-bczRLJ.dqChRk").Text()

			// Roll
			s.Find(".sc-bczRLJ.sc-gsnTZi.Card-sc-2npckq-2.cisoKP.jnFvAE.jbjTQa").Each(func(i int, s *goquery.Selection) {
				imageUrl, _ := s.Find(".sc-bczRLJ.jBqnLQ.Image-sc-2npckq-4.kPlshG").Attr("src") // item picture
				itemName := s.Find(".sc-bczRLJ.gxeRtG").Text()
				itemCostStr := s.Find(".sc-bczRLJ.hWxACP").Text()

				itemCost, _ := strconv.ParseFloat(strings.TrimSuffix(itemCostStr, " ₽"), 64)

				curMenu.Positions = append(curMenu.Positions, &position.Position{
					Name:     itemName,
					ImageUrl: imageUrl,
					Price:    itemCost,
					Type:     itemType,
				})

				//log.Println(imageUrl, itemName, itemCost)
			})
		}
	})

	restaurant.Menu = curMenu

	log.Println(curMenu)

	return &restaurant
}

func parseMrLim() *restaurant.Restaurant {
	log.Println("Start parsing Mr. Lim")
	restaurant := restaurant.Restaurant{
		Name:      "Mr. Lim",
		Reference: "https://mrlim.ru/",
		ImageUrl:  "https://vsem-edu-oblako.ru/upload/store/merchant1309/1642215585mrlimheaderlogo.png",
	}

	html := scripts.GetMrLimHtml()
	myReader := strings.NewReader(html)

	doc, err := goquery.NewDocumentFromReader(myReader)
	if err != nil {
		log.Println(err)
	}
	curMenu := restaurant.Menu
	curMenu = &menu.Menu{}
	doc.Find(`div[data-fetch-key="1"]`).Each(func(i int, s *goquery.Selection) {
		s.Find(`section[class="py-8 mb-8"]`).Each(func(i int, s *goquery.Selection) {
			s.Find(`.grid`).Each(func(i int, s *goquery.Selection) {
				s.Find(`.card`).Each(func(i int, s *goquery.Selection) {
					src, is_here := s.Find(`img`).Attr("src")
					if is_here {
						if src == "/_nuxt/5e5e01b09a7e549d74e0acec108c84c6.svg" {
							src = "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.flaticon.com%2Ffree-icon%2Fburger_198416&psig=AOvVaw2E6MWWwYc15-oqZGTbB4DN&ust=1665433673828000&source=images&cd=vfe&ved=0CAwQjRxqFwoTCICXl4b-0_oCFQAAAAAdAAAAABAO"
						}
					} else {
						src = "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.flaticon.com%2Ffree-icon%2Fburger_198416&psig=AOvVaw2E6MWWwYc15-oqZGTbB4DN&ust=1665433673828000&source=images&cd=vfe&ved=0CAwQjRxqFwoTCICXl4b-0_oCFQAAAAAdAAAAABAO"
					}

					name := s.Find(`div.mb-4`).Find(`span.font-semibold`).Text()
					name = strings.TrimSpace(name)
					priceRaw := s.Find(`div.mt-auto`).Find(`span.text-lg`).Text()
					price, _ := strconv.ParseFloat(strings.Trim(strings.TrimSpace(priceRaw), " ₽"), 64)
					curMenu.Positions = append(curMenu.Positions, &position.Position{
						Name:     name,
						ImageUrl: src,
						Price:    price,
						Type:     "Прочее",
					})
					// log.Println("name: ", name, " price: ", price)
				})
			})
		})
	})
	log.Println("End dodo")

	restaurant.Menu = curMenu

	return &restaurant
}
