package utils

import "strconv"

func Uint64ToString(number uint64) string {
	s := strconv.FormatUint(number, 10)
	return s
}