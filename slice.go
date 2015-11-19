package goutil

import (
	"strings"
	"strconv"
)

func MemberOfSliceInt(element int, slice []int) bool {
	for _, i := range slice {
		if element == i {
			return true
		}
	}
	return false
}

func JoinInt(array []int, sep string) string {
	var buf []string
	for _, v := range array {
		buf = append(buf, strconv.Itoa(v))
	}
	return strings.Join(buf, sep)
}
