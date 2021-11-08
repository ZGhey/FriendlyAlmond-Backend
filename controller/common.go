package controller

import (
	"FriendlyAlmond_backend/pkg/logger"
	baseModel "FriendlyAlmond_backend/pkg/model"
	"FriendlyAlmond_backend/pkg/model/consulreg"
	pbConfig "FriendlyAlmond_backend/proto/configuration"
	pbJobModule "FriendlyAlmond_backend/proto/jobModule"
	pbLogin "FriendlyAlmond_backend/proto/login"
	pbOrder "FriendlyAlmond_backend/proto/order"
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

//register RPC service for controller use
var (
	rpcLoginService     pbLogin.LoginService
	rpcConfigService    pbConfig.ConfigurationService
	rpcOrderService     pbOrder.OrderService
	rpcJobModuleService pbJobModule.JobModuleService
)

//InitRpcCaller init the rpc service
func InitRpcCaller() {
	rpcLoginService = pbLogin.NewLoginService("login", consulreg.MicroSer.Client())
	rpcConfigService = pbConfig.NewConfigurationService("configuration", consulreg.MicroSer.Client())
	rpcOrderService = pbOrder.NewOrderService("order", consulreg.MicroSer.Client())
	rpcJobModuleService = pbJobModule.NewJobModuleService("job_module", consulreg.MicroSer.Client())
}
