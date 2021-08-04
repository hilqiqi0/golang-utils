package strs

import (
	"fmt"

	"github.com/hilqiqi0/golang-utils/common/consts"
	"github.com/hilqiqi0/golang-utils/tools/maths"
)

// cut strs
func CutStrList(s []string, l int) []string {
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(s))
	return s[:size:size]
}

// cut strs
func CutStrListAndFilling(s []string, f []string, l int) []string {
	if cut := CutStrList(s, l); len(cut) < l {
		return UniqueStrsLen(append(s, f...), l)
	} else {
		return cut
	}
}

// string list Loop
func GetStrListNoLoop(s []string, size, num int) ([]string, error) {
	if num <= consts.ZERO || size <= consts.ZERO {
		return nil, fmt.Errorf("Input paras error!!!")
	}
	minThr := maths.MinInt(size*(num-1), len(s))
	maxThr := maths.MinInt(size*num, len(s))
	return s[minThr:maxThr], nil
}

func GetStrListLoop(s []string, size, num int) ([]string, error) {
	var sLen = len(s)
	if num <= consts.ZERO || size <= consts.ZERO || size >= sLen {
		return nil, fmt.Errorf("Input parameter error!!!")
	}
	start := (size * (num - 1)) % sLen
	end := (num * size) % sLen
	if start < end {
		return s[start:end:end], nil
	} else {
		var out []string
		out = append(out, s[start:sLen:sLen]...)
		out = append(out, s[:end:end]...)
		return out, nil
	}
}
