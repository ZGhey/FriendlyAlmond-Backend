package configuration

import (
	"FriendlyAlmond_backend/pkg/model"
	pbConfig "FriendlyAlmond_backend/proto/configuration"
	"time"
)

type QuerySection struct {
	model.JSONResult
	Data []RespSection `json:"data"`
}

type RespSection struct {
	CategoryName string    `json:"category_name"`
	Data         []Section `json:"data"`
}

type Section struct {
	Id         int64     `json:"id" gorm:"not null;unique;primaryKey;autoIncrement"`
	CategoryId int64     `json:"category_id" gorm:"not null"`
	Specs      string    `json:"specs"`
	Name       string    `json:"name" gorm:"not null"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	Price      float32   `json:"price" gorm:"not null"`
	Detail     string    `json:"detail"`
}

func (Section) TableName() string {
	return "section"
}

func (q *QuerySection) Pb2Normal(pbListSection *pbConfig.ListSection) {
	for _, pbValue := range pbListSection.Data {
		respSection := new(RespSection)
		section := new(Section)
		respSection.CategoryName = pbValue.CategoryName
		section.Id = pbValue.Id
		section.CategoryId = pbValue.CategoryId
		section.Name = pbValue.Name
		section.Price = pbValue.Price
		respSection.Data = append(respSection.Data, *section)
		if len(q.Data) == 0 {
			q.Data = append(q.Data, *respSection)
		} else if len(q.Data) > 0 {
			for index, queryValue := range q.Data {
				if respSection.CategoryName == queryValue.CategoryName {
					q.Data[index].Data = append(q.Data[index].Data, *section)
				} else if index == len(q.Data)-1 {
					q.Data = append(q.Data, *respSection)
				} else {
					continue
				}
			}
		}
	}
}
