package test

import (
	"FriendlyAlmond_backend/pkg/daemon"
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/model/consulreg"
	"FriendlyAlmond_backend/pkg/model/mysql"
	"FriendlyAlmond_backend/pkg/utils"
	"fmt"
	"os"
	"testing"
)

//same config with the corresponding config file in conf doc
var (
	configStr = `
log_path: ./log/login
micro:
  addr: localhost
  name: configuration
  version: 1.0
mysql:
  user: root:friendly_almond_database@tcp(47.74.85.143:3306)/user?charset=utf8mb4&parseTime=True&loc=Local
  config_boat: root:friendly_almond_database@tcp(47.74.85.143:3306)/config_boat?charset=utf8mb4&parseTime=True&loc=Local
  order: root:friendly_almond_database@tcp(47.74.85.143:3306)/order?charset=utf8mb4&parseTime=True&loc=Local
  job_module: root:friendly_almond_database@tcp(47.74.85.143:3306)/job_module?charset=utf8mb4&parseTime=True&loc=Local
  max_idle_conns: 10
  max_open_conns: 100
  conn_max_lifetime: 60 #second
redis:
  addr: test.opsnft.net:49162
  pwd: opsnft_redis
  db: 0
  dial_timeout: 11
  read_timeout: 30
  write_timeout: 30
  pool_size: 20
  pool_timeout: 30
`
)

//TestMain before run the unit test, this method will run firstly.
func TestMain(m *testing.M) {
	fmt.Println("running test main")
	logger.InitDefault(&logger.Options{
		AbsLogDir: "",
		Debug:     true,
	})
	defer logger.Sync()

	if err := utils.LoadConfigString(configStr); err != nil {
		daemon.Exit(-1, err.Error())
	}

	//init outside resources
	err := mysql.InitMyDB()
	if err != nil {
		daemon.Exit(-1, err.Error())
	}

	//Create service
	err = consulreg.InitMicro("localhost", "configuration")
	if err != nil {
		daemon.Exit(-1, err.Error())
	}

	code := m.Run()

	logger.Info("test main running after")

	os.Exit(code)
}
