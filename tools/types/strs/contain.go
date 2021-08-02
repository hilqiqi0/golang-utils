package strs

import (
	"reflect"
	"strings"
)

// split
func ContainStrNum(str string, sep string) int {
	return len(strings.Split(str, sep)) - 1
}

func IsContainStr(slice []string, s string) bool {
	if IsEmptyStrs(slice) {
		return false
	}
	for _, i := range slice {
		if s == i {
			return true
		}
	}
	return false
}

func IsNotContainStr(slice []string, s string) bool {
	return !IsContainStr(slice, s)
}

// Returns true if slice strings contains any of the strings.
func IsContainAny(slice []string, checks []string) bool {
	for _, s := range checks {
		if IsContainStr(slice, s) {
			return true
		}
	}
	return false
}

func IsContainsMap(maps []map[string]string, item map[string]string) bool {
	for _, m := range maps {
		if reflect.DeepEqual(m, item) {
			return true
		}
	}
	return false
}
