package util

import (
	"regexp"
)

//IsNumber 判断字符串是否全为数字
func IsNumber(str string) bool {
	regex := "^\\d+$"
	reg := regexp.MustCompile(regex)
	return reg.MatchString(str)
}

//
func InterToStr(reply interface{}) string {
	switch reply := reply.(type) {
	case []byte:
		return string(reply)
	case string:
		return reply
	case nil:
		return ""
	}
	return ""
}
