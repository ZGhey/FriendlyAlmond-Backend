package mysql

import (
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InitDB(url string) (*gorm.DB, error) {
	var (
		err error
	)
	once.Do(func() {
		db, err = setConnect(url)
	})
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
	//defer func(sqlDb *sql.DB) {
	//	err := sqlDb.Close()
	//	if err != nil {
	//		logger.Fatal(err)
	//	}
	//}(sqlDb)
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
