package utils

import (
	"strconv"
)

func Uint64ToString(number uint64) string {
	s := strconv.FormatUint(number, 10)
	return s
}

func StringToUint64(str string) uint64 {
	r, _ := strconv.ParseUint(str, 10, 64)
	return r
}