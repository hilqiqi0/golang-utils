package errs

/*
有关报错打印的封装
*/
import (
	"errors"
	"reflect"
	"runtime"
	"strconv"

	"github.com/cihub/seelog"
)

//错误处理函数
func CheckCommonInfo(err string, params ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	seelog.Infof(file+":"+strconv.Itoa(line)+" "+err, params...)
}

//错误处理函数
func CheckCommonDebug(err string, params ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	seelog.Debugf(file+":"+strconv.Itoa(line)+" "+err, params...)
}

//错误处理函数
func CheckCommonWarn(err string, params ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	_ = seelog.Warnf(file+":"+strconv.Itoa(line)+" "+err, params...)
}

//错误处理函数
func CheckCommonErr(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		_ = seelog.Error(file, ":", line, err)
	}
}

//错误处理函数
func CheckFatalErr(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		_ = seelog.Critical("Important error:", file, ":", line, err)
		panic(err)
	}
}

func CheckEmptyValue(val interface{}) {
	if reflect.TypeOf(val).Kind() == reflect.Int {
		if val.(int) == 0 {
			panic("this value shouldn't be 0")
		}
	} else if reflect.TypeOf(val).Kind() == reflect.Int64 {
		if val.(int64) == 0 {
			panic(`this value shouldn't be 0`)
		}
	} else if reflect.TypeOf(val).Kind() == reflect.String {
		if val.(string) == "" {
			panic(`this value shouldn't be ""`)
		}
	} else if reflect.TypeOf(val).Kind() == reflect.Float32 {
		if val.(float32) == 0.0 {
			panic(`this value shouldn't be 0.0`)
		}
	} else if reflect.TypeOf(val).Kind() == reflect.Float64 {
		if val.(float64) == 0.0 {
			panic(`this value shouldn't be 0.0`)
		}
	} else if reflect.TypeOf(val).Kind() == reflect.Slice {
		if len(val.([]interface{})) == 0 {
			panic(`this value shouldn't be empty slice`)
		}
	} else if reflect.TypeOf(val).Kind() == reflect.Map {
		if len(val.(map[interface{}]interface{})) == 0 {
			panic(`this value shouldn't be empty map`)
		}
	}
}

func CheckValueStat(v, min, max int) int {
	if v > max {
		CheckCommonErr(errors.New("input value is too large!!!"))
		return max
	}
	if v < min {
		CheckCommonErr(errors.New("input value is too small!!!"))
		return min
	}
	return v
}
