package rdb

import (
	"FriendlyAlmond_backend/pkg/logger"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redisdb struct {
	RedisClient *redis.Client
	LockTimeout time.Duration // 循环获取锁的等待超时时间
	LockExpire  time.Duration // 生存时间
	LockSleep   time.Duration // 循环等待时间
	LockPrefix  string
	locks       *lockMap
}

var (
	REDISDB *Redisdb
	once    sync.Once
)

func InitRedis(opts *redis.Options) error {
	var err error
	once.Do(func() {
		rdb := redis.NewClient(opts)
		err = rdb.Ping(context.TODO()).Err()
		if err != nil {
			logger.Error("The connection of redis is failed.")

		}
		REDISDB = &Redisdb{
			RedisClient: rdb,
			LockTimeout: time.Second,
			LockExpire:  time.Second,
			LockSleep:   time.Second,
			LockPrefix:  "lock:",
			locks:       NewlockMap(),
		}
		logger.Infof("connecting redis: %s successfully", opts.Addr)
	})

	return err
}

// Lock
// 传入key 锁住key
func Lock(name string, times ...time.Duration) bool {
	key := fmt.Sprintf("%s%s", REDISDB.LockPrefix, name)
	var expire = REDISDB.LockExpire
	var timeout = REDISDB.LockTimeout
	var sleep = REDISDB.LockSleep
	if len(times) > 0 {
		timeout = times[0]
	}
	if len(times) > 1 {
		expire = times[1]
	}
	if len(times) > 2 {
		sleep = times[2]
	}
	expireAt := time.Now().Add(expire).UnixNano()
	timeoutAt := time.Now().Add(timeout).UnixNano()
	for {
		// 写入成功直接返回
		if ok, _ := REDISDB.RedisClient.SetNX(context.Background(), key, expireAt, expire).Result(); ok {
			REDISDB.locks.add(name, fmt.Sprint(expireAt))
			return true
		}
		// 不成功获取ttl
		ttl := REDISDB.RedisClient.TTL(context.Background(), key).Val()
		// ttl 小于0 可以重新写入
		if ttl.Nanoseconds() < 0 {
			// 写入成功直接返回
			if ok, _ := REDISDB.RedisClient.SetNX(context.Background(), key, expireAt, expire).Result(); ok {
				REDISDB.locks.add(name, fmt.Sprint(expireAt))
				return true
			}
		}
		// 如果超时
		if timeoutAt < time.Now().UnixNano() {
			break
		}
		// 1miao 循环一次
		time.Sleep(sleep)
	}
	return false
}

// UnLock
// 传入key 解锁key
func UnLock(name string) bool {
	if ex := REDISDB.locks.find(name); ex != "" {
		key := fmt.Sprintf("%s%s", REDISDB.LockPrefix, name)
		if strings.Compare(ex, REDISDB.RedisClient.Get(context.Background(), key).Val()) == 0 {
			if k := REDISDB.RedisClient.Del(context.Background(), key).Val(); k > 0 {
				REDISDB.locks.del(name)
				return true
			}
		}
	}
	return false
}

// IsLock .
func IsLock(name string) bool {
	if ttl := REDISDB.RedisClient.TTL(context.Background(), fmt.Sprintf("%s%s", REDISDB.LockPrefix, name)).Val(); ttl.Nanoseconds() > 0 {
		return true
	}
	return false
}

type lockMap struct {
	Data map[string]string
	ch   chan func()
}

func NewlockMap() *lockMap {
	m := &lockMap{
		Data: make(map[string]string),
		ch:   make(chan func()),
	}
	go func() {
		for {
			(<-m.ch)()
		}
	}()
	return m
}
func (m *lockMap) add(num string, data string) {
	m.ch <- func() {
		m.Data[num] = data
	}
}
func (m *lockMap) del(num string) {
	m.ch <- func() {
		delete(m.Data, num)
	}
}
func (m *lockMap) find(num string) (data string) {
	m.ch <- func() {
		if res, ok := m.Data[num]; ok {
			data = res
		}
	}
	return
}
