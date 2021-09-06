package carts

type Cart struct {
	ID    string  `json:"id"`
	Items ItemMap `json:"items"`
}

type ItemMap map[string]int
