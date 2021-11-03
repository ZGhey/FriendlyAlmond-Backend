package order

import "time"

type Section struct {
	OrderId   int32     `json:"order_id" gorm:"not null"`
	SectionId int32     `json:"section_id" gorm:"not null"`
	Created   time.Time `gorm:"autoCreateTime"`
	Updated   time.Time `gorm:"autoUpdateTime"`
	TaskId    int32     `gorm:"not null"`
}

func (Section) TableName() string {
	return "order_section"
}

type Component struct {
	OrderId     int32     `json:"order_id" gorm:"not null"`
	ComponentId int32     `json:"component_id" gorm:"not null"`
	Created     time.Time `gorm:"autoCreateTime"`
	Updated     time.Time `gorm:"autoUpdateTime"`
	TaskId      int32     `gorm:"not null"`
}

func (Component) TableName() string {
	return "order_component"
}
