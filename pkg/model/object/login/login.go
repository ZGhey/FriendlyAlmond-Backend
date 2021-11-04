package login

import (
	"FriendlyAlmond_backend/pkg/model"
)

type Captcha struct {
	model.JSONResult
	Data CaptchaInfo `json:"data"`
}
type CaptchaInfo struct {
	Id     string `json:"id"`
	Image  string `json:"image"`
	Answer string `json:"answer"`
}
type RegisterUser struct {
	UserInfo
	Captcha CaptchaInfo `gorm:"-" json:"captcha"`
}
type QueryUser struct {
	model.JSONResult
	Data UserInfo `json:"data"`
}
type UpdateUser struct {
	UserInfo
}
type RespUser struct {
	model.JSONResult
	Data UserInfo `json:"data"`
}

type UserInfo struct {
	FirstName  string `json:"first_name" gorm:"not null"`
	LastName   string `json:"last_name" gorm:"not null"`
	MiddleName string `json:"middle_name" gorm:"not null"`
	Password   string `gorm:"not null" json:"password"`
	Email      string `gorm:"not null" json:"email" bind:"required"`
	Uid        string `gorm:"not null;unique;primaryKey"`
	Created    int64  `gorm:"autoCreateTime"`
	Updated    int64  `gorm:"autoUpdateTime"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	AreaCode   string `gorm:"not null;" json:"area_code"`
	IsAdmin    bool   `json:"is_admin"`
	StaffId    int32  `json:"staff_id"`
	Skill      string `json:"skill"`
	Account    string `json:"account"`
}
