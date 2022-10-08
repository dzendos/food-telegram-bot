package position

type Position struct {
	Name     string  `json:"name"`
	ImageUrl string  `json:"url"`
	Price    float64 `json:"price"`
	Type     string  `json:"type"`
}
