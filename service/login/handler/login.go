package handler

import (
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/model/mysql"
	"FriendlyAlmond_backend/pkg/model/object/login"
	"FriendlyAlmond_backend/pkg/utils"
	"FriendlyAlmond_backend/proto/login"
	"context"
)

type Login struct{}

//GenerateCaptcha create a captcha, it used for user login and sign in
func (l *Login) GenerateCaptcha(ctx context.Context, req *pbLogin.Empty, resp *pbLogin.Captcha) error {
	defer func() {
		logger.Infof("calling GenerateCaptcha success, resp.CaptchaId=%+v", resp.Id)
	}()
	captchaId, captchaImage, err := CaptMake()
	if err != nil {
		logger.Error(err.Error())
		resp.StatusCode = utils.RECODE_CAPTCHA_GENERATEERR
		return nil
	}
	resp.StatusCode = utils.RECODE_OK
	resp.Id = captchaId
	resp.Image = captchaImage
	return nil
}

//Register receive the user information when user sign in a new account
func (l *Login) Register(ctx context.Context, req *pbLogin.UserInfo, resp *pbLogin.OperationResult) error {
	defer func() {
		logger.Infof("calling Register success, req=%+v resp=%+v", req, resp)
	}()
	if CaptVerify(req.Captcha.Id, req.Captcha.Answer) {
		userInfo := new(login.UserInfo)
		userInfo.Uid = req.Uid
		userInfo.FirstName = req.FirstName
		userInfo.MiddleName = req.MiddleName
		userInfo.LastName = req.LastName
		userInfo.Email = req.Email
		userInfo.Password = req.Password
		userInfo.Phone = req.Phone
		userInfo.Address = req.Address
		userInfo.AreaCode = req.AreaCode
		mysql.UserDB.Create(&userInfo)
		resp.StatusCode = utils.RECODE_OK
	} else {
		resp.StatusCode = utils.RECODE_CAPTCHA_VERIFYERR
	}
	return nil
}

//Update when user update their info like phone number or email
func (l *Login) Update(ctx context.Context, req *pbLogin.UserInfo, resp *pbLogin.OperationResult) error {
	defer func() {
		logger.Infof("calling Update success, req=%+v resp=%+v", req, resp)
	}()
	var (
		userInfo login.UserInfo
	)
	userInfo.Email = req.Email
	userInfo.Address = req.Address
	userInfo.FirstName = req.FirstName
	userInfo.MiddleName = req.MiddleName
	userInfo.LastName = req.LastName
	userInfo.Phone = req.Phone
	userInfo.Password = req.Password
	result := mysql.UserDB.Model(&login.UserInfo{}).Where("uid = ?", req.Uid).Updates(&userInfo)
	if result.Error != nil {
		resp.StatusCode = utils.RECODE_DBERR
		return nil
	}
	resp.StatusCode = utils.RECODE_OK
	return nil

}

//Query when user get their info from the database, the function will return it
func (l *Login) Query(ctx context.Context, req *pbLogin.UserInfo, resp *pbLogin.UserInfo) error {
	defer func() {
		logger.Infof("calling Query success, req=%+v resp=%+v", req, resp)
	}()
	var (
		userInfo        login.UserInfo
		OperationResult pbLogin.OperationResult
	)
	result := mysql.UserDB.Where("uid = ?", req.Uid).Find(&userInfo)
	if result.Error != nil {
		OperationResult.StatusCode = utils.RECODE_DBERR
		return nil
	}
	resp.Uid = userInfo.Uid
	userInfo.FirstName = req.FirstName
	userInfo.MiddleName = req.MiddleName
	userInfo.LastName = req.LastName
	resp.Email = userInfo.Email
	resp.Password = userInfo.Password
	resp.Phone = userInfo.Phone
	resp.Address = userInfo.Address
	OperationResult.StatusCode = utils.RECODE_OK
	resp.OperationResult = &OperationResult
	return nil
}

//Login when user login the system, the first step need user pass the captcha test, then the func will verify the email and password, finally, return the statusCode
func (l *Login) Login(ctx context.Context, req *pbLogin.UserInfo, resp *pbLogin.UserInfo) error {
	defer func() {
		logger.Infof("calling Login success,  req=%+v resp=%+v", req, resp)
	}()
	op := new(pbLogin.OperationResult)
	if CaptVerify(req.Captcha.Id, req.Captcha.Answer) {
		userInfo := new(login.UserInfo)
		mysql.UserDB.Where(&login.UserInfo{Uid: req.Uid, Email: req.Email, Password: req.Password}).Find(&userInfo)
		if req.Email == userInfo.Email && req.Password == userInfo.Password {
			op.StatusCode = utils.RECODE_OK
			resp.FirstName = userInfo.FirstName
			resp.Password = userInfo.Password
			resp.Email = userInfo.Email
			resp.Uid = userInfo.Uid
			resp.Phone = userInfo.Phone
			resp.Address = userInfo.Address
			resp.MiddleName = userInfo.MiddleName
			resp.LastName = userInfo.LastName
			resp.AreaCode = userInfo.AreaCode
			resp.OperationResult = op
			return nil
		} else if req.Email != userInfo.Email || req.Password != userInfo.Password {
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
