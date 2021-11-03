package utils

const (
	RECODE_OK               = "0"
	RECODE_DBERR            = "4001"
	RECODE_NODATA           = "4002"
	RECODE_DATAERR          = "4004"
	RECODE_MICROERR         = "4005"
	RECODE_STOREDATA_FAILED = "4006"
	RECODE_DATAINEXISTENCE  = "4007"
	RECODE_STOREDATA_OK     = "4008"

	RECODE_TOKENERR = "4101"
	RECODE_LOGINERR = "4102"

	RECODE_TOKENEXPIRED = "4115"

	//verify image
	RECODE_CAPTCHA_GENERATEERR = "4125"
	RECODE_CAPTCHA_VERIFYERR   = "4126"

	//4500
	RECODE_UNKNOWERR = "4501"
)

var recodeText = map[string]string{
	//access success
	RECODE_OK: "success",

	//database operation
	RECODE_DBERR:            "database problem",
	RECODE_DATAERR:          "the data in database have problem",
	RECODE_MICROERR:         "the micro service problem",
	RECODE_STOREDATA_FAILED: "store data into database is failed",
	RECODE_DATAINEXISTENCE:  "the data is not exists in the database",
	RECODE_STOREDATA_OK:     "the data store successful",

	//token and login
	RECODE_TOKENERR:     "verify the token failed",
	RECODE_LOGINERR:     "the account or password is not correct",
	RECODE_TOKENEXPIRED: "the token is expired",

	//captcha
	RECODE_CAPTCHA_GENERATEERR: "generate the captcha failed",
	RECODE_CAPTCHA_VERIFYERR:   "the answer of captcha is not correct",
}

func RecodeTest(code string) string {
	str, ok := recodeText[code]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWERR]
}
