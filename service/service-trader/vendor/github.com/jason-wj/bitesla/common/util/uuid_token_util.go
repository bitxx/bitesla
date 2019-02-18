package util

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
	"time"
)

//用于生成appkey、appSecret

func GetKey() (string, error) {
	id, err := uuid.NewV4()
	return id.String(), err
}

func GetSecret(key string) string {
	sum := md5.Sum([]byte(key))
	return fmt.Sprintf("%x", sum)
}

func GetID(key string) string {
	sum := md5.Sum([]byte(key))
	return fmt.Sprintf("%x", sum)
}

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

func GetToken(key, secret, issuer, jwtSecret string, d time.Duration) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(d)
	claims := Claims{
		key,
		secret,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString([]byte(jwtSecret))
}

func ParseToken(token, jwtSecret string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, err
}
