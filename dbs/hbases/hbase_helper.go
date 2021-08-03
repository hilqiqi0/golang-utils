package hbases

import (
	"context"
	"fmt"
	"reflect"

	"time"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

type HBaseConfig struct {
	Table      string
	RowKey     string
	ColumnName string
}

func HBasePut(hbaseClient gohbase.Client, table, rowkey, columnName string, data map[string][]byte) (err error) {
	defer func() {
		if e := recover(); e != nil {
			// global.Logger.Errorf("Invoke recall failed: %s, trace:\n%s", e, debug.Stack())
			err = fmt.Errorf("put hbase error")
			return
		}
	}()
	infoData := make(map[string]map[string][]byte)
	infoData[columnName] = data
	ctx, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond) // 200ms, 本地映射后会增长时间
	defer cancel()
	putRequest, err := hrpc.NewPutStr(ctx, table, rowkey, infoData)
	if err != nil {
		// global.Logger.Errorf("NewPutStr error, rowKey=%s,error:%s", config.RowKey, err)
		return err
	}
	_, err = hbaseClient.Put(putRequest)
	if err != nil {
		// global.Logger.Errorf("client.Put error, rowKey=%s,error:%s", config.RowKey, err)
		return err
	} else {
		// global.Logger.Infof("client.Put new record, table=%s, rowKey=%s", config.Table, config.RowKey)
	}
	return nil
}

func HBaseQuery(hbaseClient gohbase.Client, table, rowkey string) (info map[string]string, err error) {
	defer func() {
		if e := recover(); e != nil {
			// global.Logger.Errorf("Invoke recall failed: %s, trace:\n%s", e, debug.Stack())
			err = fmt.Errorf("put hbase error")
			return
		}
	}()

	resultMap := make(map[string]string)
	getRequest, err := hrpc.NewGetStr(context.Background(), table, rowkey)
	if err != nil {
		// global.Logger.Error("NewGetStr error, rowKey=%s,error:%s", rowkey, err)
		return resultMap, err
	}
	getRsp, err := hbaseClient.Get(getRequest)
	if err != nil {
		// global.Logger.Error("Get error, rowKey=%s,error:%s", rowkey, err)
		return resultMap, err
	}

	for _, cell := range getRsp.Cells {
		resultMap[string(cell.Qualifier)] = string(cell.Value)
	}
	return resultMap, nil
}

func HBaseScan(hbaseClient gohbase.Client, table string) (info []map[string]string, err error) {
	defer func() {
		if e := recover(); e != nil {
			// global.Logger.Errorf("Invoke recall failed: %s, trace:\n%s", e, debug.Stack())
			err = fmt.Errorf("put hbase error")
			return
		}
	}()
	var resultMaps []map[string]string

	getRequest, err := hrpc.NewScanStr(context.Background(), table)
	if err != nil {
		// global.Logger.Error("NewScan error, table=%s,error:%s", table, err)
		return resultMaps, err
	}
	scanResult := hbaseClient.Scan(getRequest)

	if err != nil {
		// global.Logger.Error("Scan error, table=%s,error:%s", table, err)
		return resultMaps, err
	}
	for {
		getRsp, err := scanResult.Next()
		if err != nil {
			break
		}
		resultMap := make(map[string]string)
		for _, cell := range getRsp.Cells {
			resultMap[string(cell.Qualifier)] = string(cell.Value)
		}
		resultMaps = append(resultMaps, resultMap)
	}

	return resultMaps, nil
}

func Map2Info(info interface{}, infoMap map[string]string) {
	rVal := reflect.ValueOf(info).Elem()
	for i := 0; i < rVal.NumField(); i++ {
		if fieldType := rVal.Field(i).Kind(); fieldType == reflect.String && rVal.Field(i).CanSet() {
			jsonTag := rVal.Type().Field(i).Tag.Get("json")
			if v, ok := infoMap[jsonTag]; ok {
				f := rVal.Field(i)
				f.Set(reflect.ValueOf(v))
			}
		}
	}
}
