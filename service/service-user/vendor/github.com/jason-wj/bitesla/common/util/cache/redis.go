package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"time"
)

var (
	redisCache *Cache
)

const (
	DefaultTimeout = 30 * time.Minute
)

// Cache is Redis cache adapter.
type Cache struct {
	p           *redis.Pool // redis connection pool
	Url         string
	DBIndex     int
	Key         string
	Password    string
	MaxIdle     int
	MaxActive   string
	IdleTimeout int
}

func GetRedisCache() (*Cache, error) {
	if redisCache == nil {
		return nil, errors.New("please init cache before")
	}
	return redisCache, nil
}

// NewRedisCache create new redis cache with default collection name.
func Init(cache *Cache) error {
	err := cache.StartAndGC()
	if err != nil {
		return err
	}
	redisCache = cache
	return nil
}

// Get cache from redis.
func (rc *Cache) Get(key string) interface{} {
	if v, err := rc.Do("GET", key); err == nil {
		return v
	}
	return nil
}

// Get cache from redis.
func (rc *Cache) GetStructByJson(key string, obj interface{}) error {
	var v interface{}
	var err error
	v, err = rc.Do("GET", key)
	if err != nil {
		return err
	}
	if v == nil {
		return errors.New("no data match the key")
	}
	tmp := fmt.Sprintf("%s", v)
	err = json.Unmarshal([]byte(tmp), &obj)
	return err
}

func (rc *Cache) RedisBytes(key string) (data []byte, err error) {
	data, err = redis.Bytes(rc.Get(key), err)
	return
}

func (rc *Cache) RedisString(key string) (data string, err error) {
	data, err = redis.String(rc.Get(key), err)
	return
}

func (rc *Cache) RedisInt(key string) (data int, err error) {
	data, err = redis.Int(rc.Get(key), err)
	return
}

// PutJson 若timeout为-1，则使用默认的超时时间
//timeout单位秒
func (rc *Cache) PutJson(key string, val interface{}, timeout time.Duration) error {
	val, err := json.Marshal(val)
	if err != nil {
		return err
	}

	if _, err = rc.Do("SETEX", key, int64(timeout/time.Second), val); err != nil {
		return err
	}

	if _, err = rc.Do("HSET", rc.Key, key, true); err != nil {
		return err
	}
	return err
}

// Put put cache to redis.
func (rc *Cache) Put(key string, val interface{}, timeout time.Duration) error {

	var err error

	if _, err = rc.Do("SETEX", key, int64(timeout/time.Second), val); err != nil {
		return err
	}

	if _, err = rc.Do("HSET", rc.Key, key, true); err != nil {
		return err
	}
	return err
}

func (rc *Cache) SetNX(key string, val interface{}, timeout time.Duration) bool {
	if timeout < 0 {
		timeout = DefaultTimeout
	}
	if result, err := rc.Do("SET", key, val, "NX", "EX", int64(timeout/time.Second)); err == nil && result == "OK" {
		return true
	} else {
		return false
	}
}

// Delete delete cache in redis.
func (rc *Cache) Delete(key string) error {
	var err error
	if _, err = rc.Do("DEL", key); err != nil {
		return err
	}
	_, err = rc.Do("HDEL", rc.Key, key)
	return err
}

// IsExist check cache's existence in redis.
func (rc *Cache) IsExist(key string) bool {
	v, err := redis.Bool(rc.Do("EXISTS", key))
	if err != nil {
		return false
	}
	if v == false {
		if _, err = rc.Do("HDEL", rc.Key, key); err != nil {
			return false
		}
	}
	return v
}

// Put put cache to redis.
func (rc *Cache) LPush(key string, val interface{}) error {
	data, _ := json.Marshal(val)
	_, err := rc.Do("lpush", key, data)
	return err
}

func (rc *Cache) Brpop(key string, callback func([]byte)) {
	if reply, err := rc.Do("brpop", key, 1); err == nil && reply != nil {
		if values, err1 := redis.Values(reply, err); err1 == nil {
			value, ok := values[1].([]byte)
			if ok {
				callback(value)
			} else {
				fmt.Println("assert is wrong")
			}
		} else {
			fmt.Println("assert is wrong!")
		}
	}
}

func (rc *Cache) GetRedisTTL(key string) time.Duration {
	reply, _ := rc.Do("TTL", key)
	if value, ok := reply.(int64); ok {
		return time.Duration(value)
	}
	return 0
}

// Decr decrease counter in redis.
func (rc *Cache) Incrby(key string, num int) (interface{}, error) {
	return rc.Do("INCRBY", key, num)
}

// actually do the redis cmds
func (rc *Cache) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	c := rc.p.Get()
	defer c.Close()

	return c.Do(commandName, args...)
}

// ClearAll clean all cache in redis. delete this redis collection.
func (rc *Cache) ClearAll() error {
	cachedKeys, err := redis.Strings(rc.Do("HKEYS", rc.Key))
	if err != nil {
		return err
	}
	for _, str := range cachedKeys {
		if _, err = rc.Do("DEL", str); err != nil {
			return err
		}
	}
	_, err = rc.Do("DEL", rc.Key)
	return err
}

// StartAndGC start redis cache adapter.
// config is like {"key":"collection key","conn":"connection info","dbNum":"0"}
// the cache item in redis are stored forever,
// so no gc operation.
func (rc *Cache) StartAndGC() error {
	rc.connectInit()
	//初始化
	c := rc.p.Get()
	defer c.Close()
	return c.Err()
}

func (rc *Cache) Lock(key string, timeout time.Duration) bool {
	for true {
		result, err := rc.Do("SET", key, 1, "NX", "EX", int64(timeout/time.Second))
		if err == nil && result == "OK" {
			return true
		}
		if err != nil {
			fmt.Println("Locker Lock Error", err)
			break
		}
		sleepTimeInterval()
	}
	return false
}

// connect to redis.
func (rc *Cache) connectInit() {
	dialFunc := func() (c redis.Conn, err error) {
		c, err = redis.Dial("tcp", rc.Url)
		if err != nil {
			return nil, err
		}

		if rc.Password != "" {
			if _, err := c.Do("AUTH", rc.Password); err != nil {
				c.Close()
				return nil, err
			}
		}

		_, selecterr := c.Do("SELECT", rc.DBIndex)
		if selecterr != nil {
			c.Close()
			return nil, selecterr
		}
		return
	}
	// initialize a new pool
	rc.p = &redis.Pool{
		MaxIdle:     rc.MaxIdle,
		IdleTimeout: time.Duration(rc.IdleTimeout) * time.Second,
		Dial:        dialFunc,
	}
}

// 当一次Lock失败之后,会随机在RetryMinTimeInterval和RetryMaxTimeInterval之间
// 取一个随机值作为下次请求锁的时间间隔
// Lock请求锁的最小时间间隔(毫秒)
var RetryMinTimeInterval int64 = 5

// Lock请求锁的最大时间间隔(毫秒)
var RetryMaxTimeInterval int64 = 30

// sleepTimeInterval 随机休眠一段时间
// 随机时间范围[RetryMinTimeInterval,RetryMaxTimeInterval)
func sleepTimeInterval() {
	var unixNano = time.Now().UnixNano()
	var r = rand.New(rand.NewSource(unixNano))
	var randValue = RetryMinTimeInterval + r.Int63n(RetryMaxTimeInterval-RetryMinTimeInterval)
	time.Sleep(time.Duration(randValue) * time.Millisecond)
}
