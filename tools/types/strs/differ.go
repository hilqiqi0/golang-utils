package strs

import "github.com/hilqiqi0/golang-utils/common/consts"

var PerfThrDiff = 100

// s1对s2做差
func differStrByLoop(s1, s2 []string) []string {
	if IsEmptyStrs(s2) {
		return s1
	}
	var strs []string
	for _, i := range s1 {
		if IsNotContainStr(s2, i) {
			strs = append(strs, i)
		}
	}
	return strs
}

// 通过map主键唯一的特性过滤重复元素
func differStrByMap(s1, s2 []string) []string {
	var dup []string
	var tmpMap = make(map[string]byte, len(s2)) // 存放不重复主键
	for _, v := range s2 {
		tmpMap[v] = consts.ZERO
	}
	for _, i := range s1 {
		if _, ok := tmpMap[i]; !ok {
			dup = append(dup, i)
		}
	}
	return dup
}

// string间做去重
func DifferStrs(s1, s2 []string) []string {
	if len(s1) < PerfThrDiff {
		return differStrByLoop(s1, s2)
	}
	return differStrByMap(s1, s2)
}

func DifferStrsLen(s1, s2 []string, l int) []string {
	return CutStrList(DifferStrs(s1, s2), l)
}
