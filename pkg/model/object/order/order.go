package order

import (
	"FriendlyAlmond_backend/pkg/model"
	"time"
)

type ReqOrder struct {
	OrderId       int32   `json:"order_id"`
	BoatName      string  `json:"boat_name"`
	Color         string  `json:"color"`
	BoatmodelName string  `json:"boatmodel_name"`
	SectionId     []int32 `json:"section_id"`
	ComponentId   []int32 `json:"component_id"`
	TotalPrice    float32 `json:"total_price"`
	CategoryName  string  `json:"category_name"`
	Uid           string  `json:"uid"`
	Comment       string  `json:"comment"`
}

type ApiOrder struct {
	model.JSONResult
	Data []*RespOrder `json:"data"`
}

type RespOrder struct {
	BoatName      string   `json:"boat_name"`
	Color         string   `json:"color"`
	BoatmodelName string   `json:"boatmodel_name"`
	SectionId     []int64  `json:"section_id"`
	ComponentId   []int64  `json:"component_id"`
	Options       []string `json:"options"`
	TotalPrice    float32  `json:"total_price"`
	CategoryName  string   `json:"category_name"`
	Uid           string   `json:"uid"`
	JobId         int32    `json:"job_id"`
	OrderId       int32    `json:"order_id"`
	UserName      string   `json:"user_name"`
	OrderDate     string   `json:"order_date"`
	Comment       string   `json:"comment"`
}

type Order struct {
	OrderId     int32     `json:"order_id" gorm:"primaryKey;autoIncrement"`
	CategoryId  int64     `json:"category_id" gorm:"not null"`
	BoatmodelId int64     `json:"boatmodel_id" gorm:"not null"`
	Uid         string    `json:"uid" gorm:"not null"`
	TotalPrice  float32   `json:"total_price" gorm:"not null"`
	BoatName    string    `json:"boat_name" gorm:"not null"`
	Color       string    `json:"color" gorm:"not null"`
	Created     time.Time `gorm:"autoCreateTime"`
	Updated     time.Time `gorm:"autoUpdateTime"`
	JobId       int32     `json:"job_id" gorm:"not null"`
	Comment     string    `json:"comment"`
}

func (Order) TableName() string {
	return "order"
}
