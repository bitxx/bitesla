package util

import "regexp"

//IsNumber 判断字符串是否全为数字
func IsNumber(str string) bool {
	regex := "^\\d+$"
	reg := regexp.MustCompile(regex)
	return reg.MatchString(str)
}
