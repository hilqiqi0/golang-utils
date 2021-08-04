package strs

import (
	"github.com/hilqiqi0/golang-utils/common/consts"
	"github.com/hilqiqi0/golang-utils/tools/maths"
)

// ints mix
func MixListStr(s ...[]string) []string {
	var count int
	for _, item := range s {
		count += len(item)
	}
	mix := make([]string, consts.ZERO, count)
	for _, ss := range s {
		mix = append(mix, ss...)
	}
	return UniqueStrs(mix)
}

// ints mix
func MixListStrV2(s ...[]string) []string {
	var count int
	var maxLen int
	for _, item := range s {
		count += len(item)
		maxLen = maths.MaxInt(maxLen, len(item))
	}
	mix := make([]string, consts.ZERO, count)
	for i := 0; i < maxLen; i++ {
		for j := range s {
			if i < len(s[j]) {
				mix = append(mix, s[j][i])
			}
		}
	}
	return UniqueStrs(mix)
}

func MixListStrLenV2(l int, s ...[]string) []string {
	mix := MixListStrV2(s...)
	size := maths.MinInt(maths.MaxInt(l, consts.ZERO), len(mix))
	return mix[:size:size]
}
