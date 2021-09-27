package rdb

import (
	"FriendlyAlmond_backend/pkg/utils"
	"context"
	"time"

	"github.com/asim/go-micro/v3/logger"
	"github.com/go-redis/redis/v8"
)

func Get(key string) (string, string) {
	result, err := REDISDB.RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		return "", utils.RECODE_DATAINEXISTENCE
	}
	return result, utils.RECODE_OK
}

func SetS(key string, value interface{}, expiration time.Duration) string {
	err := REDISDB.RedisClient.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return utils.RECODE_STOREDATA_FAILED
	}
	return utils.RECODE_STOREDATA_OK
}
func LPushS(key string, value interface{}) string {
	err := REDISDB.RedisClient.LPush(context.Background(), key, value).Err()
	if err != nil {
		logger.Error(err)
		return utils.RECODE_STOREDATA_FAILED
	}
	return utils.RECODE_STOREDATA_OK
}

func LRem(key string, count int64, value interface{}) string {
	err := REDISDB.RedisClient.LRem(context.Background(), key, count, value)
	if err.Err() != nil {
		logger.Error(err.Err())
		return utils.RECODE_NODATA
	}
	return utils.RECODE_OK
}

func Del(key string) error {
	err := REDISDB.RedisClient.Del(context.Background(), key).Err()
	return err
}

func LPush(key string, value interface{}) error {
	err := REDISDB.RedisClient.LPush(context.Background(), key, value).Err()
	return err
}

func RPop(key string) (string, error) {
	result, err := REDISDB.RedisClient.RPop(context.Background(), key).Result()
	return result, err
}

func LLen(key string) (int, error) {
	result, err := REDISDB.RedisClient.LLen(context.Background(), key).Result()
	return int(result), err
}
func SetM(key string, value interface{}, expiration time.Duration) error {
	err := REDISDB.RedisClient.Set(context.Background(), key, value, expiration).Err()
	return err
}

func GetM(key string) (string, error) {
	val, err := REDISDB.RedisClient.Get(context.Background(), key).Result()
	return val, err
}

func SAdd(key string, member interface{}) (int64, error) {
	return REDISDB.RedisClient.SAdd(context.Background(), key, member).Result()
}

func SRem(key string, member interface{}) (int64, error) {
	return REDISDB.RedisClient.SRem(context.Background(), key, member).Result()
}

func SIsMember(key string, member interface{}) (bool, error) {
	return REDISDB.RedisClient.SIsMember(context.Background(), key, member).Result()
}

func SMembers(key string) ([]string, error) {
	return REDISDB.RedisClient.SMembers(context.Background(), key).Result()
}

func ZAdd(key string, members ...*redis.Z) (int64, error) {
	return REDISDB.RedisClient.ZAdd(context.Background(), key, members...).Result()
}

func ZRem(key string, members ...interface{}) (int64, error) {
	return REDISDB.RedisClient.ZRem(context.Background(), key, members...).Result()
}

func ZRemRangeByScore(key string, min string, max string) (int64, error) {
	return REDISDB.RedisClient.ZRemRangeByScore(context.Background(), key, min, max).Result()
}

func ZRank(key string, member string) (int64, error) {
	return REDISDB.RedisClient.ZRank(context.Background(), key, member).Result()
}

func HGet(key string, field string) (string, error) {
	return REDISDB.RedisClient.HGet(context.Background(), key, field).Result()
}

func HSet(key string, values ...interface{}) (int64, error) {
	return REDISDB.RedisClient.HSet(context.Background(), key, values...).Result()
}

func HIncrBy(key string, field string, incr int64) (int64, error) {
	return REDISDB.RedisClient.HIncrBy(context.Background(), key, field, incr).Result()
}

func SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return REDISDB.RedisClient.SetNX(context.Background(), key, value, expiration).Result()
}
