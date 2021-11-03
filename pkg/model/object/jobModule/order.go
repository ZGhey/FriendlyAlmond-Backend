package jobModule

import "FriendlyAlmond_backend/pkg/model"

type ApiOrder struct {
	model.JSONResult
	Data []*RespOrder `json:"data"`
}

type RespOrder struct {
	BoatName      string  `json:"boat_name"`
	Color         string  `json:"color"`
	BoatmodelName string  `json:"model"`
	TotalPrice    float32 `json:"total_price"`
	UserName      string  `json:"user_name"`
	OrderNum      string  `json:"order_num"`
	OrderDate     string  `json:"order_date"`
}
