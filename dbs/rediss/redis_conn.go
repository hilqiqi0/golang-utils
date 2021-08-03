package rediss

import (
	"github.com/go-redis/redis"
	"github.com/hilqiqi0/golang-utils/config"
	"github.com/hilqiqi0/golang-utils/tools/errs"
)

type RedisDbInfo struct {
	redisDataDb *redis.Client
	poolSize    int
}

func createSingleClient(redisConf *config.RedisData) *redis.Client {
	redisdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		PoolSize: redisConf.Pool_size,
		// Password: redisConf.Password,
		// Database to be selected after connecting to the server.
		DB: redisConf.Db,
	})
	_, err := redisdb.Ping().Result()
	errs.CheckFatalErr(err)
	return redisdb
}

func (that *RedisDbInfo) GetRedisConnFromConf(c *config.ConfigEngine, name string) {
	redis_login := c.GetRedisDataFromConf(name)
	that.redisDataDb = createSingleClient(redis_login)
	that.poolSize = redis_login.Pool_size
}

// func (that *RedisDbInfo) CreateSingleClient(addr, password string, poolSize int) {
// 	redisDB := redis.NewClient(&redis.Options{
// 		Addr:     addr,
// 		PoolSize: poolSize,
// 		// Password: password,
// 	})
// 	_, err := redisDB.Ping().Result()
// 	errs.CheckFatalErr(err)
// 	that.redisDataDb = redisDB
// 	that.poolSize = poolSize
// }

func (that *RedisDbInfo) GetDb() *redis.Client {
	return that.redisDataDb
}
