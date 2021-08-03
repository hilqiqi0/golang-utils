package converter

import (
	"strconv"

	"github.com/hilqiqi0/golang-utils/tools/types/nums"
)

/**
int 转换
*/
func IntToInt32(i int) int32 {
	return int32(i)
}

func IntToInt64(i int) int64 {
	return int64(i)
}

func IntToFloat32(a int) float32 {
	return float32(a)
}

func IntsToInt32s(a []int) []int32 {
	ret := make([]int32, len(a))
	for i, b := range a {
		ret[i] = IntToInt32(b)
	}
	return ret
}

func IntsToInt64s(a []int) []int64 {
	ret := make([]int64, len(a))
	for i, b := range a {
		ret[i] = IntToInt64(b)
	}
	return ret
}

func IntsToFloats(a []int) []float32 {
	ret := make([]float32, len(a))
	for i, b := range a {
		ret[i] = IntToFloat32(b)
	}
	return ret
}

/**
int32 转换
*/
func Int32ToInt(i int32) int {
	return int(i)
}

func Int32ToInt64(i int32) int64 {
	return int64(i)
}

func Int32ToFloat(a int32) float32 {
	return float32(a)
}

func Int32sToInts(a []int32) []int {
	ret := make([]int, len(a))
	for i, b := range a {
		ret[i] = Int32ToInt(b)
	}
	return ret
}

func Int32sToInt64s(a []int32) []int64 {
	ret := make([]int64, len(a))
	for i, b := range a {
		ret[i] = Int32ToInt64(b)
	}
	return ret
}

func Int32sToFloats(a []int32) []float32 {
	ret := make([]float32, len(a))
	for i, b := range a {
		ret[i] = Int32ToFloat(b)
	}
	return ret
}

/**
int64 转换
*/
func Int64ToInt(i int64) int {
	return int(i)
}

func Int64ToInt32(i int64) int32 {
	return int32(i)
}

func Int64ToFloat(a int64) float32 {
	return float32(a)
}

func Int64sToInts(a []int64) []int {
	ret := make([]int, len(a))
	for i, b := range a {
		ret[i] = Int64ToInt(b)
	}
	return ret
}

func Int64sToInt32s(a []int64) []int32 {
	ret := make([]int32, len(a))
	for i, b := range a {
		ret[i] = Int64ToInt32(b)
	}
	return ret
}

func Int64sToFloats(a []int64) []float32 {
	ret := make([]float32, len(a))
	for i, b := range a {
		ret[i] = Int64ToFloat(b)
	}
	return ret
}

/**
string 转换
*/
func IntToStr(a int) string {
	return strconv.Itoa(a)
}

func Int32ToStr(a int32) string {
	return strconv.FormatInt(int64(a), 10)
}

func Int64ToStr(a int64) string {
	return strconv.FormatInt(a, 10)
}

func FloatToStr(a float64) string {
	return strconv.FormatFloat(a, 'f', -1, 64)
}

func IntsToStrs(a []int) []string {
	if nums.IsEmptyInts(a) {
		return nil
	}
	ret := make([]string, len(a))
	for i, b := range a {
		ret[i] = IntToStr(b)
	}
	return ret
}

func Int32sToStrs(a []int32) []string {
	if nums.IsEmptyInt32s(a) {
		return nil
	}
	ret := make([]string, len(a))
	for i, b := range a {
		ret[i] = Int32ToStr(b)
	}
	return ret
}

func Int64sToStrs(a []int64) []string {
	if nums.IsEmpty(a) {
		return nil
	}
	ret := make([]string, len(a))
	for i, b := range a {
		ret[i] = Int64ToStr(b)
	}
	return ret
}
