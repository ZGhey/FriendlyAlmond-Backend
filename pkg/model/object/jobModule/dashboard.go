package jobModule

import "FriendlyAlmond_backend/pkg/model"

type RespMostPopular struct {
	model.JSONResult
	Data MostPopular `json:"data"`
}

type MostPopular struct {
	Colors     []Color     `json:"colors"`
	Sections   []Section   `json:"sections"`
	Components []Component `json:"components"`
}
type Color struct {
	Color string `json:"color"`
	Total int32  `json:"total"`
}

type Section struct {
	Section string `json:"section"`
	Total   int32  `json:"total"`
}

type Component struct {
	Component string `json:"component"`
	Total     int32  `json:"total"`
}

type RespTotalSales struct {
	model.JSONResult
	Data TotalSales `json:"data"`
}

type TotalSales struct {
	LastOneMonth   float32 `json:"last_one_month"`
	LastThreeMonth float32 `json:"last_three_month"`
	LastSixMonth   float32 `json:"last_six_month"`
	LastOneYear    float32 `json:"last_one_year"`
}
