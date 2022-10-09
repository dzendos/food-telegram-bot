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

	return &restaurant
}
