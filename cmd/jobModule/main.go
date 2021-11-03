package main

import (
	"FriendlyAlmond_backend/pkg/daemon"
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/model/consulreg"
	"FriendlyAlmond_backend/pkg/model/mysql"
	"FriendlyAlmond_backend/pkg/model/rdb"
	"FriendlyAlmond_backend/pkg/utils"
	"FriendlyAlmond_backend/pkg/version"
	pbJobModule "FriendlyAlmond_backend/proto/jobModule"
	"FriendlyAlmond_backend/service/jobModule/handler"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	flag := daemon.NewCmdFlags()
	if err := flag.Parse(); err != nil {
		daemon.Exit(-1, err.Error())
	}
	err := utils.LoadConfigFile(flag.ConfigFile)
	if err != nil {
		daemon.Exit(-1, err.Error())
	}

	loggerOpt := logger.NewOpts(utils.GetConfigStr("log_path"))
	logger.InitDefault(loggerOpt)
	defer logger.Sync()
	logger.Info(version.Long)

	//init outside resources
	err = mysql.InitMyDB()
	if err != nil {
		daemon.Exit(-1, err.Error())
	}

	err = rdb.InitRedis(&redis.Options{
		Addr:         utils.GetConfigStr("redis.addr"),
		Password:     utils.GetConfigStr("redis.pwd"),
		DB:           utils.GetConfigInt("redis.db"), // use default DB
		DialTimeout:  time.Duration(utils.GetConfigInt64("redis.dial_timeout")) * time.Second,
		ReadTimeout:  time.Duration(utils.GetConfigInt64("redis.read_timeout")) * time.Second,
		WriteTimeout: time.Duration(utils.GetConfigInt64("redis.write_timeout")) * time.Second,
		PoolSize:     utils.GetConfigInt("redis.pool_size"),
		PoolTimeout:  time.Duration(utils.GetConfigInt64("redis.pool_timeout")) * time.Second,
	})
	if err != nil {
		daemon.Exit(-1, err.Error())
	}

	//Create service
	err = consulreg.InitMicro(utils.GetConfigStr("micro.addr"), utils.GetConfigStr("micro.name"))
	if err != nil {
		daemon.Exit(-1, err.Error())
	}
	// Create service
	srv := consulreg.MicroSer
	err = pbJobModule.RegisterJobModuleHandler(srv.Server(), new(handler.JobModule))
	if err != nil {
		return
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
