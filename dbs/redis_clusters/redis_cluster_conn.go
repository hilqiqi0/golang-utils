package redis_clusters

import (
	"github.com/hilqiqi0/golang-utils/config"
	"github.com/hilqiqi0/golang-utils/tools/errs"

	"github.com/go-redis/redis"
)

type RedisClusterConnMethod interface {
	GetRedisClusterConnFromConf(c *config.ConfigEngine, name string)
	CreateSingleClient(addr, password string, poolSize int)
}

/*
有关redis cluster连接的封装
*/
type RedisClusterDbInfo struct {
	RedisClusterDataDb *redis.ClusterClient
	PoolSize           int
}

func createClusterClient(redisConf *config.RedisClusterData) *redis.ClusterClient {
	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    redisConf.Addrs,
		PoolSize: redisConf.Pool_size,
		// Password:     redisConf.Password,
	})
	_, err := redisdb.Ping().Result()
	errs.CheckFatalErr(err)
	return redisdb
}

func (redisClusterDbInfo *RedisClusterDbInfo) GetRedisClusterConnFromConf(c *config.ConfigEngine, name string) {
	redis_login := c.GetRedisClusterDataFromConf(name)
	redisClusterDbInfo.RedisClusterDataDb = createClusterClient(redis_login)
	redisClusterDbInfo.PoolSize = redis_login.Pool_size
}

func (redisClusterDbInfo *RedisClusterDbInfo) GetDb() *redis.ClusterClient {
	return redisClusterDbInfo.RedisClusterDataDb
}
