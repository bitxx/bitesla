package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Pasword  string `json:"password"`
	UserId   int64  `json:"userId"`
	jwt.StandardClaims
}

func GetToken(username, password, issuer, jwtSecret string, userId int64, d time.Duration) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(d)
	claims := Claims{
		username,
		password,
		userId,
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
