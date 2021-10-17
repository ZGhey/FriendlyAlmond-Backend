package configuration

import (
	"FriendlyAlmond_backend/pkg/model"
	pbConfig "FriendlyAlmond_backend/proto/configuration"
)

type QueryCategory struct {
	model.JSONResult
	Data []Category `json:"data"`
}

type Category struct {
	Id          int64  `json:"id" gorm:"not null;unique;primaryKey;autoIncrement"`
	Type        string `json:"type" gorm:"not null"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	Created     int64  `gorm:"autoCreateTime"`
	Updated     int64  `gorm:"autoUpdateTime"`
}

func (Category) TableName() string {
	return "category"
}

func (q *QueryCategory) Pb2Normal(pbListCategory *pbConfig.ListCategory) {
	for index, value := range pbListCategory.Data {
		q.Data[index].Id = value.Id
		q.Data[index].Type = value.Type
		q.Data[index].Name = value.Name
		q.Data[index].Description = value.Description
	}
}
