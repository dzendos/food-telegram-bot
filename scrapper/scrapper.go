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
	state.ServerState.Restaurants = append(state.ServerState.Restaurants, parseMrLim())
}

func parseDodo() *restaurant.Restaurant {
	log.Println("Start Dubna")
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

	doc.Find(".sc-bczRLJ.hJYFVB").Each(func(i int, s *goquery.Selection) {
		// type
		s.Find(".sc-bczRLJ.dqChRk").Each(func(i int, s *goquery.Selection) {
			log.Println(s.Text())
		})

		// Roll
		s.Find(".sc-bczRLJ.sc-gsnTZi.Card-sc-2npckq-2.cisoKP.jnFvAE.jbjTQa").Each(func(i int, s *goquery.Selection) {
			// roll name
			s.Find(".sc-bczRLJ.gxeRtG").Each(func(i int, s *goquery.Selection) {
				log.Println(s.Text())
			})
		})

		log.Println(i+1, s.Text())
		s.Find("")
	})

	log.Println("End dodo")

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
		// type
		// s.Find("")
		// log.Println("aboba")
		s.Find(`section[class="py-8 mb-8"]`).Each(func(i int, s *goquery.Selection) {
			s.Find(`.grid`).Each(func(i int, s *goquery.Selection) {
				log.Println("section")
				s.Find(`.card`).Each(func(i int, s *goquery.Selection) {
					log.Println("card")
					src, is_here := s.Find(`img`).Attr("src")
					if is_here {
						if src == "/_nuxt/5e5e01b09a7e549d74e0acec108c84c6.svg" {
							src = "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.flaticon.com%2Ffree-icon%2Fburger_198416&psig=AOvVaw2E6MWWwYc15-oqZGTbB4DN&ust=1665433673828000&source=images&cd=vfe&ved=0CAwQjRxqFwoTCICXl4b-0_oCFQAAAAAdAAAAABAO"
						}
					} else {
						src = "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.flaticon.com%2Ffree-icon%2Fburger_198416&psig=AOvVaw2E6MWWwYc15-oqZGTbB4DN&ust=1665433673828000&source=images&cd=vfe&ved=0CAwQjRxqFwoTCICXl4b-0_oCFQAAAAAdAAAAABAO"
					}
					name := s.Find(`.font-semibold.text-gray-700.dark:text-white.text-lgg"`).Text()
					name = strings.TrimSpace(name)
					price, _ := strconv.ParseFloat(strings.Trim(s.Find(`.text-lg.font-semibold.text-gray-700.dark:text-white`).Text(), " ₽"), 64)
					curMenu.Positions = append(curMenu.Positions, &position.Position{
						Name:     "aboba",
						ImageUrl: src,
						Price:    12,
						Type:     "Прочее",
					})
					log.Println(name, price, src)
				})
			})
		})
	})
	// <div class="grid grid-cols-4 gap-8">
	log.Println("End dodo")

	return &restaurant
}
