package controller

import (
	"FriendlyAlmond_backend/pkg/model"
	"FriendlyAlmond_backend/pkg/model/object/order"
	"FriendlyAlmond_backend/pkg/utils"
	pbOrder "FriendlyAlmond_backend/proto/order"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateOrder godoc
// @Summary create order data。
// @Description When post this API, the API will return the result of creating order, and insert the order data into database
// @ID CreateOrder
// @tags Order
// @Accept  json
// @Produce  json
// @Param type body order.ReqOrder true "all the information of the order."
// @Success 0 {object} model.JSONResult
// @Header 200 {header} string
// @Failure 4005 {object} model.JSONResult "The micro-service can't be reachable"
// @Failure 4001 {object} model.JSONResult "Database problem"
// @Router /order/create [post]
func CreateOrder(ctx *gin.Context) {
	var (
		statusCode  int
		req         order.ReqOrder
		resp        model.JSONResult
		pbOrderInfo pbOrder.OrderInfo
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.SetError(utils.RECODE_DATAERR, utils.RecodeTest(utils.RECODE_DATAERR), err)
		statusCode = http.StatusBadRequest
	}
	pbOrderInfo.BoatmodelName = req.BoatmodelName
	pbOrderInfo.CategoryName = req.CategoryName
	pbOrderInfo.BoatName = req.BoatName
	pbOrderInfo.Uid = req.Uid
	pbOrderInfo.Color = req.Color
	pbOrderInfo.TotalPrice = req.TotalPrice
	pbOrderInfo.SectionId = req.SectionId
	pbOrderInfo.ComponentId = req.ComponentId
	result, err := rpcOrderService.CreateOrder(context.TODO(), &pbOrderInfo)
	if err != nil {
		resp.SetError(utils.RECODE_MICROERR, utils.RecodeTest(utils.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if result.StatusCode != utils.RECODE_OK {
		resp.SetError(result.StatusCode, result.StatusCode, err)
		statusCode = http.StatusBadRequest
		return
	}

	statusCode = http.StatusOK
	resp.NewSuccess()
}

// QueryOrder godoc
// @Summary query order data。
// @Description When post this API, the API will return the order details via uid
// @ID QueryOrder
// @tags Order
// @Accept  json
// @Produce  json
// @Param type body order.ReqOrder true "just need the uid."
// @Success 0 {object} order.ApiOrder
// @Header 200 {header} string
// @Failure 4005 {object} order.ApiOrder "The micro-service can't be reachable"
// @Failure 4001 {object} order.ApiOrder "Database problem"
// @Router /order/query [post]
func QueryOrder(ctx *gin.Context) {
	var (
		statusCode  int
		req         order.ReqOrder
		resp        order.ApiOrder
		pbOrderInfo pbOrder.OrderInfo
	)
	defer func() {
		responseHTTP(ctx, statusCode, &resp)
	}()

	if err := ctx.ShouldBindJSON(&req); err != nil {
		resp.SetError(utils.RECODE_DATAERR, utils.RecodeTest(utils.RECODE_DATAERR), err)
		statusCode = http.StatusBadRequest
	}
	if req.Uid != "" {
		pbOrderInfo.Uid = req.Uid
	}
	result, err := rpcOrderService.QueryOrder(context.TODO(), &pbOrderInfo)
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
		respOrder.BoatmodelName = value.BoatmodelName
		respOrder.Options = value.Options
		respOrder.Uid = value.Uid
		respOrder.Color = value.Color
		respOrder.BoatName = value.BoatName
		respOrder.TotalPrice = value.TotalPrice
		respOrder.CategoryName = value.CategoryName
		resp.Data = append(resp.Data, respOrder)
	}

	statusCode = http.StatusOK
	resp.NewSuccess()
}
