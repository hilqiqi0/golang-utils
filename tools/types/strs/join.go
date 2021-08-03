package strs

import "strings"

// join
func JoinStrs(sep string, str ...string) string {
	return strings.Join(str, sep)
}

func JoinStrSlice(sep string, str []string) string {
	return strings.Join(str, sep)
}
