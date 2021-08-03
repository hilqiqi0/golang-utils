package nums

import (
	"github.com/hilqiqi0/golang-utils/common/consts"
)

func IsEmptyInts(s []int) bool {
	return len(s) == consts.ZERO
}

func IsEmptyInt32s(s []int32) bool {
	return len(s) == consts.ZERO
}

func IsEmptyInt64s(s []int64) bool {
	return len(s) == consts.ZERO
}

func IsEmpty(s ...interface{}) bool {
	return len(s) == consts.ZERO
}
