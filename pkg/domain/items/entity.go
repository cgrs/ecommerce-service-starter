package items

import "net/url"

type Item struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       *url.URL `json:"image"`
	UnitPrice   float64  `json:"price"`
}
