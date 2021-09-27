package mysql

import (
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/model/object/login"
	"FriendlyAlmond_backend/pkg/utils"
	"gorm.io/gorm"
)

var (
	UserDB  *gorm.DB
	OrderDB *gorm.DB
	Staff   *gorm.DB
)

func InitMyDB() error {
	var err error
	UserDB, err = InitDB(utils.GetConfigStr("mysql.user"))
	OrderDB, err = InitDB(utils.GetConfigStr("mysql.order"))
	Staff, err = InitDB(utils.GetConfigStr("mysql.staff"))

	err = UserDB.AutoMigrate(&login.UserInfo{})
	if err != nil {
		logger.Info(err)
		return err
	}
	return err
}
