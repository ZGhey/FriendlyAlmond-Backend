package jobModule

import (
	"FriendlyAlmond_backend/pkg/model"
	"FriendlyAlmond_backend/pkg/model/object/login"
	"time"
)

type ListUserInfo struct {
	model.JSONResult
	Data []*login.UserInfo `json:"data"`
}

type ListStaff struct {
	model.JSONResult
	Data []*Staff `json:"data"`
}

type Staff struct {
	StaffId    int32     `json:"staff_id" gorm:"primaryKey;autoIncrement"`
	Account    string    `json:"account" gorm:"not null"`
	Password   string    `json:"password" gorm:"not null"`
	Firstname  string    `json:"firstname" gorm:"not null"`
	Lastname   string    `json:"lastname" gorm:"not null"`
	Middlename string    `json:"middlename" gorm:"not null"`
	Email      string    `json:"email" gorm:"not null"`
	Address    string    `json:"address" gorm:"not null"`
	Phone      string    `json:"phone" gorm:"not null"`
	Skill      string    `json:"skill" gorm:"not null"`
	AreaCode   string    `json:"area_code" gorm:"not null"`
	Created    time.Time `gorm:"autoCreateTime"`
	Updated    time.Time `gorm:"autoUpdateTime"`
}

func (Staff) TableName() string {
	return "staff"
}
