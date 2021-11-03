package handler

import (
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/model/mysql"
	config "FriendlyAlmond_backend/pkg/model/object/configuration"
	"FriendlyAlmond_backend/pkg/model/object/order"
	"FriendlyAlmond_backend/pkg/utils"
	pbOrder "FriendlyAlmond_backend/proto/order"
	"context"
)

type Order struct {
}

//QueryOrder query order by uid, it will respond a list of order
func (o Order) QueryOrder(ctx context.Context, req *pbOrder.OrderInfo, resp *pbOrder.ListQueryOrder) error {
	var (
		orderData      []order.Order
		orderSection   []order.Section
		orderComponent []order.Component
		boatmodel      config.Boat
		category       config.Category
	)
	defer func() {
		logger.Infof("calling QueryOrder success, req=%+v, resp=%+v", req, resp)
	}()

	if req.Uid == "" {
		if result := mysql.OrderDB.Order("created desc").Find(&orderData); result.Error != nil {
			logger.Error("the data is not exists in the database")
			logger.Error(result.Error)
			resp.StatusCode = utils.RECODE_DATAINEXISTENCE
			return nil
		}
	} else {
		if result := mysql.OrderDB.Where("uid = ?", req.Uid).Order("created desc").Find(&orderData); result.Error != nil {
			logger.Error("the data is not exists in the database" + req.BoatmodelName)
			logger.Error(result.Error)
			resp.StatusCode = utils.RECODE_DATAINEXISTENCE
			return nil
		}
	}

	for _, value := range orderData {
		pbQueryOrder := new(pbOrder.QueryOrder)
		pbQueryOrder.Color = value.Color
		pbQueryOrder.BoatName = value.BoatName
		pbQueryOrder.Uid = value.Uid
		pbQueryOrder.TotalPrice = value.TotalPrice
		pbQueryOrder.OrderId = value.OrderId
		pbQueryOrder.JobId = value.JobId
		if result := mysql.ConfigBoatDB.Where("id = ?", value.BoatmodelId).Find(&boatmodel); result.Error != nil {
			logger.Error(result.Error)
			resp.StatusCode = utils.RECODE_DATAINEXISTENCE
			return nil
		}
		pbQueryOrder.BoatmodelName = boatmodel.Name

		if result := mysql.ConfigBoatDB.Where("id = ?", value.CategoryId).Find(&category); result.Error != nil {
			logger.Error(result.Error)
			resp.StatusCode = utils.RECODE_DATAINEXISTENCE
			return nil
		}
		pbQueryOrder.CategoryName = category.Name

		if result := mysql.OrderDB.Where("order_id = ?", value.OrderId).Find(&orderComponent); result.Error != nil {
			logger.Error(result.Error)
			resp.StatusCode = utils.RECODE_DATAINEXISTENCE
			return nil
		}

		for _, comValue := range orderComponent {
			component := new(config.Component)
			if result := mysql.ConfigBoatDB.Where("id = ?", comValue.ComponentId).Find(&component); result.Error != nil {
				logger.Error(result.Error)
				resp.StatusCode = utils.RECODE_DATAINEXISTENCE
				return nil
			}
			pbQueryOrder.Options = append(pbQueryOrder.Options, component.Details)
		}

		if result := mysql.OrderDB.Where("order_id = ?", value.OrderId).Find(&orderSection); result.Error != nil {
			logger.Error(result.Error)
			resp.StatusCode = utils.RECODE_DATAINEXISTENCE
			return nil
		}

		for _, secValue := range orderSection {
			section := new(config.Section)
			if result := mysql.ConfigBoatDB.Where("id = ?", secValue.SectionId).Find(&section); result.Error != nil {
				logger.Error(result.Error)
				resp.StatusCode = utils.RECODE_DATAINEXISTENCE
				return nil
			}
			pbQueryOrder.Options = append(pbQueryOrder.Options, section.Name)
		}

		resp.QueryOrder = append(resp.QueryOrder, pbQueryOrder)
	}

	resp.StatusCode = utils.RECODE_OK
	return nil
}

//CreateOrder based on the order info like component or section, this method will create an order to the database
func (o Order) CreateOrder(ctx context.Context, req *pbOrder.OrderInfo, resp *pbOrder.OperateResult) error {
	var (
		orderData      order.Order
		orderSection   order.Section
		orderComponent order.Component
		boatmodel      config.Boat
		category       config.Category
	)
	defer func() {
		logger.Infof("calling CreateOrder success, req=%+v, resp=%+v", req, resp)
	}()
	if result := mysql.ConfigBoatDB.Where("name = ?", req.BoatmodelName).Find(&boatmodel); result.Error != nil {
		logger.Error("the data is not exists in the database" + req.BoatmodelName)
		logger.Error(result.Error)
		resp.StatusCode = utils.RECODE_DATAINEXISTENCE
		return nil
	}

	if result := mysql.ConfigBoatDB.Where("type = ?", req.CategoryName).Find(&category); result.Error != nil {
		logger.Error("the data is not exists in the database" + req.CategoryName)
		logger.Error(result.Error)
		resp.StatusCode = utils.RECODE_DATAINEXISTENCE
		return nil
	}
	orderData.BoatmodelId = boatmodel.Id
	orderData.CategoryId = category.Id
	orderData.Uid = req.Uid
	orderData.BoatName = req.BoatName
	orderData.Color = req.Color
	orderData.TotalPrice = req.TotalPrice
	if result := mysql.OrderDB.Create(&orderData); result.Error != nil {
		logger.Error(result.Error)
		resp.StatusCode = utils.RECODE_STOREDATA_FAILED
		return nil
	}

	for _, value := range req.ComponentId {
		orderComponent.OrderId = orderData.OrderId
		orderComponent.ComponentId = value
		if result := mysql.OrderDB.Create(&orderComponent); result.Error != nil {
			logger.Error(result.Error)
			resp.StatusCode = utils.RECODE_STOREDATA_FAILED
			return nil
		}
	}

	for _, value := range req.SectionId {
		orderSection.OrderId = orderData.OrderId
		orderSection.SectionId = value
		if result := mysql.OrderDB.Create(&orderSection); result.Error != nil {
			logger.Error(result.Error)
			resp.StatusCode = utils.RECODE_STOREDATA_FAILED
			return nil
		}
	}

	resp.StatusCode = utils.RECODE_OK
	return nil
}
