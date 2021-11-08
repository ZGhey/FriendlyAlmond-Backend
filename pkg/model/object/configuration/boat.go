package configuration

import (
	"FriendlyAlmond_backend/pkg/model"
	pbConfig "FriendlyAlmond_backend/proto/configuration"
)

type QueryBoat struct {
	model.JSONResult
	Data []Boat `json:"data"`
}

type Boat struct {
	Id          int64  `json:"id" gorm:"not null;unique;primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"not null"`
	Size        int64  `json:"size"`
	Shape       string `json:"shape"`
	Description string `json:"description"`
	Created     int64  `gorm:"autoCreateTime"`
	Updated     int64  `gorm:"autoUpdateTime"`
}

//TableName set table name
func (Boat) TableName() string {
	return "boat"
}

//Pb2Normal convert protobuf objects to struct
func (q *QueryBoat) Pb2Normal(pbListBoat *pbConfig.ListBoat) {
	for _, value := range pbListBoat.Data {
		boat := new(Boat)
		boat.Id = value.Id
		boat.Name = value.Name
		boat.Shape = value.Shape
		boat.Description = value.Description
		q.Data = append(q.Data, *boat)
	}
}
