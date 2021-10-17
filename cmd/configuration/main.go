package main

import (
	"FriendlyAlmond_backend/pkg/daemon"
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/model/consulreg"
	"FriendlyAlmond_backend/pkg/model/mysql"
	"FriendlyAlmond_backend/pkg/utils"
	"FriendlyAlmond_backend/pkg/version"
	pbConfig "FriendlyAlmond_backend/proto/configuration"
	"FriendlyAlmond_backend/service/configuration/handler"
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
	err = pbConfig.RegisterConfigurationHandler(srv.Server(), new(handler.Config))
	if err != nil {
		return
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
