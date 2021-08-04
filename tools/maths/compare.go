package maths

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/hilqiqi0/golang-utils/common/consts"

	"github.com/pkg/errors"
)

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinInt32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// Int64Min returns the minimum of the params
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func MinFloat(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func MinInts(a []int) int {
	if consts.ZERO == len(a) {
		return consts.ZERO
	}
	var min = a[consts.ZERO]
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func MinInt32s(a []int32) int32 {
	if consts.ZERO == len(a) {
		return consts.ZERO
	}
	var min = a[consts.ZERO]
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func MinFloats(a []float32) float32 {
	if consts.ZERO == len(a) {
		return consts.ZERO
	}
	var min = a[consts.ZERO]
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MaxInt32(a, b int32) int32 {
	if a < b {
		return b
	}
	return a
}

// Int64Max returns the maximum of the params
func MaxInt64(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func MaxFloat(a, b float32) float32 {
	if a < b {
		return b
	}
	return a
}

func MaxInts(a []int) int {
	if consts.ZERO == len(a) {
		return consts.ZERO
	}
	var max = a[consts.ZERO]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

func MaxInt32s(a []int32) int32 {
	if consts.ZERO == len(a) {
		return consts.ZERO
	}
	var max = a[consts.ZERO]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

func MaxFloats(a []float32) float32 {
	if consts.ZERO == len(a) {
		return consts.ZERO
	}
	var max = a[consts.ZERO]
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}

func Min(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			return MinInt(a.(int), b.(int))
		case float32:
			return MinFloat(a.(float32), b.(float32))
		default:
			return errors.New("input num type error!!!")
		}
	case float32:
		switch b.(type) {
		case int:
			return MinFloat(a.(float32), b.(float32))
		case float32:
			return MinFloat(a.(float32), b.(float32))
		default:
			return errors.New("input num type error!!!")
		}
	default:
		return errors.New("input num type error!!!")
	}
}

func Max(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			return MaxInt(a.(int), b.(int))
		case float32:
			return MaxFloat(a.(float32), b.(float32))
		default:
			return errors.New("input num type error!!!")
		}
	case float32:
		switch b.(type) {
		case int:
			return MaxFloat(a.(float32), b.(float32))
		case float32:
			return MaxFloat(a.(float32), b.(float32))
		default:
			return errors.New("input num type error!!!")
		}
	default:
		return errors.New("input num type error!!!")
	}
}

// RoundToInt32 rounds floats into integer nums.
func RoundToInt32(a float64) int32 {
	if a < 0 {
		return int32(a - 0.5)
	}
	return int32(a + 0.5)
}

// 小数点后 n 位 - 四舍五入
func RoundedFixed(val float64, n int) float64 {
	shift := math.Pow(10, float64(n))
	fv := 0.0000000001 + val //对浮点数产生.xxx999999999 计算不准进行处理
	return math.Floor(fv*shift+.5) / shift
}

// 小数点后 n 位 - 舍去
func TruncRound(val float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n+1)+"f", val)
	temp := strings.Split(floatStr, ".")
	var newFloat string
	if len(temp) < 2 || n >= len(temp[1]) {
		newFloat = floatStr
	} else {
		newFloat = temp[0] + "." + temp[1][:n]
	}
	inst, _ := strconv.ParseFloat(newFloat, 64)
	return inst
}
