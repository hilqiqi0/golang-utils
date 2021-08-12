package hbases_test

import (
	"context"
	"fmt"
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

func TestScanData(t *testing.T) {
	t.Log(Config_path)
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	t.Log(c)
	hbasedb := new(hbases.HBaseDbInfo)
	hbasedb.GetHbaseConnFromConf(&c, "Hbase")

	// client := hbasedb.GetDb()
	infos, _ := hbasedb.HBaseScan("lobby_test:item_scene")

	for _, info := range infos {
		fmt.Println(info)
	}
}

func TestScanOptionData(t *testing.T) {
	t.Log(Config_path)
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	t.Log(c)
	hbasedb := new(hbases.HBaseDbInfo)
	hbasedb.GetHbaseConnFromConf(&c, "Hbase")

	// client := hbasedb.GetDb()
	f := map[string][]string{"info": []string{"description"}}
	infos, _ := hbasedb.HBaseScanOption("lobby_test:item_scene", hrpc.Families(f))

	for k, info := range infos {
		fmt.Println(k, info)
	}
}
func TestGetData(t *testing.T) {
	t.Log(Config_path)
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	t.Log(c)
	hbasedb := new(hbases.HBaseDbInfo)
	hbasedb.GetHbaseConnFromConf(&c, "Hbase")

	// client := hbasedb.GetDb()
	infos, _ := hbasedb.HBaseQuery("lobby_test:item_scene", "237")

	for k, v := range infos {
		fmt.Println(k, v)
	}
}

func TestGetFilterData(t *testing.T) {
	t.Log(Config_path)
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	t.Log(c)
	hbasedb := new(hbases.HBaseDbInfo)
	hbasedb.GetHbaseConnFromConf(&c, "Hbase")

	f := map[string][]string{"info": []string{"description"}}
	infos, _ := hbasedb.GetsByOption("lobby_test:item_scene", "237", hrpc.Families(f))

	for k, v := range infos {
		fmt.Println(k, v)
	}
}
func TestPutData(t *testing.T) {
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
// https://blog.csdn.net/Ssxy0606/article/details/99945479
