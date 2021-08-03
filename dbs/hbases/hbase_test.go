package hbases_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/hilqiqi0/golang-utils/config"
	"github.com/hilqiqi0/golang-utils/dbs/hbases"
	"github.com/hilqiqi0/golang-utils/tools/errs"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

const (
	Config_path = "../../config/test/hbase_test.yaml"
)

type ItemInfo struct {
	ItemId   string `json:"item_id"`
	ItemType string `json:"item_type"`
}

func (item *ItemInfo) Item2Map() map[string][]byte {
	infoMap := make(map[string][]byte)
	val := reflect.ValueOf(*item)
	for i := 0; i < val.Type().NumField(); i++ {
		jsonTag := val.Type().Field(i).Tag.Get("json")
		infoMap[jsonTag] = []byte(val.Field(i).String())
	}
	return infoMap
}
func TestGetHbaseItemRangeFromConf(t *testing.T) {
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	// redisItem := new(rediss.ItemInfo)
	// hbaseItem := new(hbases.HBaseConfig)
	// redisItem.GetRedisItemFromConf(&c, "Redis_items.test_item2")
	// t.Log(redisItem)
	// redisdb := new(rediss.RedisDbInfo)
	hbasedb := new(hbases.HBaseDbInfo)
	// redisdb.GetRedisConnFromConf(&c, "Redis")
	hbasedb.GetHbaseConnFromConf(&c, "Hbase")
	// t.Log(&hbasedb)
	var req ItemInfo
	req.ItemId = "asss"
	req.ItemType = "qqqq"
	data := req.Item2Map()
	fmt.Println(data)
	hbases.HBasePut(hbasedb.GetDb(), "lobby_test:item_scene", "100", "info", data)
	// if redisdb.GetDb().Ping().Val() != "PONG" {
	// 	t.Error("can't connect to redis db")
	// }
	// err = redisItem.ItemZAdd(redisdb.GetDb(), converter.IntsToStrs([]int{3, 4, 5, 6, 7}), "just redis_test")
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("this error")
	// }
	// cmd, err := redisItem.ItemGetZRange(redisdb.GetDb(), "just redis_test")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(cmd)
}

func TestGetDate(t *testing.T) {
	// thrift := "ip-172-17-49-77.us-west-2.compute.internal"
	// client := gohbase.NewClient(thrift)

	// infos, _ := hbases.HBaseScan(client, "lobby_test:item_scene")

	// for _, info := range infos {
	// 	fmt.Println(info)
	// }

	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	hbasedb := new(hbases.HBaseDbInfo)
	hbasedb.GetHbaseConnFromConf(&c, "Hbase")

	infos, _ := hbases.HBaseScan(hbasedb.GetDb(), "lobby_test:item_scene")

	for _, info := range infos {
		fmt.Println(info)
	}
}

func TestPutDate(t *testing.T) {
	thrift := "ip-172-17-49-77.us-west-2.compute.internal"
	client := gohbase.NewClient(thrift)

	// Values maps a ColumnFamily -> Qualifiers -> Values.
	values := map[string]map[string][]byte{"cf": map[string][]byte{"a": []byte{0}}}
	putRequest, err := hrpc.NewPutStr(context.Background(), "test", "key", values)
	if err != nil {

	}
	rsp, err := client.Put(putRequest)
	if rsp != nil {

	}

}

// 参考：https://github.com/tsuna/gohbase
