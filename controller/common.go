package controller

import (
	"FriendlyAlmond_backend/pkg/logger"
	baseModel "FriendlyAlmond_backend/pkg/model"
	"FriendlyAlmond_backend/pkg/model/consulreg"
	pbConfig "FriendlyAlmond_backend/proto/configuration"
	pbLogin "FriendlyAlmond_backend/proto/login"
	"github.com/gin-gonic/gin"
	"net/http"
)

type jsonResult interface {
	GetMessage() string
	NewError(recode string) *baseModel.JSONResult
	NewSuccess() *baseModel.JSONResult
}

func responseHTTP(ctx *gin.Context, statusCode int, data jsonResult) {
	if statusCode != http.StatusOK {
		logger.Error(data.GetMessage())
	}
	ctx.JSON(statusCode, &data)
}

var (
	rpcLoginService  pbLogin.LoginService
	rpcConfigService pbConfig.ConfigurationService
)

func InitRpcCaller() {
	rpcLoginService = pbLogin.NewLoginService("login", consulreg.MicroSer.Client())
	rpcConfigService = pbConfig.NewConfigurationService("configuration", consulreg.MicroSer.Client())
}
