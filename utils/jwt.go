package utils

import (
	"fmt"
	"github.com/c479096292/Spinach_blog/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = config.ConfObj.SecretKey

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	StandardClaims jwt.StandardClaims
}

func (c Claims) Valid() error {
	return fmt.Errorf("generate token failed")
}

func GenerateToken(username, password string)  string {
	nowTime := time.Now()
	expireTime := nowTime.Add(1 * time.Hour)
	claims := Claims{
		Username:EncodeMD5(username),
		Password:EncodeMD5(password),
		StandardClaims:jwt.StandardClaims{
			ExpiresAt:expireTime.Unix(),
			Issuer:   "cole",
		},
	}
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err !=nil{
		config.Warn(fmt.Sprintf("create token failed: %s", err))
	}
	return tokenString
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(jwtSecret), nil
	})
	if tokenClaims != nil{
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}


