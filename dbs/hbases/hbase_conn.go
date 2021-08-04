package hbases

import (
	"github.com/hilqiqi0/golang-utils/config"

	"github.com/tsuna/gohbase"
)

// var (
// 	HBaseClient config.HBaseDbInfo
// )

// func ConnectHBase() {
// 	conf := global.GetConfig()
// 	// user := gohbase.EffectiveUser(conf.HBaseCli.Account)
// 	// options := []gohbase.Option{user}
// 	HBaseClient.Zkquorum = conf.HBaseCli.Zkquorum
// 	HBaseClient.Namespace = conf.HBaseCli.Namespace
// 	// HBaseClient.Client = gohbase.NewClient(conf.HBaseCli.Zkquorum, options...)

// 	// print(conf)

// 	HBaseClient.Client = gohbase.NewClient(conf.HBaseCli.Thrift)
// 	return
// }

type HBaseDbInfo struct {
	// Zkquorum  string
	// Option    string
	Namespace string
	Client    gohbase.Client
}

func (that *HBaseDbInfo) GetHbaseConnFromConf(c *config.ConfigEngine, name string) {
	// redis_login := c.GetRedisDataFromConf(name)
	// that.redisDataDb = createSingleClient(redis_login)
	// that.poolSize = redis_login.Pool_size
	// fmt.Println("======", name)
	hbase_login := c.GetHbaseFromConf(name)
	// that.Zkquorum = hbase_login.Zkquorum
	that.Namespace = hbase_login.Namespace
	// fmt.Println("======", hbase_login)
	that.Client = gohbase.NewClient(hbase_login.Thrift)
}

func (that *HBaseDbInfo) GetDb() gohbase.Client {
	return that.Client
}
