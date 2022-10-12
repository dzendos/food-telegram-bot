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
	state.ServerState.Restaurants = append(state.ServerState.Restaurants, parseDodo())
	state.ServerState.Restaurants = append(state.ServerState.Restaurants, parseDubna())
	state.ServerState.Restaurants = append(state.ServerState.Restaurants, parseMrLim())
	state.ServerState.Restaurants = append(state.ServerState.Restaurants, parseSushi())
	state.ServerState.Restaurants = append(state.ServerState.Restaurants, parseNearFood())
}

func parseDodo() *restaurant.Restaurant {
	restaurant := restaurant.Restaurant{
		Name:            "Dodo pizza",
		Reference:       "https://dodopizza.ru/dubna",
		TelephoneNumber: "+7-800-302-00-60",
		ImageUrl:        "https://sun9-15.userapi.com/impg/Ocu8UXi770N3V3E2b-K2xuUNmpgKiEvp_Ezhhw/8qUjm-hrswI.jpg?size=2560x2560&quality=95&sign=f45821a68464bb2105d35791262485c8&type=album",
	}

	html := scripts.GetDodoHtml()
	myReader := strings.NewReader(html)

	doc, err := goquery.NewDocumentFromReader(myReader)
	if err != nil {
		log.Println(err)
	}

	curMenu := restaurant.Menu
	curMenu = &menu.Menu{}

	// Roll
	doc.Find(".sc-1tpn8pe-3.duQqqx").Each(func(i int, s *goquery.Selection) {
		imageUrl, _ := s.Find(".img").Attr("src") // item picture
		itemName := s.Find(".sc-1tpn8pe-1.jbQjVh").Text()
		itemCostStr := s.Find(".product-control-price").Text()

		if itemName == "" {
			return
		}

		wp := strings.TrimPrefix(strings.TrimSuffix(itemCostStr, " ₽"), "от ")
		itemCost, _ := strconv.ParseFloat(wp, 64)

		curMenu.Positions = append(curMenu.Positions, &position.Position{
			Name:     itemName,
			ImageUrl: imageUrl,
			Price:    itemCost,
			Type:     "Прочее",
		})
	})

	restaurant.Menu = curMenu

	return &restaurant
}

func parseDubna() *restaurant.Restaurant {
	restaurant := restaurant.Restaurant{
		Name:            "Dubna china",
		Reference:       "https://dubna-china.ru/",
		TelephoneNumber: "+7-925-155-02-07",
		ImageUrl:        "https://sun9-north.userapi.com/sun9-88/s/v1/ig2/5hoHe1SjuCfHi6xUTIYi5BKxl-zBcyoso8UypBcaqzlvyAiwHVrwd5OjnZ1HFZKX0MwPPeOLxDGe390k8ysKPg-R.jpg?size=604x604&quality=96&type=album",
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

	return &restaurant
}

func parseSushi() *restaurant.Restaurant {
	restaurant := restaurant.Restaurant{
		Name:            "Полёт",
		Reference:       "https://dubna-sushi.ru/",
		TelephoneNumber: "+7-910-424-00-99",
		ImageUrl:        "https://sun9-4.userapi.com/impf/c849328/v849328504/1703f8/aMlfiWt4qs8.jpg?size=882x882&quality=96&sign=2b4d57c5a161c921977d98d0283cd84a&type=album",
	}

	html := scripts.GetDubnaSushi()
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
			s.Find(".sc-bczRLJ.sc-gsnTZi.Card-sc-mpiw7l-2.cisoKP.jnFvAE.hPFePH").Each(func(i int, s *goquery.Selection) {
				imageUrl, _ := s.Find(".sc-bczRLJ.jBqnLQ.Image-sc-mpiw7l-4.bYtpaX").Attr("src") // item picture
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

	return &restaurant
}

func parseNearFood() *restaurant.Restaurant {
	restaurant := restaurant.Restaurant{
		Name:            "Еда рядом",
		Reference:       "https://app.nearfood.ru/",
		TelephoneNumber: "https://app.nearfood.ru/",
		ImageUrl:        "https://sun9-59.userapi.com/impg/SjTcGsraRihqTpawb5kTWM4vPAIZu3Ga7GVEpg/He5zw2p6a6Q.jpg?size=512x512&quality=96&sign=34a5e8939990c1ee167d253eaa677594&type=album",
	}

	html := scripts.GetNearFoodHtml()
	myReader := strings.NewReader(html)

	doc, err := goquery.NewDocumentFromReader(myReader)
	if err != nil {
		log.Println(err)
	}

	curMenu := restaurant.Menu
	curMenu = &menu.Menu{}

	doc.Find(".sc-fHsOPI.kNcJwe").Each(func(i int, s *goquery.Selection) {
		if s.ChildrenFiltered(".sc-jSMfEi.sc-gUAEMC.jlgugb.jniTUu").Text() != "" {
			// type
			itemType := s.ChildrenFiltered(".sc-jSMfEi.sc-gUAEMC.jlgugb.jniTUu").Text()

			// Roll
			s.Find(".sc-elYLMi.fwTcsh").Each(func(i int, s *goquery.Selection) {
				imageUrl, _ := s.Find("img").Attr("src") // item picture
				itemName := s.Find(".sc-eCYdqJ.jqjtAw").Text()
				itemCostStr := s.Find(".sc-bczRLJ.sc-crXcEl.sc-iqGgem.cTqYlq.iAJwHh.fDzNE").Text()

				itemCost, _ := strconv.ParseFloat(strings.TrimSuffix(itemCostStr, " ₽"), 64)

				curMenu.Positions = append(curMenu.Positions, &position.Position{
					Name:     itemName,
					ImageUrl: imageUrl,
					Price:    itemCost,
					Type:     itemType,
				})
			})
		}
	})

	restaurant.Menu = curMenu

	return &restaurant
}

func parseMrLim() *restaurant.Restaurant {
	//log.Println("Start parsing Mr. Lim")
	restaurant := restaurant.Restaurant{
		Name:            "Mr. Lim",
		Reference:       "https://mrlim.ru/",
		TelephoneNumber: "+7-909-159-15-99",
		ImageUrl:        "https://vsem-edu-oblako.ru/upload/store/merchant1309/1642215585mrlimheaderlogo.png",
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
	// <div class="grid grid-cols-4 gap-8">
	// log.Println("End dodo")

	restaurant.Menu = curMenu

	return &restaurant
}
