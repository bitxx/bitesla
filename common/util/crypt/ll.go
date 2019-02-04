package crypt

import (
	"crypto"
	"encoding/base64"
	"math/rand"
	"time"
)

func LLEncrypt(plaintext, public_key []byte) string {
	hmack_key := genLetterDigitRandom(32)
	version := "lianpay1_0_1"
	aes_key := genLetterDigitRandom(32)
	nonce := genLetterDigitRandom(8)
	return lianlianpayEncrypt(plaintext, public_key, hmack_key, version, aes_key, nonce)
}

func lianlianpayEncrypt(plaintext, pu_key []byte, hmack_key, version, aes_key, nonce string) string {
	B64hmack_key, _ := RsaEncrypt([]byte(hmack_key), pu_key)
	B64aes_key, _ := RsaEncrypt([]byte(aes_key), pu_key)
	B64nonce := base64.StdEncoding.EncodeToString([]byte(nonce))
	aesData, _ := AesCTREncrypt(plaintext, []byte(aes_key), append([]byte(nonce), 0, 0, 0, 0, 0, 0, 0, 1))
	encry := base64.StdEncoding.EncodeToString(aesData)
	message := B64nonce + "$" + encry
	sign := HmacEncryptToBase64([]byte(message), []byte(hmack_key), crypto.SHA256)
	return version + "$" + B64hmack_key + "$" + B64aes_key + "$" + B64nonce + "$" + encry + "$" + sign
}

//随机数种子
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func genLetterDigitRandom(size int) string {
	allLetterDigit := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	randomSb := ""
	digitSize := len(allLetterDigit)
	for i := 0; i < size; i++ {
		randomSb += allLetterDigit[rnd.Intn(digitSize)]
	}
	return randomSb
}
