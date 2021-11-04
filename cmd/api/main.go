package main

import (
	"FriendlyAlmond_backend/cmd/api/docs"
	"FriendlyAlmond_backend/pkg/daemon"
	"FriendlyAlmond_backend/pkg/model/consulreg"
	"net/http"
	"strings"
	"time"

	"FriendlyAlmond_backend/controller"
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/model"
	"FriendlyAlmond_backend/pkg/model/jwt"
	"FriendlyAlmond_backend/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title FriendlyAlmond API Interface
// @version 1.0
// @description This the FriendlyAlmond backend API for Frontend.
// @termsOfService http://swagger.io/terms/

// @contact.name Gong Zhang
// @contact.url http://www.swagger.io/support
// @contact.email gong.zhang10@uon.edu.au

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1.0
// @query.collection.format multi
// @schemes http

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

//router used for unit test
func ginRouter(middleWare gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	// swagger
	docs.SwaggerInfo.Host = utils.GetConfigStr("tls_domain") + ":" + utils.GetConfigStr("port")
	url := ginSwagger.URL(utils.GetConfigStr("swagger_doc_url"))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//define middleware
	router.Use(ginLogger())
	router.Use(CORS())

	r1 := router.Group("api/v1.0/login")
	{
		r1.GET("/get-captcha", controller.GetCaptcha)
		r1.POST("/register", controller.Register)
		r1.POST("/update", controller.Update)
		r1.POST("/query", controller.Query)
		r1.POST("/login", controller.Login)
	}

	r2 := router.Group("api/v1.0/config")
	{
		r2.POST("/query-boat", controller.QueryBoat)
		r2.POST("/query-component", controller.QueryComponent)
		r2.POST("/query-section", controller.QuerySection)
		r2.POST("/query-component-id", controller.QueryComponentById)
		r2.POST("/query-section-id", controller.QuerySectionById)
	}

	r3 := router.Group("api/v1.0/order")
	{
		r3.POST("/create", controller.CreateOrder)
		r3.POST("/query", controller.QueryOrder)
	}

	r4 := router.Group("api/v1.0/job")
	{
		r4.POST("update-staff", controller.UpdateStaff)
		r4.POST("add-staff", controller.AddStaff)
		r4.GET("query-staff", controller.QueryStaff)
		r4.GET("query-user", controller.QueryUser)
		r4.GET("query-order", controller.QueryNoJobOrder)
		r4.POST("create-job", controller.CreateJob)
		r4.POST("create-task", controller.CreateTask)
		r4.POST("query-task", controller.QueryTask)
		r4.GET("most-popular", controller.MostPopular)
		r4.GET("total-sales", controller.TotalSales)
	}
	//if middleWare != nil {
	//	router.Use(middleWare)
	//}

	return router
}

func main() {
	flag := daemon.NewCmdFlags()
	if err := flag.Parse(); err != nil {
		panic(err)
	}
	err := utils.LoadConfigFile(flag.ConfigFile)
	if err != nil {
		panic(err)
	}

	loggerOpt := logger.NewOpts(utils.GetConfigStr("log_path"))
	if gin.Mode() != gin.ReleaseMode {
		loggerOpt.Debug = true
	}

	logger.InitDefault(loggerOpt)
	defer logger.Sync()

	//Create service
	err = consulreg.InitMicro(utils.GetConfigStr("micro.addr"), "fa_api")
	if err != nil {
		daemon.Exit(-1, err.Error())
	}
	controller.InitRpcCaller()

	router := ginRouter(JWTAuth())

	logger.Info(gin.Version)
	port := utils.GetConfigInt("port")
	var builder strings.Builder
	builder.WriteString(":")
	builder.WriteString(strconv.Itoa(port))
	err = router.Run(builder.String())
	if err != nil {
		logger.Fatal(err.Error())
	}
}

// JWTAuth middleware for check token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var result model.JSONResult
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, result.NewError(utils.RECODE_TOKENERR))
			c.Abort()
			return
		}

		j := jwt.NewJWT()
		// parse token
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				c.JSON(http.StatusOK, result.NewError(utils.RECODE_TOKENEXPIRED))
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// continue pass to next router
		c.Set("claims", claims)
	}
}

func ginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := utils.Now()
		// before request
		c.Next()
		// after request
		latency := time.Since(t)

		// request way
		reqMethod := c.Request.Method

		// request router
		reqUrl := c.Request.RequestURI

		// status code
		statusCode := c.Writer.Status()

		// request ip address
		clientIP := c.ClientIP()

		logger.APIAccess(reqMethod, clientIP, reqUrl, statusCode, c.Request.ContentLength,
			int64(c.Writer.Size()), t, latency)
	}

}

func CORS() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, token")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
			return
		}

		context.Next()
	}
}
