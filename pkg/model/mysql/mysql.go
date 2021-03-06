package mysql

import (
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	db *gorm.DB
)

func InitDB(url string) (*gorm.DB, error) {
	var (
		err error
	)
	db, err = setConnect(url)
	return db, err
}

func setConnect(url string) (*gorm.DB, error) {
	logger.Info("connecting mysql...., the url is: ", url)
	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	sqlDb, err := db.DB()

	if err != nil {
		return nil, err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDb.SetMaxIdleConns(utils.GetConfigInt("mysql.max_idle_conns"))

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDb.SetMaxOpenConns(utils.GetConfigInt("mysql.max_open_conns"))

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDb.SetConnMaxLifetime(time.Second * time.Duration(utils.GetConfigInt64("mysql.conn_max_lifetime")))

	err = sqlDb.Ping()
	if err != nil {
		logger.Info(err)
		return nil, err
	}

	return db, nil
}
