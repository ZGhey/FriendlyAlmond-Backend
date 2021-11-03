package handler

import (
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/model/mysql"
	config "FriendlyAlmond_backend/pkg/model/object/configuration"
	"FriendlyAlmond_backend/pkg/utils"
	pbConfig "FriendlyAlmond_backend/proto/configuration"
	"context"
	"encoding/json"
	"strconv"
)

type Config struct{}

func (c Config) QuerySecById(ctx context.Context, req *pbConfig.Section, resp *pbConfig.Section) error {
	var (
		sections config.Section
	)
	defer func() {
		logger.Infof("calling QuerySecById success, req=%+v, resp=%+v", req, resp)
	}()
	if sectionResult := mysql.ConfigBoatDB.First(&sections, "id = ?", req.Id); sectionResult.Error != nil {
		logger.Error(sectionResult.Error)
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}
	marshal, err := json.Marshal(&sections)
	if err != nil {
		return err
	}
	err = json.Unmarshal(marshal, resp)
	if err != nil {
		return err
	}
	resp.StatusCode = utils.RECODE_OK
	return nil

}

func (c Config) GetComById(ctx context.Context, req *pbConfig.Component, resp *pbConfig.Component) error {
	var (
		component config.Component
	)
	defer func() {
		logger.Infof("calling QueryComById success, req=%+v, resp=%+v", req, resp)
	}()
	if componentResult := mysql.ConfigBoatDB.First(&component, "id = ?", req.Id); componentResult.Error != nil {
		logger.Error(componentResult.Error)
		resp.StatusCode = utils.RECODE_DBERR

		return nil
	}
	marshal, err := json.Marshal(&component)
	if err != nil {
		return err
	}
	err = json.Unmarshal(marshal, resp)
	if err != nil {
		return err
	}
	resp.StatusCode = utils.RECODE_OK
	return nil
}

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

	if categoryResult := mysql.ConfigBoatDB.Where("type = ?", req.Type).Find(&categories); categoryResult.Error != nil {
		logger.Error(categoryResult.Error)
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}

	for _, value := range categories {
		categoryId = append(categoryId, strconv.FormatInt(value.Id, 10))
	}

	if componentResult := mysql.ConfigBoatDB.Where("category_id IN ?", categoryId).Find(&components); componentResult.Error != nil {
		logger.Error(componentResult.Error)
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}

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
