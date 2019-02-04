package crypt

import (
	"crypto"
	"crypto/hmac"
	// "crypto/sha1"
	// "fmt"
	// "io"
	"encoding/base64"
	"encoding/hex"
)

func HmacEncrypt(origData, key []byte, hash crypto.Hash) string {
	mac := hmac.New(hash.New, key)
	mac.Write(origData)
	return hex.EncodeToString(mac.Sum(nil))
}

func HmacEncryptToBase64(origData, key []byte, hash crypto.Hash) string {
	mac := hmac.New(hash.New, key)
	mac.Write(origData)
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
