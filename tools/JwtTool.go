package tools

import (
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"log"
	"time"
)

/**
@Title JWT(Json Web Token)的生成、刷新和解析
@CreateTime 2022年9月1日10:26:18
@Author 薛智敏
*/

type Users struct {
	ExeEmail        string    `json:"LoginEmail"`      //被执行的邮箱
	LoginStatus     bool      `json:"loginStatus"`     //登录状态:True 表示已经登录
	SendEmailStatus bool      `json:"sendEmailStatus"` //发送验证码的状态:True 表示已经发送
	Captcha         string    `json:"captcha"`         //验证码
	ExpirationTime  time.Time `json:"expirationTime"`  //过期时间
}

type CustomClaims struct {
	Users
	jwt.StandardClaims
}

var MySecret = []byte("密钥")

//  GenToken 创建 Token
//  @Description:
//  @param user Token携带的字段
//  @param times  过期时间
//  @return string 加密的Token
//  @return error

func GenToken(user Users, times time.Duration) (string, error) {
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(times)), //过期的时间
			Issuer:    "XueZhiMin",                   //签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

//  ParseToken // 解析 token
//  @Description:
//  @param tokenStr 加密的Token字段
//  @return *CustomClaims
//  @return error
//

func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		log.Println(" token parse err:", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

//
//  RefreshToken 刷新 Token
//  @Description:
//  @param tokenStr Token字段
//  @param times 重新设置的时间
//  @return string Token
//  @return error
//

func RefreshToken(tokenStr string, times time.Duration) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = jwt.At(time.Now().Add(times))
		return GenToken(claims.Users, times)
	}
	return "", errors.New("Cloudn't handle this token")
}
