package app

import (
	"time"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/util"

	"github.com/golang-jwt/jwt"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/global"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

// getJWTSecret 获取jwt secret
func getJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// GenerateToken 生成Token
func GenerateToken(appKey, appSecret string) (string, error) {
	now := time.Now()
	expireTime := now.Add(global.JWTSetting.Expire)

	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(getJWTSecret())
}

// ParseToken 解析Token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getJWTSecret(), nil
	})

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
