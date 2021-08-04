package strs

import (
	"github.com/hilqiqi0/golang-utils/common/consts"
)

var PerfThrUnique = 300

// int slice自去重,通过两重循环过滤重复元素
func uniqueStrsByLoop(s []string) []string {
	var strings = make([]string, consts.ZERO, len(s))
	for _, v := range s {
		if !IsContainStr(strings, v) {
			strings = append(strings, v)
		}
	}
	return strings
}

// 通过map主键唯一的特性过滤重复元素
func uniqueStrsByMap(s []string) []string {
	var strings = make([]string, consts.ZERO, len(s))
	var tmpMap = make(map[string]byte, len(s)) // 存放不重复主键
	for _, v := range s {
		if _, ok := tmpMap[v]; !ok {
			tmpMap[v] = consts.ZERO
			strings = append(strings, v)
		}
	}
	return strings
}

// 自去重
func UniqueStrs(s []string) []string {
	if len(s) < PerfThrUnique {
		return uniqueStrsByLoop(s)
	}
	return uniqueStrsByMap(s)
}

func UniqueStrsLen(s []string, l int) []string {
	return CutStrList(UniqueStrs(s), l)
}
