package controller

import (
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/model"
	"FriendlyAlmond_backend/pkg/model/object/jobModule"
	"FriendlyAlmond_backend/pkg/model/object/login"
	"FriendlyAlmond_backend/pkg/model/object/order"
	"FriendlyAlmond_backend/pkg/utils"
	pbJobModule "FriendlyAlmond_backend/proto/jobModule"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateStaff godoc
// @Summary update staff info。
// @Description When POST this API, the API will update the staff info into DB
// @ID UpdateStaff
// @tags Job_Module
// @Accept  json
// @Produce  json
// @Param jobModule.Staff body jobModule.Staff true "account field is required"
// @Success 0 {object} model.JSONResult
// @Header 200 {header} string
// @Failure 4005 {object} model.JSONResult "The micro-service can't be reachable"
// @Failure 4125 {object} model.JSONResult "Generate the captcha failed"
// @Router /job/update-staff [post]
func UpdateStaff(ctx *gin.Context) {
	var (
		statusCode int
		req        jobModule.Staff
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

	pbReq := &pbJobModule.Staff{}
	pbReq.Password = req.Password
	pbReq.Account = req.Account
	pbReq.Email = req.Email
	pbReq.Firstname = req.Firstname
	pbReq.Middlename = req.Middlename
	pbReq.Lastname = req.Lastname
	pbReq.Phone = req.Phone
	pbReq.Address = req.Address
	pbReq.AreaCode = req.AreaCode
	pbReq.Skill = req.Skill
	remoteResult, err := rpcJobModuleService.UpdateStaff(context.TODO(), pbReq)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if remoteResult.StatusCode != utils.RECODE_OK {
		resp.SetError(remoteResult.StatusCode, "verify the captcha failed", err)
		statusCode = http.StatusBadRequest
		return
	}
	statusCode = http.StatusOK
	resp.NewSuccess()
}

// AddStaff godoc
// @Summary add a new staff。
// @Description When POST this API, the API will create a staff
// @ID AddStaff
// @tags Job_Module
// @Accept  json
// @Produce  json
// @Param jobModule.Staff body jobModule.Staff true "all the info for current staff"
// @Success 0 {object} model.JSONResult
// @Header 200 {header} string
// @Failure 4005 {object} model.JSONResult "The micro-service can't be reachable"
// @Failure 4125 {object} model.JSONResult "Generate the captcha failed"
// @Router /job/add-staff [post]
func AddStaff(ctx *gin.Context) {
	var (
		statusCode int
		req        jobModule.Staff
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

	pbReq := &pbJobModule.Staff{}
	pbReq.Password = req.Password
	pbReq.Account = req.Account
	pbReq.Email = req.Email
	pbReq.Firstname = req.Firstname
	pbReq.Middlename = req.Middlename
	pbReq.Lastname = req.Lastname
	pbReq.Phone = req.Phone
	pbReq.Address = req.Address
	pbReq.AreaCode = req.AreaCode
	pbReq.Skill = req.Skill
	remoteResult, err := rpcJobModuleService.AddStaff(context.TODO(), pbReq)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if remoteResult.StatusCode != utils.RECODE_OK {
		resp.SetError(remoteResult.StatusCode, "verify the captcha failed", err)
		statusCode = http.StatusBadRequest
		return
	}
	statusCode = http.StatusOK
	resp.NewSuccess()
}

// QueryStaff godoc
// @Summary query staff data。
// @Description When get this API, the API will return all staff info
// @ID QueryStaff
// @tags Job_Module
// @Accept  json
// @Produce  json
// @Success 0 {object} jobModule.ListStaff
// @Header 200 {header} string
// @Failure 4005 {object} jobModule.ListStaff "The micro-service can't be reachable"
// @Failure 4001 {object} jobModule.ListStaff "Database problem"
// @Router /job/query-staff [get]
func QueryStaff(ctx *gin.Context) {
	var (
		statusCode int
		resp       jobModule.ListStaff
		pbEmpty    pbJobModule.Empty
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	result, err := rpcJobModuleService.QueryListStaff(context.TODO(), &pbEmpty)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if result.StatusCode != utils.RECODE_OK {
		resp.SetError(result.StatusCode, result.StatusCode, err)
		statusCode = http.StatusBadRequest
		return
	}

	for _, value := range result.Staff {
		respStaff := new(jobModule.Staff)
		respStaff.Account = value.Account
		respStaff.Password = value.Password
		respStaff.Firstname = value.Firstname
		respStaff.Lastname = value.Lastname
		respStaff.Middlename = value.Middlename
		respStaff.Email = value.Email
		respStaff.Address = value.Address
		respStaff.Phone = value.Phone
		respStaff.Skill = value.Skill
		respStaff.AreaCode = value.AreaCode
		respStaff.StaffId = value.StaffId
		resp.Data = append(resp.Data, respStaff)
	}

	statusCode = http.StatusOK
	resp.NewSuccess()
}

// QueryUser godoc
// @Summary query user。
// @Description When get this API, the API will return all user info
// @ID QueryUser
// @tags Job_Module
// @Accept  json
// @Produce  json
// @Success 0 {object} jobModule.ListUserInfo
// @Header 200 {header} string
// @Failure 4005 {object} jobModule.ListUserInfo "The micro-service can't be reachable"
// @Failure 4001 {object} jobModule.ListUserInfo "Database problem"
// @Router /job/query-user [get]
func QueryUser(ctx *gin.Context) {
	var (
		statusCode int
		resp       jobModule.ListUserInfo
		pbEmpty    pbJobModule.Empty
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	result, err := rpcJobModuleService.QueryListUser(context.TODO(), &pbEmpty)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if result.StatusCode != utils.RECODE_OK {
		resp.SetError(result.StatusCode, result.StatusCode, err)
		statusCode = http.StatusBadRequest
		return
	}

	for _, value := range result.UserInfo {
		respOrder := new(login.UserInfo)
		respOrder.Uid = value.Uid
		respOrder.FirstName = value.FirstName
		respOrder.MiddleName = value.MiddleName
		respOrder.LastName = value.LastName
		respOrder.Email = value.Email
		respOrder.Password = value.Password
		respOrder.Phone = value.Phone
		respOrder.Address = value.Address
		respOrder.AreaCode = value.AreaCode
		resp.Data = append(resp.Data, respOrder)
	}

	statusCode = http.StatusOK
	resp.NewSuccess()
}

// QueryNoJobOrder godoc
// @Summary return the order haven't a job yet。
// @Description When get this API, the API will return the order haven't assigned a job
// @ID QueryNoJobOrder
// @tags Job_Module
// @Accept  json
// @Produce  json
// @Success 0 {object} jobModule.ListUserInfo
// @Header 200 {header} string
// @Failure 4005 {object} jobModule.ListUserInfo "The micro-service can't be reachable"
// @Failure 4001 {object} jobModule.ListUserInfo "Database problem"
// @Router /job/query-order [get]
func QueryNoJobOrder(ctx *gin.Context) {
	var (
		statusCode int
		resp       order.ApiOrder
		pbEmpty    pbJobModule.Empty
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	result, err := rpcJobModuleService.QueryNoJobOrder(context.TODO(), &pbEmpty)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if result.StatusCode != utils.RECODE_OK {
		resp.SetError(result.StatusCode, result.StatusCode, err)
		statusCode = http.StatusBadRequest
		return
	}

	for _, value := range result.QueryOrder {
		respOrder := new(order.RespOrder)
		respOrder.OrderId = value.OrderId
		respOrder.Color = value.Color
		respOrder.CategoryName = value.CategoryName
		respOrder.TotalPrice = value.TotalPrice
		respOrder.BoatName = value.BoatName
		respOrder.SectionId = value.SectionId
		respOrder.ComponentId = value.ComponentId
		respOrder.BoatmodelName = value.BoatmodelName
		respOrder.UserName = value.UserName
		respOrder.OrderDate = value.OrderDate
		resp.Data = append(resp.Data, respOrder)
	}

	statusCode = http.StatusOK
	resp.NewSuccess()
}

// CreateJob godoc
// @Summary create job。
// @Description When POST this API, the API will create a job
// @ID CreateJob
// @tags Job_Module
// @Accept  json
// @Produce  json
// @Param jobModule.Job body jobModule.Job true "just need the order id"
// @Success 0 {object} model.JSONResult
// @Header 200 {header} string
// @Failure 4005 {object} model.JSONResult "The micro-service can't be reachable"
// @Router /job/create-job [post]
func CreateJob(ctx *gin.Context) {
	var (
		statusCode int
		req        jobModule.Job
		resp       model.JSONResult
		pbReq      pbJobModule.Job
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		statusCode = http.StatusOK
		resp.SetError(utils.RECODE_DATAERR, "the data may have some problem", err)
		return
	}
	pbReq.OrderId = req.OrderId
	remoteResult, err := rpcJobModuleService.CreateJob(context.TODO(), &pbReq)
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

// CreateTask godoc
// @Summary create task。
// @Description When POST this API, the API will create a task
// @ID CreateTask
// @tags Job_Module
// @Accept  json
// @Produce  json
// @Param jobModule.Task body jobModule.Task true "just need the job id"
// @Success 0 {object} model.JSONResult
// @Header 200 {header} string
// @Failure 4005 {object} model.JSONResult "The micro-service can't be reachable"
// @Router /job/create-task [post]
func CreateTask(ctx *gin.Context) {
	var (
		statusCode int
		req        jobModule.Task
		resp       model.JSONResult
		pbReq      pbJobModule.Task
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		statusCode = http.StatusOK
		resp.SetError(utils.RECODE_DATAERR, "the data may have some problem", err)
		return
	}
	pbReq.JobId = req.JobId
	pbReq.SectionId = req.SectionId
	pbReq.ComponentId = req.ComponentId
	pbReq.Description = req.Description
	pbReq.StartDate = req.StartDate
	pbReq.DueDate = req.DueDate
	pbReq.StaffId = req.StaffId
	remoteResult, err := rpcJobModuleService.CreateTask(context.TODO(), &pbReq)
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

// QueryTask godoc
// @Summary query task。
// @Description When POST this API, the API will all the task in a job
// @ID QueryTask
// @tags Job_Module
// @Accept  json
// @Produce  json
// @Param jobModule.Task body jobModule.Task true "just need the job id"
// @Success 0 {object} jobModule.RespTask
// @Header 200 {header} string
// @Failure 4005 {object} jobModule.RespTask "The micro-service can't be reachable"
// @Router /job/query-task [post]
func QueryTask(ctx *gin.Context) {
	var (
		statusCode int
		req        jobModule.Task
		resp       jobModule.RespTask
		pbReq      pbJobModule.Task
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		statusCode = http.StatusOK
		resp.SetError(utils.RECODE_DATAERR, "the data may have some problem", err)
		return
	}
	pbReq.JobId = req.JobId
	remoteResult, err := rpcJobModuleService.QueryTask(context.TODO(), &pbReq)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if remoteResult.StatusCode != utils.RECODE_OK {
		resp.SetError(remoteResult.StatusCode, utils.RecodeTest(remoteResult.StatusCode), err)
		statusCode = http.StatusBadRequest
		return
	}
	for _, value := range remoteResult.Task {
		task := new(jobModule.Task)
		task.JobId = value.JobId
		task.TaskId = value.TaskId
		task.SectionId = value.SectionId
		task.ComponentId = value.ComponentId
		task.Description = value.Description
		task.StartDate = value.StartDate
		task.DueDate = value.DueDate
		resp.Data = append(resp.Data, *task)
	}
	statusCode = http.StatusOK
	resp.NewSuccess()
}

// MostPopular godoc
// @Summary get the most popular things。
// @Description When POST this API, the API will return the most popular things. like color, component, section.
// @ID MostPopular
// @tags Job_Module
// @Accept  json
// @Produce  json
// @Success 0 {object} jobModule.RespMostPopular
// @Header 200 {header} string
// @Failure 4005 {object} jobModule.RespMostPopular "The micro-service can't be reachable"
// @Router /job/most-popular [get]
func MostPopular(ctx *gin.Context) {
	var (
		statusCode int
		resp       jobModule.RespMostPopular
		pbReq      pbJobModule.Empty
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	remoteResult, err := rpcJobModuleService.QueryMostPopular(context.TODO(), &pbReq)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if remoteResult.StatusCode != utils.RECODE_OK {
		resp.SetError(remoteResult.StatusCode, utils.RecodeTest(remoteResult.StatusCode), err)
		statusCode = http.StatusBadRequest
		return
	}
	marshal, err := json.Marshal(remoteResult)
	if err != nil {
		logger.Error(err)
		return
	}
	err = json.Unmarshal(marshal, &resp.Data)
	if err != nil {
		logger.Error(err)
		return
	}
	statusCode = http.StatusOK
	resp.NewSuccess()
}

// TotalSales godoc
// @Summary get the total sales
// @Description When POST this API, the API will return the for last 1, 3, 6 month and 1 year.
// @ID TotalSales
// @tags Job_Module
// @Accept  json
// @Produce  json
// @Param jobModule.Task body jobModule.Task true "just need the job id"
// @Success 0 {object} jobModule.RespTask
// @Header 200 {header} string
// @Failure 4005 {object} jobModule.RespTask "The micro-service can't be reachable"
// @Router /job/total-sales [get]
func TotalSales(ctx *gin.Context) {
	var (
		statusCode int
		resp       jobModule.RespTotalSales
		pbReq      pbJobModule.Empty
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	remoteResult, err := rpcJobModuleService.QueryTotalSales(context.TODO(), &pbReq)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if remoteResult.StatusCode != utils.RECODE_OK {
		resp.SetError(remoteResult.StatusCode, utils.RecodeTest(remoteResult.StatusCode), err)
		statusCode = http.StatusBadRequest
		return
	}
	resp.Data.LastOneMonth = remoteResult.LastOneMonth
	resp.Data.LastThreeMonth = remoteResult.LastThreeMonth
	resp.Data.LastSixMonth = remoteResult.LastSixMonth
	resp.Data.LastOneYear = remoteResult.LastOneYear
	statusCode = http.StatusOK
	resp.NewSuccess()
}
