package converter

import (
	"strconv"

	"github.com/hilqiqi0/golang-utils/tools/errs"

	"github.com/hilqiqi0/golang-utils/tools/types/strs"
)

// str2other
func StrToByte(s string) []byte {
	return []byte(s)
}

func StrToInt(a string, defaultInt int) int {
	if r, err := strconv.Atoi(a); err == nil {
		return r
	} else {
		errs.CheckCommonErr(err)
		return defaultInt
	}
}

func StrToInt32(a string, defaultInt int32) int32 {
	if r, err := strconv.ParseInt(a, 10, 32); err == nil {
		return int32(r)
	} else {
		errs.CheckCommonErr(err)
		return defaultInt
	}
}

func StrToInt64(a string, defaultInt int64) int64 {
	if r, err := strconv.ParseInt(a, 10, 64); err == nil {
		return r
	} else {
		errs.CheckCommonErr(err)
		return defaultInt
	}
}

func StrToFloat32(a string, defaultFloat float32) float32 {
	if r, err := strconv.ParseFloat(a, 32); err == nil {
		return float32(r)
	} else {
		errs.CheckCommonErr(err)
		return defaultFloat
	}
}

func StrToFloat64(a string, defaultInt float64) float64 {
	if r, err := strconv.ParseFloat(a, 64); err == nil {
		return r
	} else {
		errs.CheckCommonErr(err)
		return defaultInt
	}
}

func StrsToInts(a []string, defaultInt int) []int {
	if strs.IsEmptyStrs(a) {
		return nil
	}
	ret := make([]int, len(a))
	for i, b := range a {
		ret[i] = StrToInt(b, defaultInt)
	}
	return ret
}

func StrsToInt32s(a []string, defaultInt int32) []int32 {
	if strs.IsEmptyStrs(a) {
		return nil
	}
	ret := make([]int32, len(a))
	for i, b := range a {
		ret[i] = StrToInt32(b, defaultInt)
	}
	return ret
}

func StrsToInt64s(a []string, defaultInt int64) []int64 {
	if strs.IsEmptyStrs(a) {
		return nil
	}
	ret := make([]int64, len(a))
	for i, b := range a {
		ret[i] = StrToInt64(b, defaultInt)
	}
	return ret
}
