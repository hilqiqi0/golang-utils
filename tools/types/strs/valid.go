package strs

import (
	"strings"

	"github.com/hilqiqi0/golang-utils/common/consts"
)

// 有效字符串判断
func IsEmptyStrs(strs []string) bool {
	return len(strs) == consts.ZERO
}

func IsNotEmptyStrs(strs []string) bool {
	return !IsEmptyStrs(strs)
}

// IsEmpty returns true if the string is empty
func IsEmptyStr(str string) bool {
	return len(str) == 0
}

// IsNotEmpty returns true if the string is not empty
func IsNotEmptyStr(str string) bool {
	return !IsEmptyStr(str)
}

// IsBlank returns true if the string is blank (all whitespace)
func IsBlankStr(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// IsNotBlank returns true if the string is not blank
func IsNotBlankStr(str string) bool {
	return !IsBlankStr(str)
}
