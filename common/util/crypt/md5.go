package crypt

import (
	"crypto/md5"
	"encoding/hex"
)

//md5加密
func MD5(data string) string {
	m := md5.Sum([]byte(data))
	return hex.EncodeToString(m[:])
	//return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

//md5加密16位
func MD5Encrypt16(data string) string {
	m := md5.Sum([]byte(data))
	return hex.EncodeToString(m[:])[8:24]
}
