package mysql

import (
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/model/object/jobModule"
	"FriendlyAlmond_backend/pkg/model/object/login"
	"FriendlyAlmond_backend/pkg/model/object/order"
	"FriendlyAlmond_backend/pkg/utils"
	"gorm.io/gorm"
)

var (
	UserDB       *gorm.DB
	OrderDB      *gorm.DB
	JobModule    *gorm.DB
	ConfigBoatDB *gorm.DB
)

func InitMyDB() error {
	var err error
	UserDB, err = InitDB(utils.GetConfigStr("mysql.user"))
	OrderDB, err = InitDB(utils.GetConfigStr("mysql.order"))
	JobModule, err = InitDB(utils.GetConfigStr("mysql.job_module"))
	ConfigBoatDB, err = InitDB(utils.GetConfigStr("mysql.config_boat"))

	err = UserDB.AutoMigrate(&login.UserInfo{})
	if err != nil {
		logger.Info(err)
		return err
	}

	err = OrderDB.AutoMigrate(&order.Order{})
	if err != nil {
		logger.Info(err)
		return err
	}

	err = OrderDB.AutoMigrate(&order.Section{})
	if err != nil {
		logger.Info(err)
		return err
	}
	err = OrderDB.AutoMigrate(&order.Component{})
	if err != nil {
		logger.Info(err)
		return err
	}

	err = JobModule.AutoMigrate(&jobModule.Staff{})
	if err != nil {
		logger.Info(err)
		return err
	}

	err = JobModule.AutoMigrate(&jobModule.Job{})
	if err != nil {
		logger.Info(err)
		return err
	}

	err = JobModule.AutoMigrate(&jobModule.Task{})
	if err != nil {
		logger.Info(err)
		return err
	}

	return err
}
