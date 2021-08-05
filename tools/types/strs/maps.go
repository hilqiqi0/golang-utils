package strs

import (
	"sort"
)

func GetMapKeysAndValues(m map[string]string) ([]string, []string) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var values []string
	for _, k := range keys {
		values = append(values, m[k])
	}
	return keys, values
}

func Strs2Map(s []string, value string) (m map[string]string) {
	m = make(map[string]string, len(s))
	for _, ss := range s {
		m[ss] = value
	}
	return
}
