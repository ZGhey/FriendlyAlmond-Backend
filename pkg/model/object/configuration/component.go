package configuration

import (
	"FriendlyAlmond_backend/pkg/model"
	pbConfig "FriendlyAlmond_backend/proto/configuration"
)

type QueryComponent struct {
	model.JSONResult
	Data []RespComponent `json:"data"`
}

type RespComponent struct {
	CategoryName string      `json:"category_name"`
	Data         []Component `json:"data"`
}

type Component struct {
	Id         int64   `json:"id" gorm:"not null;unique;primaryKey;autoIncrement"`
	Name       string  `json:"name" gorm:"not null"`
	Details    string  `json:"details" gorm:"not null"`
	SupplierId int64   `json:"supplier_id" gorm:"not null"`
	CategoryId int64   `json:"category_id" gorm:"not null"`
	Price      float32 `json:"price" gorm:"not null"`
	Created    int64   `gorm:"autoCreateTime"`
	Updated    int64   `gorm:"autoUpdateTime"`
}

func (Component) TableName() string {
	return "component"
}

func (q *QueryComponent) Pb2Normal(pbListComponent *pbConfig.ListComponent) {
	for _, pbValue := range pbListComponent.Data {
		respComponent := new(RespComponent)
		component := new(Component)
		respComponent.CategoryName = pbValue.CategoryName
		component.Id = pbValue.Id
		component.CategoryId = pbValue.CategoryId
		component.Name = pbValue.Name
		component.Details = pbValue.Details
		component.SupplierId = pbValue.SupplierId
		component.Price = pbValue.Price
		respComponent.Data = append(respComponent.Data, *component)
		if len(q.Data) == 0 {
			q.Data = append(q.Data, *respComponent)
		} else if len(q.Data) > 0 {
			for index, queryValue := range q.Data {
				if respComponent.CategoryName == queryValue.CategoryName {
					q.Data[index].Data = append(q.Data[index].Data, *component)
				} else if index == len(q.Data)-1 {
					q.Data = append(q.Data, *respComponent)
				} else {
					continue
				}
			}
		}
	}
}
