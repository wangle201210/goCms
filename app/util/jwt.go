package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(id int, username string, role int) (string, error) {
	nowTime := time.Now()
	//ExpireTime
	expireTime := nowTime.Add(time.Minute * time.Duration(AppSetting.ExpireTime))

	claims := Claims{
		id,
		username,
		role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "goCms",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func init() {
	jwtSecret = []byte(AppSetting.JwtSecret)
}
