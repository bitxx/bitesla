package pay

import (
	// "bytes"
	"crypto"
	// "encoding/json"
	"errors"
	"fmt"
	// "net/url"
	// "sort"
	"strings"
	"zcm_tools/crypt"
	"zcm_tools/http"
)

//签名加密方式
const (
	EncryptMD5           = iota //md5加密
	EncryptSHA1withRsa          //SHA1withRsa
	EncryptSHA256withRsa        //SHA256withRsa
	EncryptMD5withRsa           //MD5withRsa
)

//请求参数结构
const (
	ContentTypeJson = iota
	ContentTypeForm
)

//需要添加sign的请求 SHA1withRsa，SHA256withRsa，MD5withRsa，MD5
func DoRequest(param http.Values, request_url, key string, contentType, encryptType int) ([]byte, error) {
	sign, err := Sign(param, key, encryptType)
	if err != nil {
		return nil, err
	}
	param.Add("sign", sign)
	strContentType := ""
	data := ""
	if contentType == ContentTypeJson {
		data = param.ToString()
		strContentType = "application/json;charset=UTF-8"
	} else {
		strContentType = "application/x-www-form-urlencoded;charset=utf-8"
		data = param.Encode()
	}
	fmt.Println(request_url, data, strContentType)
	if strings.HasPrefix(request_url, "https://") {
		return http.HttpsPost(request_url, data, strContentType)
	} else {
		return http.HttpPost(request_url, data, strContentType)
	}
}

//签名 SHA1withRsa，SHA256withRsa，MD5withRsa，MD5
func Sign(param http.Values, key string, encryptType int) (string, error) {
	signData := param.GetSignDataNoSpace() //
	fmt.Println(signData)
	sign := ""
	var err error
	switch encryptType {
	case EncryptSHA1withRsa:
		fmt.Println("SHA1withRsa")
		sign, err = crypt.SignPKCS1v15([]byte(signData), []byte(key), crypto.SHA1)
		if err != nil {
			return "", err
		}
	case EncryptSHA256withRsa:
		fmt.Println("SHA256withRsa")
		sign, err = crypt.SignPKCS1v15([]byte(signData), []byte(key), crypto.SHA256)
		if err != nil {
			return "", err
		}
	case EncryptMD5withRsa:
		fmt.Println("MD5withRsa")
		sign, err = crypt.SignPKCS1v15([]byte(signData), []byte(key), crypto.MD5)
		if err != nil {
			return "", err
		}
	case EncryptMD5:
		fmt.Println("MD5")
		sign = crypt.MD5(signData + key)
	default:
		return "", errors.New("缺少加密方式")
	}
	return sign, nil

}

//验签
func VerifySign(param http.Values, key string, encryptType int) bool {
	sign := param.Get("sign")
	param.Del("sign")
	calSign, _ := Sign(param, key, encryptType)
	if calSign == sign {
		return true
	}
	return false
}
