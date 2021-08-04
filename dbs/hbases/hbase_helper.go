package hbases

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"runtime/debug"

	"time"

	"github.com/hilqiqi0/golang-utils/tools/errs"
	"github.com/tsuna/gohbase/hrpc"
)

type HBaseConfig struct {
	Table      string
	RowKey     string
	ColumnName string
}

func (that *HBaseDbInfo) HBasePut(table, rowkey, columnName string, data map[string][]byte) (err error) {
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
		errs.CheckCommonErr(fmt.Errorf(fmt.Sprintf("NewPutStr error, rowKey=%s,error:%s", rowkey, err)))
		return err
	}
	_, err = that.Client.Put(putRequest)
	if err != nil {
		errs.CheckCommonErr(fmt.Errorf(fmt.Sprintf("client.Put error, rowKey=%s,error:%s", rowkey, err)))
		return err
	} else {
		errs.CheckCommonErr(fmt.Errorf(fmt.Sprintf("client.Put new record, table=%s, rowKey=%s", table, rowkey)))
	}
	return nil
}

func (that *HBaseDbInfo) HBaseScan(table string) (info []map[string]string, err error) {
	defer func() {
		if e := recover(); e != nil {
			errs.CheckCommonErr(fmt.Errorf(fmt.Sprintf("Invoke recall failed: %s, trace:\n%s", e, debug.Stack())))
			err = fmt.Errorf("put hbase error")
			return
		}
	}()
	var resultMaps []map[string]string

	getRequest, err := hrpc.NewScanStr(context.Background(), table)
	if err != nil {
		errs.CheckCommonErr(fmt.Errorf(fmt.Sprintf("NewScan error, table=%s,error:%s", table, err)))
		return resultMaps, err
	}
	scanResult := that.Client.Scan(getRequest)

	if err != nil {
		errs.CheckCommonErr(fmt.Errorf(fmt.Sprintf("Scan error, table=%s,error:%s", table, err)))
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

func (that *HBaseDbInfo) HBaseQuery(table, rowkey string) (info map[string]string, err error) {
	defer func() {
		if e := recover(); e != nil {
			errs.CheckCommonErr(fmt.Errorf(fmt.Sprintf("nvoke recall failed: %s, trace:\n%s", e, debug.Stack())))
			err = fmt.Errorf("put hbase error")
			return
		}
	}()

	resultMap := make(map[string]string)
	getRequest, err := hrpc.NewGetStr(context.Background(), table, rowkey)
	if err != nil {
		errs.CheckCommonErr(fmt.Errorf(fmt.Sprintf("NewGetStr error, rowKey=%s,error:%s", rowkey, err)))
		return resultMap, err
	}
	getRsp, err := that.Client.Get(getRequest)
	if err != nil {
		errs.CheckCommonErr(fmt.Errorf(fmt.Sprintf("Get error, rowKey=%s,error:%s", rowkey, err)))
		return resultMap, err
	}

	for _, cell := range getRsp.Cells {
		resultMap[string(cell.Qualifier)] = string(cell.Value)
	}
	return resultMap, nil
}

//指定表，通过options筛选数据，例如Families函数，或者filter函数
func (that *HBaseDbInfo) GetsByOption(table string, rowkey string, options func(hrpc.Call) error) (info map[string]string, err error) {
	defer func() {
		if errs := recover(); errs != nil {
			switch fmt.Sprintf("%v", errs) {
			case "runtime error: index out of range":
				err = errors.New("NoSuchRowKeyOrQualifierException")
			case "runtime error: invalid memory address or nil pointer dereference":
				err = errors.New("NoSuchColFamilyException")
			default:
				err = fmt.Errorf("%v", errs)
			}
			return
		}
	}()

	resultMap := make(map[string]string)
	getRequest, err := hrpc.NewGetStr(context.Background(), table, rowkey, options)
	if err != nil {
		errs.CheckCommonErr(fmt.Errorf(fmt.Sprintf("hrpc.NewGetStr: %s", err.Error())))
	}
	getRsp, err := that.Client.Get(getRequest)
	if err != nil {
		errs.CheckCommonErr(fmt.Errorf(fmt.Sprintf("hbase clients: %s", err.Error())))
	}

	for _, cell := range getRsp.Cells {
		resultMap[string(cell.Qualifier)] = string(cell.Value)
	}
	return resultMap, nil
}
