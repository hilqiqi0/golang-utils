package redis_clusters_test

import (
	"fmt"

	"testing"
	"time"

	"github.com/hilqiqi0/golang-utils/config"
	"github.com/hilqiqi0/golang-utils/dbs/redis_clusters"
	"github.com/hilqiqi0/golang-utils/tools/errs"
)

const (
	Config_path = "../../config/test/redis_clusters_test.yaml"
)

func TestGetRedisClusterItemFromConfNew(t *testing.T) {
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	redisItem := redis_clusters.ItemInfo{}
	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
	t.Log(redisItem)
	redisCli := new(redis_clusters.RedisClusterDbInfo)
	redisCli.GetRedisClusterConnFromConf(&c, "RedisCluster")
	t.Log(redisCli)
	if redisCli.RedisClusterDataDb.Ping().Val() != "PONG" {
		t.Error("can't connect to redis cluster db")
	}
	if err = redisItem.ItemSet(redisCli.RedisClusterDataDb, "just redis_test", "test2"); err != nil {
		fmt.Println(err)
	}
	cmd, err := redisItem.ItemGet(redisCli.RedisClusterDataDb, "test2")
	if err != nil {
		fmt.Println(err)
	}
	println(cmd.Val())
	if cmd.Val() != "just redis_test" {
		t.Error("set redis item in redis db failed")
	}

}
func TestGetRedisClusterItemFromConf(t *testing.T) {
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	redisItem := redis_clusters.ItemInfo{}
	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
	t.Log(redisItem)
	redisCli := new(redis_clusters.RedisClusterDbInfo)
	redisCli.GetRedisClusterConnFromConf(&c, "RedisCluster")
	t.Log(redisCli)
	if redisCli.GetDb().Ping().Val() != "PONG" {
		t.Error("can't connect to redis cluster db")
	}
	if err = redisItem.ItemSet(redisCli.GetDb(), "just redis_test", "test1"); err != nil {
		fmt.Println(err)
	}
	cmd, err := redisItem.ItemGet(redisCli.GetDb(), "test1")
	if err != nil {
		fmt.Println(err)
	}
	if cmd.Val() != "just redis_test" {
		t.Error("set redis item in redis db failed")
	}
}

func TestGetRedisItemRangeFromConf(t *testing.T) {
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	redisItem := new(redis_clusters.ItemInfo)
	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item2")
	t.Log(redisItem)
	redisdb := new(redis_clusters.RedisClusterDbInfo)
	redisdb.GetRedisClusterConnFromConf(&c, "RedisCluster")
	t.Log(redisdb)
	if redisdb.GetDb().Ping().Val() != "PONG" {
		t.Error("can't connect to redis db")
	}
	// err = redisItem.ItemZAdd(redisdb.GetDb(), converter.IntsToStrs([]int{3, 4, 5, 6, 7}), "just redis_test")
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("this error")
	// }
	cmd, err := redisItem.ItemGetZRange(redisdb.GetDb(), "just redis_test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cmd)
}

func TestGetRedisItemFromConf(t *testing.T) {
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	if nil != err {
		panic(err)
	}
	redisItem := new(redis_clusters.ItemInfo)
	redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item")
	fmt.Println("redisItem.GetExpire().Seconds()", redisItem.GetExpire().Seconds())
	t.Log(redisItem)
	redisdb := new(redis_clusters.RedisClusterDbInfo)
	redisdb.GetRedisClusterConnFromConf(&c, "Redis")
	//t.Log(redisdb)
	if redisdb.GetDb().Ping().Val() != "PONG" {
		t.Error("can't connect to redis db")
	}
	err = redisItem.ItemZAdd(redisdb.GetDb(), []string{"ugender2", "ok1"}, "just redis_test")
	if err != nil {
		fmt.Println(err)
		fmt.Println("this error")
	}
	time.Sleep(5 * time.Second)
	cmd, err := redisItem.ItemGetZRange(redisdb.GetDb(), "just redis_test")
	if err != nil {
		fmt.Println(err)
	}
	for _, x := range cmd {
		fmt.Println(x)
	}
	fmt.Println("获取数据长度：", len(cmd))
}
