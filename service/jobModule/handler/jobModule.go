package handler

import (
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/model/mysql"
	config "FriendlyAlmond_backend/pkg/model/object/configuration"
	"FriendlyAlmond_backend/pkg/model/object/jobModule"
	"FriendlyAlmond_backend/pkg/model/object/login"
	"FriendlyAlmond_backend/pkg/model/object/order"
	"FriendlyAlmond_backend/pkg/utils"
	pbJobModule "FriendlyAlmond_backend/proto/jobModule"
	"context"
)

type JobModule struct{}

func (j JobModule) QueryMostPopular(ctx context.Context, req *pbJobModule.Empty, resp *pbJobModule.MostPopular) error {
	var (
		orderData      order.Order
		orderSection   order.Section
		orderComponent order.Component
	)
	defer func() {
		logger.Infof("calling QueryMostPopular success, req=%+v, resp=%+v", req, resp)
	}()
	if result := mysql.OrderDB.Model(&orderData).Select("color, count(color) as total").Group("color").Order("total desc").Find(&resp.Colors); result.Error != nil {
		logger.Error("the data is not exists in the database")
		logger.Error(result.Error)
		resp.StatusCode = utils.RECODE_DATAINEXISTENCE
		return nil
	}
	if result := mysql.OrderDB.Model(&orderSection).Select("section_id as section,count(section_id) as total").Group("section_id").Order("total desc").Find(&resp.Sections); result.Error != nil {
		logger.Error("the data is not exists in the database")
		logger.Error(result.Error)
		resp.StatusCode = utils.RECODE_DATAINEXISTENCE
		return nil
	}
	for index, value := range resp.Sections {
		section := new(config.Section)
		if result := mysql.ConfigBoatDB.Model(&section).Where("id = ?", value.Section).Find(&section); result.Error != nil {
			logger.Error("the data is not exists in the database")
			logger.Error(result.Error)
			resp.StatusCode = utils.RECODE_DATAINEXISTENCE
			return nil
		}
		resp.Sections[index].Section = section.Name
	}

	if result := mysql.OrderDB.Model(&orderComponent).Select("component_id as component,count(component_id) as total").Group("component_id").Order("total desc").Find(&resp.Components); result.Error != nil {
		logger.Error("the data is not exists in the database")
		logger.Error(result.Error)
		resp.StatusCode = utils.RECODE_DATAINEXISTENCE
		return nil
	}
	for index, value := range resp.Components {
		component := new(config.Component)
		if result := mysql.ConfigBoatDB.Model(&component).Where("id = ?", value.Component).Find(&component); result.Error != nil {
			logger.Error("the data is not exists in the database")
			logger.Error(result.Error)
			resp.StatusCode = utils.RECODE_DATAINEXISTENCE
			return nil
		}
		resp.Components[index].Component = component.Details
	}
	resp.StatusCode = utils.RECODE_OK
	return nil
}

func (j JobModule) QueryTotalSales(ctx context.Context, req *pbJobModule.Empty, resp *pbJobModule.TotalSales) error {
	var (
		orderData order.Order
	)
	defer func() {
		logger.Infof("calling QueryTotalSales success, req=%+v, resp=%+v", req, resp)
	}()
	//calculate total sales for last 1 month
	if result := mysql.OrderDB.Model(&orderData).Select("sum(total_price) as LastOneMonth").Where("PERIOD_DIFF(date_format(now(), '%Y%m'), date_format(created, '%Y%m')) = 1").Find(&resp.LastOneMonth); result.Error != nil {
		logger.Error(result.Error)
		resp.StatusCode = utils.RECODE_DATAINEXISTENCE
		return nil
	}

	//calculate total sales for last 3 month
	if result := mysql.OrderDB.Model(&orderData).Select("sum(total_price) as LastThreeMonth").Where("PERIOD_DIFF(date_format(now(), '%Y%m'), date_format(created, '%Y%m')) = 3").Find(&resp.LastThreeMonth); result.Error != nil {
		logger.Error(result.Error)
	}

	//calculate total sales for last 6 month
	if result := mysql.OrderDB.Model(&orderData).Select("sum(total_price) as LastSixMonth").Where("PERIOD_DIFF(date_format(now(), '%Y%m'), date_format(created, '%Y%m')) = 6").Find(&resp.LastSixMonth); result.Error != nil {
		logger.Error(result.Error)
	}

	//calculate total sales for last year
	if result := mysql.OrderDB.Model(&orderData).Select("sum(total_price) as LastOneYear").Where("year(created) = year(date_sub(now(), interval 1 year))").Find(&resp.LastOneYear); result.Error != nil {
		logger.Error(result.Error)
	}
	resp.StatusCode = utils.RECODE_OK
	return nil
}

func (j JobModule) QueryNoJobOrder(ctx context.Context, req *pbJobModule.Empty, resp *pbJobModule.ListQueryOrder) error {
	var (
		orderData      []order.Order
		orderSection   []order.Section
		orderComponent []order.Component
		boatmodel      config.Boat
		category       config.Category
		userInfo       login.UserInfo
	)
	defer func() {
		logger.Infof("calling QueryOrder success, req=%+v, resp=%+v", req, resp)
	}()

	if result := mysql.OrderDB.Where("job_id = ?", 0).Order("created desc").Find(&orderData); result.Error != nil {
		logger.Error("the data is not exists in the database")
		logger.Error(result.Error)
		resp.StatusCode = utils.RECODE_DATAINEXISTENCE
		return nil
	}

	for _, value := range orderData {
		pbQueryOrder := new(pbJobModule.QueryOrder)
		pbQueryOrder.Color = value.Color
		pbQueryOrder.BoatName = value.BoatName
		pbQueryOrder.OrderDate = value.Created.String()
		pbQueryOrder.TotalPrice = value.TotalPrice
		pbQueryOrder.OrderId = value.OrderId

		if result := mysql.UserDB.Where("uid = ?", value.Uid).Find(&userInfo); result.Error != nil {
			logger.Error(result.Error)
			resp.StatusCode = utils.RECODE_DATAINEXISTENCE
			return nil
		}

		if userInfo.MiddleName != "" {
			pbQueryOrder.UserName = userInfo.FirstName + " " + userInfo.MiddleName + " " + userInfo.LastName
		} else {
			pbQueryOrder.UserName = userInfo.FirstName + " " + userInfo.LastName
		}

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
			pbQueryOrder.ComponentId = append(pbQueryOrder.ComponentId, component.Id)
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
			pbQueryOrder.SectionId = append(pbQueryOrder.SectionId, section.Id)
		}

		resp.QueryOrder = append(resp.QueryOrder, pbQueryOrder)
	}

	resp.StatusCode = utils.RECODE_OK
	return nil
}

func (j JobModule) CreateJob(ctx context.Context, req *pbJobModule.Job, resp *pbJobModule.OperationResult) error {
	defer func() {
		logger.Infof("calling CreateJob success, req=%+v resp=%+v", req, resp)
	}()
	var (
		job    jobModule.Job
		order1 order.Order
	)
	job.OrderId = req.OrderId
	result := mysql.JobModule.Create(&job)
	if result.Error != nil {
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}
	result = mysql.OrderDB.Model(&order1).Where("order_id = ?", req.OrderId).Update("job_id", job.JobId)
	if result.Error != nil {
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}

	resp.StatusCode = utils.RECODE_OK
	return nil
}

func (j JobModule) CreateTask(ctx context.Context, req *pbJobModule.Task, resp *pbJobModule.OperationResult) error {
	defer func() {
		logger.Infof("calling CreateTask success, req=%+v resp=%+v", req, resp)
	}()
	var (
		task jobModule.Task
	)
	task.JobId = req.JobId
	task.SectionId = req.SectionId
	task.ComponentId = req.ComponentId
	task.Description = req.Description
	task.StartDate = req.StartDate
	task.DueDate = req.DueDate
	task.StaffId = req.StaffId
	result := mysql.JobModule.Create(&task)
	if result.Error != nil {
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}
	resp.StatusCode = utils.RECODE_OK
	return nil
}

func (j JobModule) QueryTask(ctx context.Context, req *pbJobModule.Task, resp *pbJobModule.ListTask) error {
	var (
		task []jobModule.Task
	)
	defer func() {
		logger.Infof("calling QueryTask success, req=%+v, resp=%+v", req, resp)
	}()
	if result := mysql.JobModule.Find(&task, "job_id", req.JobId); result.Error != nil {
		logger.Error("the data is not exists in the database")
		logger.Error(result.Error)
		resp.StatusCode = utils.RECODE_DATAINEXISTENCE
		return nil
	}

	for _, value := range task {
		pbTask := new(pbJobModule.Task)
		pbTask.JobId = value.JobId
		pbTask.TaskId = value.TaskId
		pbTask.SectionId = value.SectionId
		pbTask.ComponentId = value.ComponentId
		pbTask.Description = value.Description
		pbTask.StartDate = value.StartDate
		pbTask.DueDate = value.DueDate
		resp.Task = append(resp.Task, pbTask)
	}

	resp.StatusCode = utils.RECODE_OK
	return nil
}

func (j JobModule) Login(ctx context.Context, req *pbJobModule.Staff, resp *pbJobModule.Staff) error {
	defer func() {
		logger.Infof("calling Login success,  req=%+v resp=%+v", req, resp)
	}()
	op := new(pbJobModule.OperationResult)
	if CaptVerify(req.Captcha.Id, req.Captcha.Answer) {
		staffInfo := new(jobModule.Staff)

		mysql.JobModule.Where(&jobModule.Staff{Account: req.Email, Password: req.Password}).Find(&staffInfo)
		if req.Account == staffInfo.Account && req.Password == staffInfo.Password {
			op.StatusCode = utils.RECODE_OK
			resp.Address = staffInfo.Address
			resp.AreaCode = staffInfo.AreaCode
			resp.Phone = staffInfo.Phone
			resp.Email = staffInfo.Email
			resp.StaffId = staffInfo.StaffId
			resp.Password = staffInfo.Password
			resp.Firstname = staffInfo.Firstname
			resp.Lastname = staffInfo.Lastname
			resp.Middlename = staffInfo.Middlename
			resp.Skill = staffInfo.Skill
			resp.Account = staffInfo.Account
			resp.OperationResult = op
			return nil
		} else if req.Account != staffInfo.Account || req.Password != staffInfo.Password {
			op.StatusCode = utils.RECODE_LOGINERR
			resp.OperationResult = op
			return nil
		}
	} else {
		op.StatusCode = utils.RECODE_CAPTCHA_VERIFYERR
		resp.OperationResult = op
	}
	return nil
}

func (j JobModule) UpdateStaff(ctx context.Context, req *pbJobModule.Staff, resp *pbJobModule.OperationResult) error {
	defer func() {
		logger.Infof("calling UpdateStaff success, req=%+v resp=%+v", req, resp)
	}()
	var (
		staffInfo jobModule.Staff
	)
	staffInfo.Email = req.Email
	staffInfo.Address = req.Address
	staffInfo.Firstname = req.Firstname
	staffInfo.Middlename = req.Middlename
	staffInfo.Lastname = req.Lastname
	staffInfo.Phone = req.Phone
	staffInfo.Password = req.Password
	staffInfo.Account = req.Account
	staffInfo.Skill = req.Skill
	staffInfo.AreaCode = req.AreaCode
	result := mysql.JobModule.Model(&jobModule.Staff{}).Where("account = ?", req.Account).Updates(&staffInfo)
	if result.Error != nil {
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}
	resp.StatusCode = utils.RECODE_OK
	return nil
}

func (j JobModule) AddStaff(ctx context.Context, req *pbJobModule.Staff, resp *pbJobModule.OperationResult) error {
	defer func() {
		logger.Infof("calling AddStaff success, req=%+v resp=%+v", req, resp)
	}()
	var (
		staffInfo jobModule.Staff
	)
	staffInfo.Email = req.Email
	staffInfo.Address = req.Address
	staffInfo.Firstname = req.Firstname
	staffInfo.Middlename = req.Middlename
	staffInfo.Lastname = req.Lastname
	staffInfo.Phone = req.Phone
	staffInfo.Password = req.Password
	staffInfo.Account = req.Account
	staffInfo.Skill = req.Skill
	staffInfo.AreaCode = req.AreaCode
	staffInfo.StaffId = req.StaffId
	result := mysql.JobModule.Create(&staffInfo)
	if result.Error != nil {
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}
	resp.StatusCode = utils.RECODE_OK
	return nil
}

func (j JobModule) QueryListStaff(ctx context.Context, req *pbJobModule.Empty, resp *pbJobModule.ListStaff) error {
	var (
		staff []jobModule.Staff
	)
	defer func() {
		logger.Infof("calling QueryListStaff success, req=%+v, resp=%+v", req, resp)
	}()
	if result := mysql.JobModule.Order("created desc").Find(&staff); result.Error != nil {
		logger.Error("the data is not exists in the database")
		logger.Error(result.Error)
		resp.StatusCode = utils.RECODE_DATAINEXISTENCE
		return nil
	}

	for _, value := range staff {
		pbStaff := new(pbJobModule.Staff)
		pbStaff.Account = value.Account
		pbStaff.Password = value.Password
		pbStaff.Firstname = value.Firstname
		pbStaff.Lastname = value.Lastname
		pbStaff.Middlename = value.Middlename
		pbStaff.Email = value.Email
		pbStaff.Address = value.Address
		pbStaff.Phone = value.Phone
		pbStaff.Skill = value.Skill
		pbStaff.AreaCode = value.AreaCode
		pbStaff.StaffId = value.StaffId
		resp.Staff = append(resp.Staff, pbStaff)
	}

	resp.StatusCode = utils.RECODE_OK
	return nil
}

func (j JobModule) QueryListUser(ctx context.Context, req *pbJobModule.Empty, resp *pbJobModule.ListUser) error {
	var (
		userInfo []login.UserInfo
	)
	defer func() {
		logger.Infof("calling QueryListUser success, req=%+v, resp=%+v", req, resp)
	}()
	if result := mysql.UserDB.Order("created desc").Find(&userInfo); result.Error != nil {
		logger.Error("the data is not exists in the database")
		logger.Error(result.Error)
		resp.StatusCode = utils.RECODE_DATAINEXISTENCE
		return nil
	}

	for _, value := range userInfo {
		pbUserInfo := new(pbJobModule.UserInfo)
		pbUserInfo.Uid = value.Uid
		pbUserInfo.FirstName = value.FirstName
		pbUserInfo.MiddleName = value.MiddleName
		pbUserInfo.LastName = value.LastName
		pbUserInfo.Email = value.Email
		pbUserInfo.Password = value.Password
		pbUserInfo.Phone = value.Phone
		pbUserInfo.Address = value.Address
		pbUserInfo.AreaCode = value.AreaCode
		resp.UserInfo = append(resp.UserInfo, pbUserInfo)
	}

	resp.StatusCode = utils.RECODE_OK
	return nil
}

func String(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
