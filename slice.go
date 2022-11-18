package goutil

import (
	"strconv"
	"strings"
)

// JoinInt concatenates integers in one string with defined separator
func JoinInt(array []int, sep string) string {
	var buf []string
	for _, v := range array {
		buf = append(buf, strconv.Itoa(v))
	}
	return strings.Join(buf, sep)
}
