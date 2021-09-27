package handler

import (
	"FriendlyAlmond_backend/pkg/logger"
	"FriendlyAlmond_backend/pkg/model/rdb"
	"time"
)

const CAPTCHA = "captcha:"

type RedisStore struct {
}

//set a capt
func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	err := rdb.SetM(key, value, time.Minute*2)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

//get a capt
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := rdb.GetM(key)
	if err != nil {
		logger.Error(err)
		return ""
	}
	if clear {
		err := rdb.Del(key)
		if err != nil {
			logger.Error(err)
			return ""
		}
	}
	return val
}

//verify a capt
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	if id == "" || answer == "" {
		return false
	}
	v := RedisStore{}.Get(id, clear)
	return v == answer
}
