package hbases_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hilqiqi0/golang-utils/config"
	"github.com/hilqiqi0/golang-utils/dbs/hbases"
	"github.com/hilqiqi0/golang-utils/tools/errs"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/filter"
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

// 过滤列，限制行
func TestScanOptionDataPageFilter(t *testing.T) {
	t.Log(Config_path)
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	t.Log(c)
	hbasedb := new(hbases.HBaseDbInfo)
	hbasedb.GetHbaseConnFromConf(&c, "Hbase")

	var hf []func(hrpc.Call) error
	// 只查询固定的列{cf: [col1, col2]}
	var Families map[string][]string
	Families = map[string][]string{"info": []string{"description"}}
	hf = append(hf, hrpc.Families(Families))

	var Filter filter.Filter
	// 限制返回条数
	Filter = filter.NewPageFilter(1)
	hf = append(hf, hrpc.Filters(Filter))

	infos, _ := hbasedb.HBaseScanOption("lobby_test:item_scene", hf...)

	for k, info := range infos {
		fmt.Println(k, info)
	}

	// 参考：https://www.cnblogs.com/P--K/p/11393862.html
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
	values := map[string]map[string][]byte{"f1": map[string][]byte{"col1": []byte("val1")}}
	// value := map[string]map[string][]byte{
	// 	"cf": { // 列族名, 与创建表时指定的名字保持一致
	// 		"col1": []byte("val1"),  // 列与值, 列名可自由定义
	// 		"col2": []byte("val2"),
	// 		"col3": []byte("val3"),
	// 	}
	// }
	putRequest, err := hrpc.NewPutStr(context.Background(), "t1", "r4", values)
	if err != nil {

	}
	rsp, err := client.Put(putRequest)
	if rsp != nil {

	}

}

func TestDeleteData(t *testing.T) {
	t.Log(Config_path)
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	t.Log(c)
	hbasedb := new(hbases.HBaseDbInfo)
	hbasedb.GetHbaseConnFromConf(&c, "Hbase")

	// client := hbasedb.GetDb()
	err1 := hbasedb.HbaseDelete("t1", "r4")

	fmt.Println(err1)

}

func TestPutDataMa(t *testing.T) {
	t.Log(Config_path)
	c := config.ConfigEngine{}
	var err error
	err = c.Load(Config_path)
	errs.CheckCommonErr(err)
	t.Log(c)
	hbasedb := new(hbases.HBaseDbInfo)
	hbasedb.GetHbaseConnFromConf(&c, "Hbase")

	// client := hbasedb.GetDb()
	values := map[string][]byte{"a": []byte("val1")}
	err1 := hbasedb.HBasePut("t1", "r7", "f1", values)

	fmt.Println(err1)

}

// 参考：https://github.com/tsuna/gohbase
// https://blog.csdn.net/Ssxy0606/article/details/99945479
