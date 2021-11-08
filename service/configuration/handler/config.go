package handler

import (
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/model/mysql"
	config "FriendlyAlmond_backend/pkg/model/object/configuration"
	"FriendlyAlmond_backend/pkg/utils"
	pbConfig "FriendlyAlmond_backend/proto/configuration"
	"context"
	"strconv"
)

type Config struct{}

//QuerySecById required a list of section id, response the details of the current section
func (c Config) QuerySecById(ctx context.Context, req *pbConfig.ListId, resp *pbConfig.ListSection) error {
	var (
		sections []config.Section
	)
	defer func() {
		logger.Infof("calling QuerySecById success, req=%+v, resp=%+v", req, resp)
	}()
	//select * from section where id in req.id
	if sectionResult := mysql.ConfigBoatDB.Where("id in ?", req.Id).Find(&sections); sectionResult.Error != nil {
		logger.Error(sectionResult.Error)
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}
	//Looping adds each section object to pbConfig.ListSection
	for _, value := range sections {
		pbSec := new(pbConfig.Section)
		pbSec.Id = value.Id
		pbSec.CategoryId = value.Id
		//pbSec.Name = value.Name
		pbSec.Specs = value.Specs
		pbSec.Price = value.Price
		pbSec.Detail = value.Detail
		resp.Data = append(resp.Data, pbSec)
	}
	resp.StatusCode = utils.RECODE_OK
	return nil
}

//QueryComById has the same functional with QuerySecById, but different is focus on component table in mysql
func (c Config) QueryComById(ctx context.Context, req *pbConfig.ListId, resp *pbConfig.ListComponent) error {
	var (
		components []config.Component
	)
	defer func() {
		logger.Infof("calling QueryComById success, req=%+v, resp=%+v", req, resp)
	}()
	if componentResult := mysql.ConfigBoatDB.Where("id in ?", req.Id).Find(&components); componentResult.Error != nil {
		logger.Error(componentResult.Error)
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}
	for _, value := range components {
		pbCom := new(pbConfig.Component)
		pbCom.Id = value.Id
		pbCom.Name = value.Name
		pbCom.Details = value.Details
		pbCom.SupplierId = value.SupplierId
		pbCom.CategoryId = value.CategoryId
		pbCom.Price = value.Price
		resp.Data = append(resp.Data, pbCom)
	}
	resp.StatusCode = utils.RECODE_OK
	return nil
}

//QueryBoat return the boat module, like lightship
func (c Config) QueryBoat(ctx context.Context, req *pbConfig.Empty, resp *pbConfig.ListBoat) error {
	defer func() {
		logger.Infof("calling QueryBoat success, resp=%+v", resp)
	}()
	var (
		boat   []config.Boat
		pbBoat pbConfig.Boat
	)

	if result := mysql.ConfigBoatDB.Find(&boat); result.Error != nil {
		logger.Error(result.Error)
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}
	for _, value := range boat {
		pbBoat.Id = value.Id
		pbBoat.Name = value.Name
		pbBoat.Shape = value.Shape
		pbBoat.Description = value.Description
		resp.Data = append(resp.Data, &pbBoat)
	}
	return nil
}

func (c Config) QueryCategory(ctx context.Context, req *pbConfig.Boat, resp *pbConfig.ListCategory) error {
	panic("implement me")
}

//QueryComponent get the boat component from database via pbConfig.Category, it will require a type like motor
func (c Config) QueryComponent(ctx context.Context, req *pbConfig.Category, resp *pbConfig.ListComponent) error {
	var (
		categories []config.Category
		components []config.Component
		categoryId []string
	)
	defer func() {
		logger.Infof("calling QueryComponent success, req=%+v, resp=%+v", req, resp)
	}()

	//According to the type of the boat, like motor, return all the category information belong to the motor
	if categoryResult := mysql.ConfigBoatDB.Where("type = ?", req.Type).Find(&categories); categoryResult.Error != nil {
		logger.Error(categoryResult.Error)
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}

	//add each category id into categoryId list
	for _, value := range categories {
		categoryId = append(categoryId, strconv.FormatInt(value.Id, 10))
	}

	//based on the category id to find which component is followed, then return the details, binding the module of components
	if componentResult := mysql.ConfigBoatDB.Where("category_id IN ?", categoryId).Find(&components); componentResult.Error != nil {
		logger.Error(componentResult.Error)
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}

	//looping two different object, when both category id are same, add the detail into resp.Data
	for _, comValue := range components {
		for _, cateValue := range categories {
			if comValue.CategoryId == cateValue.Id {
				pbComponent := new(pbConfig.Component)
				pbComponent.Id = comValue.Id
				pbComponent.Details = comValue.Details
				pbComponent.Name = comValue.Name
				pbComponent.CategoryId = comValue.CategoryId
				pbComponent.SupplierId = comValue.SupplierId
				pbComponent.Price = comValue.Price
				pbComponent.CategoryName = cateValue.Name
				resp.Data = append(resp.Data, pbComponent)
			} else {
				continue
			}
		}
	}
	resp.StatusCode = utils.RECODE_OK
	return nil
}

func (c Config) QueryPackage(ctx context.Context, empty *pbConfig.Empty, boat *pbConfig.ListPackage) error {
	panic("implement me")
}

//QuerySection same function with QueryComponent
func (c Config) QuerySection(ctx context.Context, req *pbConfig.Category, resp *pbConfig.ListSection) error {
	var (
		categories []config.Category
		sections   []config.Section
		categoryId []string
	)
	defer func() {
		logger.Infof("calling QuerySection success, req=%+v, resp=%+v", req, resp)
	}()

	if categoryResult := mysql.ConfigBoatDB.Where("type = ?", req.Type).Find(&categories); categoryResult.Error != nil {
		logger.Error(categoryResult.Error)
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}

	for _, value := range categories {
		categoryId = append(categoryId, strconv.FormatInt(value.Id, 10))
	}

	if sectionResult := mysql.ConfigBoatDB.Where("category_id IN ?", categoryId).Find(&sections); sectionResult.Error != nil {
		logger.Error(sectionResult.Error)
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}

	for _, secValue := range sections {
		for _, cateValue := range categories {
			if secValue.CategoryId == cateValue.Id {
				pbSection := new(pbConfig.Section)
				pbSection.Id = secValue.Id
				pbSection.Name = secValue.Name
				pbSection.CategoryId = secValue.CategoryId
				pbSection.Price = secValue.Price
				pbSection.CategoryName = cateValue.Name
				resp.Data = append(resp.Data, pbSection)
			} else {
				continue
			}
		}
	}
	resp.StatusCode = utils.RECODE_OK
	return nil
}

func (c Config) QuerySupplier(ctx context.Context, empty *pbConfig.Empty, boat *pbConfig.ListSupplier) error {
	panic("implement me")
}
