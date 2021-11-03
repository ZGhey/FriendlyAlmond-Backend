package controller

import (
	config "FriendlyAlmond_backend/pkg/model/object/configuration"
	"FriendlyAlmond_backend/pkg/utils"
	pbConfig "FriendlyAlmond_backend/proto/configuration"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// QueryBoat godoc
// @Summary Get boat model。
// @Description When post this API, the API will return a list that shows the boat info
// @ID QueryBoat
// @tags Configuration
// @Accept  json
// @Produce  json
// @Success 0 {object} configuration.QueryBoat
// @Header 200 {header} string
// @Failure 4005 {object} configuration.QueryBoat "The micro-service can't be reachable"
// @Failure 4001 {object} configuration.QueryBoat "Database problem"
// @Router /config/query-boat [post]
func QueryBoat(ctx *gin.Context) {
	var (
		statusCode int
		resp       config.QueryBoat
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	result, err := rpcConfigService.QueryBoat(context.TODO(), &pbConfig.Empty{})
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if result.StatusCode == utils.RECODE_DBERR {
		resp.SetError(result.StatusCode, "The database has some problems", err)
		statusCode = http.StatusBadRequest
		return
	}

	statusCode = http.StatusOK
	resp.Pb2Normal(result)
	resp.NewSuccess()
}

// QueryComponent godoc
// @Summary Get component。
// @Description When post this API, the API will return a list that shows the component info
// @ID QueryComponent
// @tags Configuration
// @Accept  json
// @Produce  json
// @Param type body configuration.Category true "just need the type of the category, like Motor"
// @Success 0 {object} configuration.QueryComponent
// @Header 200 {header} string
// @Failure 4005 {object} configuration.QueryComponent "The micro-service can't be reachable"
// @Failure 4001 {object} configuration.QueryComponent "Database problem"
// @Router /config/query-component [post]
func QueryComponent(ctx *gin.Context) {
	var (
		req        config.Category
		statusCode int
		resp       config.QueryComponent
		pbReq      pbConfig.Category
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.SetError(utils.RECODE_DATAERR, utils.RecodeTest(utils.RECODE_DATAERR), err)
		statusCode = http.StatusBadRequest
	}
	pbReq.Type = req.Type
	result, err := rpcConfigService.QueryComponent(context.TODO(), &pbReq)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if result.StatusCode == utils.RECODE_DBERR {
		resp.SetError(result.StatusCode, "The database has some problems", err)
		statusCode = http.StatusBadRequest
		return
	}

	statusCode = http.StatusOK
	resp.Pb2Normal(result)
	resp.NewSuccess()
}

// QuerySection godoc
// @Summary Get section。
// @Description When post this API, the API will return a list that shows the section info
// @ID QuerySection
// @tags Configuration
// @Accept  json
// @Produce  json
// @Param type body configuration.Category true "just need the type of the category, like Exterior"
// @Success 0 {object} configuration.QuerySection
// @Header 200 {header} string
// @Failure 4005 {object} configuration.QuerySection "The micro-service can't be reachable"
// @Failure 4001 {object} configuration.QuerySection "Database problem"
// @Router /config/query-section [post]
func QuerySection(ctx *gin.Context) {
	var (
		req        config.Category
		statusCode int
		resp       config.QuerySection
		pbReq      pbConfig.Category
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.SetError(utils.RECODE_DATAERR, utils.RecodeTest(utils.RECODE_DATAERR), err)
		statusCode = http.StatusBadRequest
	}
	pbReq.Type = req.Type
	result, err := rpcConfigService.QuerySection(context.TODO(), &pbReq)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if result.StatusCode == utils.RECODE_DBERR {
		resp.SetError(result.StatusCode, "The database has some problems", err)
		statusCode = http.StatusBadRequest
		return
	}

	statusCode = http.StatusOK
	resp.Pb2Normal(result)
	resp.NewSuccess()
}

// QueryComponentById godoc
// @Summary Get component。
// @Description When post this API, the API will return a component info via id
// @ID QueryComponentById
// @tags Configuration
// @Accept  json
// @Produce  json
// @Param id path int64 true "component id"
// @Success 0 {object} configuration.RespComponentId
// @Header 200 {header} string
// @Failure 4005 {object} configuration.RespComponentId "The micro-service can't be reachable"
// @Failure 4001 {object} configuration.RespComponentId "Database problem"
// @Router /config/query-component/{id} [get]
func QueryComponentById(ctx *gin.Context) {
	var (
		statusCode int
		resp       config.RespComponentId
		pbReq      pbConfig.Component
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()
	componentId := ctx.Param("id")
	pbReq.Id, _ = strconv.ParseInt(componentId, 10, 64)
	result, err := rpcConfigService.GetComById(context.TODO(), &pbReq)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if result.StatusCode == utils.RECODE_DBERR {
		resp.SetError(result.StatusCode, "The database has some problems", err)
		statusCode = http.StatusBadRequest
		return
	}
	statusCode = http.StatusOK
	marshal, err := json.Marshal(&result)
	if err != nil {
		return
	}
	err = json.Unmarshal(marshal, &resp.Data)
	if err != nil {
		return
	}
	resp.NewSuccess()
}

// QuerySectionById godoc
// @Summary Get section。
// @Description When post this API, the API will return a section info via id
// @ID QuerySectionById
// @tags Configuration
// @Accept  json
// @Produce  json
// @Param id path string true "section id"
// @Success 0 {object} configuration.RespSectionId
// @Header 200 {header} string
// @Failure 4005 {object} configuration.RespSectionId "The micro-service can't be reachable"
// @Failure 4001 {object} configuration.RespSectionId "Database problem"
// @Router /config/query-section/{id} [get]
func QuerySectionById(ctx *gin.Context) {
	var (
		statusCode int
		resp       config.RespSectionId
		pbReq      pbConfig.Section
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()
	sectionId := ctx.Param("id")
	pbReq.Id, _ = strconv.ParseInt(sectionId, 10, 64)
	result, err := rpcConfigService.QuerySecById(context.TODO(), &pbReq)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if result.StatusCode == utils.RECODE_DBERR {
		resp.SetError(result.StatusCode, "The database has some problems", err)
		statusCode = http.StatusBadRequest
		return
	}

	statusCode = http.StatusOK
	marshal, err := json.Marshal(&result)
	if err != nil {
		return
	}
	err = json.Unmarshal(marshal, &resp.Data)
	if err != nil {
		return
	}
	resp.NewSuccess()
}
