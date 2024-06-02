package entity

type Animal struct {
	ID       string  `json:"id"`
	Name     string  `json:"animal" example:"sheep"`
	Weight   float64 `json:"weight" example:"56.9"`
	IsHungry bool    `json:"is_hungry" example:"true"`
}
