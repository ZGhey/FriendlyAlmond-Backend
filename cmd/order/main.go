package main

import (
	"FriendlyAlmond_backend/pkg/daemon"
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/model/consulreg"
	"FriendlyAlmond_backend/pkg/model/mysql"
	"FriendlyAlmond_backend/pkg/utils"
	"FriendlyAlmond_backend/pkg/version"
	pbOrder "FriendlyAlmond_backend/proto/order"
	"FriendlyAlmond_backend/service/order/handler"
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

	//Create service
	err = consulreg.InitMicro(utils.GetConfigStr("micro.addr"), utils.GetConfigStr("micro.name"))
	if err != nil {
		daemon.Exit(-1, err.Error())
	}
	// Create service
	srv := consulreg.MicroSer
	err = pbOrder.RegisterOrderHandler(srv.Server(), new(handler.Order))
	if err != nil {
		return
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
