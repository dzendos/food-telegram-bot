package scrapper

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dzendos/dubna/internal/model/restaurant"
	"github.com/dzendos/dubna/internal/model/state"
	"github.com/dzendos/dubna/scripts"
)

func InitializeServerState() {
	state.ServerState.Restaurants = append(state.ServerState.Restaurants, parseDodo())

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
