package handler

import (
	"FriendlyAlmond_backend/pkg/logger"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// use redis to store the id and answer
var store = RedisStore{}

//CaptMake generate captcha
func CaptMake() (id, b64s string, err error) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	// config the info of captcha
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 5 | 10,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	lid, lb64s, lerr := captcha.Generate()
	return lid, lb64s, lerr
}

//verify a captcha
func CaptVerify(id string, capt string) bool {
	logger.Infof("id=%+v, capt=%+v", id, capt)
	if store.Verify(id, capt, true) {
		return true
	} else {
		return false
	}
}
