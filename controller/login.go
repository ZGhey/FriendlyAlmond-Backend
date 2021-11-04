package controller

import (
	"FriendlyAlmond_backend/pkg/model"
	"FriendlyAlmond_backend/pkg/model/object/login"
	"FriendlyAlmond_backend/pkg/utils"
	pbJobModule "FriendlyAlmond_backend/proto/jobModule"
	pbLogin "FriendlyAlmond_backend/proto/login"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCaptcha godoc
// @Summary Get captcha。
// @Description When GET this API, the API will return a captcha to frontend
// @ID GetCaptcha
// @tags Login
// @Accept  json
// @Produce  json
// @Success 0 {object} login.Captcha
// @Header 200 {header} string
// @Failure 4005 {object} login.Captcha "The micro-service can't be reachable"
// @Failure 4125 {object} login.Captcha "Generate the captcha failed"
// @Router /login/get-captcha [get]
func GetCaptcha(ctx *gin.Context) {
	var (
		statusCode int
		resp       login.Captcha
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	result, err := rpcLoginService.GenerateCaptcha(context.TODO(), &pbLogin.Empty{})
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if result.StatusCode == utils.RECODE_CAPTCHA_GENERATEERR {
		resp.SetError(result.StatusCode, "Generate the captcha failed", err)
		statusCode = http.StatusBadRequest
		return
	}
	resp.Data.Id = result.Id
	resp.Data.Image = result.Image
	statusCode = http.StatusOK
	resp.NewSuccess()
}

// Register godoc
// @Summary store user info。
// @Description When POST this API, the API will store the user info into DB
// @ID Register
// @tags Login
// @Accept  json
// @Produce  json
// @Param token body login.RegisterUser true "-"
// @Success 0 {object} model.JSONResult
// @Header 200 {header} string
// @Failure 4005 {object} model.JSONResult "The micro-service can't be reachable"
// @Failure 4125 {object} model.JSONResult "Generate the captcha failed"
// @Router /login/register [post]
func Register(ctx *gin.Context) {
	var (
		statusCode int
		req        login.RegisterUser
		resp       model.JSONResult
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		statusCode = http.StatusOK
		resp.SetError(utils.RECODE_DATAERR, "the data may have some problem", err)
		return
	}

	if result := utils.IsEmailValid(req.Email); !result {
		statusCode = http.StatusOK
		resp.SetError(utils.RECODE_DATAERR, "This is not a email address", nil)
		return
	}

	pbReq := &pbLogin.UserInfo{}
	pbCaptcha := &pbLogin.Captcha{}
	pbReq.Password = req.Password
	pbReq.Uid = utils.NewLen(10)
	pbReq.Email = req.Email
	pbReq.FirstName = req.FirstName
	pbReq.MiddleName = req.MiddleName
	pbReq.LastName = req.LastName
	pbReq.AreaCode = req.AreaCode
	pbReq.Phone = req.Phone
	pbReq.Address = req.Address
	pbCaptcha.Id = req.Captcha.Id
	pbCaptcha.Answer = req.Captcha.Answer
	pbReq.Captcha = pbCaptcha

	isSameEmail, err := rpcLoginService.IsSameEmail(context.TODO(), pbReq)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if isSameEmail.StatusCode != utils.RECODE_OK {
		resp.SetError(isSameEmail.StatusCode, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	}

	remoteResult, err := rpcLoginService.Register(context.TODO(), pbReq)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if remoteResult.StatusCode != utils.RECODE_OK {
		resp.SetError(remoteResult.StatusCode, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	}
	statusCode = http.StatusOK
	resp.NewSuccess()
}

// Update godoc
// @Summary update user info。
// @Description When POST this API, the API will update the user info into DB
// @ID Update
// @tags Login
// @Accept  json
// @Produce  json
// @Param user_info body login.UserInfo true "-"
// @Success 0 {object} model.JSONResult
// @Header 200 {header} string
// @Failure 4005 {object} model.JSONResult "The micro-service can't be reachable"
// @Failure 4125 {object} model.JSONResult "Generate the captcha failed"
// @Router /login/update [post]
func Update(ctx *gin.Context) {
	var (
		statusCode int
		req        login.UpdateUser
		resp       model.JSONResult
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		statusCode = http.StatusOK
		resp.SetError(utils.RECODE_DATAERR, "the data may have some problem", err)
		return
	}

	if req.Email != "" {
		if result := utils.IsEmailValid(req.Email); !result {
			statusCode = http.StatusOK
			resp.SetError(utils.RECODE_DATAERR, "This is not a email address", nil)
			return
		}
	}

	pbReq := &pbLogin.UserInfo{}
	pbReq.Password = req.Password
	pbReq.Uid = req.Uid
	pbReq.Email = req.Email
	pbReq.FirstName = req.FirstName
	pbReq.MiddleName = req.MiddleName
	pbReq.LastName = req.LastName
	pbReq.Phone = req.Phone
	pbReq.Address = req.Address
	pbReq.AreaCode = req.AreaCode

	isSameEmail, err := rpcLoginService.IsSameEmail(context.TODO(), pbReq)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if isSameEmail.StatusCode != utils.RECODE_OK {
		resp.SetError(isSameEmail.StatusCode, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	}

	remoteResult, err := rpcLoginService.Update(context.TODO(), pbReq)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if remoteResult.StatusCode != utils.RECODE_OK {
		resp.SetError(remoteResult.StatusCode, utils.RecodeTest(remoteResult.StatusCode), err)
		statusCode = http.StatusBadRequest
		return
	}
	statusCode = http.StatusOK
	resp.NewSuccess()
}

// Query godoc
// @Summary query user info。
// @Description When POST this API, the API will query the user info into DB
// @ID Query
// @tags Login
// @Accept  json
// @Produce  json
// @Param uid body login.UserInfo true "-"
// @Success 0 {object} model.JSONResult
// @Header 200 {header} string
// @Failure 4005 {object} model.JSONResult "The micro-service can't be reachable"
// @Failure 4125 {object} model.JSONResult "Generate the captcha failed"
// @Router /login/query [post]
func Query(ctx *gin.Context) {
	var (
		statusCode int
		req        login.UserInfo
		resp       login.QueryUser
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		statusCode = http.StatusOK
		resp.SetError(utils.RECODE_DATAERR, "the data may have some problem", err)
		return
	}

	pbReq := &pbLogin.UserInfo{}
	pbReq.Uid = req.Uid
	remoteResult, err := rpcLoginService.Query(context.TODO(), pbReq)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if remoteResult.OperationResult.StatusCode != utils.RECODE_OK {
		resp.SetError(remoteResult.OperationResult.StatusCode, "error", err)
		statusCode = http.StatusBadRequest
		return
	}
	statusCode = http.StatusOK
	resp.Data.Address = remoteResult.Address
	resp.Data.Uid = remoteResult.Uid
	resp.Data.Phone = remoteResult.Phone
	resp.Data.Email = remoteResult.Email
	resp.Data.FirstName = remoteResult.FirstName
	resp.Data.MiddleName = remoteResult.MiddleName
	resp.Data.LastName = remoteResult.LastName
	resp.Data.Password = remoteResult.Password
	resp.Data.AreaCode = remoteResult.AreaCode
	resp.NewSuccess()
}

// Login godoc
// @Summary login。
// @Description When POST this API, the API will verify the captcha, email and password
// @ID Login
// @tags Login
// @Accept  json
// @Produce  json
// @Param RegisterUser body login.RegisterUser true "-"
// @Success 0 {object} login.RespUser
// @Header 200 {header} string
// @Failure 4005 {object} login.RespUser "The micro-service can't be reachable"
// @Failure 4125 {object} login.RespUser "Generate the captcha failed"
// @Router /login/login [post]
func Login(ctx *gin.Context) {
	var (
		statusCode int
		req        login.RegisterUser
		resp       login.RespUser
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		statusCode = http.StatusOK
		resp.SetError(utils.RECODE_DATAERR, "the data may have some problem", err)
		return
	}

	if result := utils.IsEmailValid(req.Email); !result {
		pbReq := &pbJobModule.Staff{}
		pbCaptcha := &pbJobModule.Captcha{}
		pbReq.Password = req.Password
		pbReq.Account = req.Email
		pbCaptcha.Id = req.Captcha.Id
		pbCaptcha.Answer = req.Captcha.Answer
		pbReq.Captcha = pbCaptcha
		remoteResult, err := rpcJobModuleService.Login(context.TODO(), pbReq)
		if err != nil {
			resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
			statusCode = http.StatusBadRequest
			return
		} else if remoteResult.OperationResult.StatusCode != utils.RECODE_OK {
			resp.SetError(remoteResult.OperationResult.StatusCode, utils.RecodeTest(remoteResult.OperationResult.StatusCode), err)
			statusCode = http.StatusBadRequest
			return
		}
		statusCode = http.StatusOK
		resp.NewSuccess()
		resp.Data.IsAdmin = true
		resp.Data.Address = remoteResult.Address
		resp.Data.AreaCode = remoteResult.AreaCode
		resp.Data.Phone = remoteResult.Phone
		resp.Data.Email = remoteResult.Email
		resp.Data.StaffId = remoteResult.StaffId
		resp.Data.Password = remoteResult.Password
		resp.Data.FirstName = remoteResult.Firstname
		resp.Data.LastName = remoteResult.Lastname
		resp.Data.MiddleName = remoteResult.Middlename
		resp.Data.Skill = remoteResult.Skill
		resp.Data.Account = remoteResult.Account
		return
	}

	pbReq := &pbLogin.UserInfo{}
	pbCaptcha := &pbLogin.Captcha{}
	pbReq.Password = req.Password
	pbReq.Email = req.Email
	pbCaptcha.Id = req.Captcha.Id
	pbCaptcha.Answer = req.Captcha.Answer
	pbReq.Captcha = pbCaptcha
	remoteResult, err := rpcLoginService.Login(context.TODO(), pbReq)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if remoteResult.OperationResult.StatusCode != utils.RECODE_OK {
		resp.SetError(remoteResult.OperationResult.StatusCode, utils.RecodeTest(remoteResult.OperationResult.StatusCode), err)
		statusCode = http.StatusBadRequest
		return
	}
	statusCode = http.StatusOK
	resp.NewSuccess()
	resp.Data.IsAdmin = false
	resp.Data.Uid = remoteResult.Uid
	resp.Data.FirstName = remoteResult.FirstName
	resp.Data.Password = remoteResult.Password
	resp.Data.Email = remoteResult.Email
	resp.Data.Uid = remoteResult.Uid
	resp.Data.Phone = remoteResult.Phone
	resp.Data.Address = remoteResult.Address
	resp.Data.MiddleName = remoteResult.MiddleName
	resp.Data.LastName = remoteResult.LastName
	resp.Data.AreaCode = remoteResult.AreaCode
}
