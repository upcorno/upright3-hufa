package utils

import (
	"law/conf"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// midAuth 登录认证中间件
func MidAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if ctx.Request().URL.Path == conf.App.Jwt.LoginPath || ctx.Request().URL.Path == conf.App.Jwt.BackendLoginPath {
			return next(ctx)
		}
		tokenRaw := ctx.Request().Header.Get("token")
		if tokenRaw == "" {
			return ctx.JSON(ErrJwt("token不可为空"))
		}
		claims, err := parseAuthToken(tokenRaw)
		if err != nil {
			return ctx.JSON(ErrJwt("请重新登陆", err.Error()))
		}
		if ctx.RealIP() != claims.Ip {
			return ctx.JSON(ErrJwt("网络变更,请重新登陆"))
		}
		ctx.Set("uid", claims.Uid) //存到了store字段里面
		return next(ctx)
	}
}

type authClaims struct {
	Uid int
	Ip  string
	jwt.StandardClaims
}

func CreateAuthToken(uid int, ip string) (string, error) {
	nowSecond := int64(time.Now().Unix())
	expireAtSecond := nowSecond + int64(conf.App.Jwt.AuthLifetime) //加了两小时,conf.App.Jwt.AuthLifetime= 7200
	claims := &authClaims{
		Uid: uid,
		Ip:  ip,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAtSecond,
			NotBefore: nowSecond,
		},
	}
	return CreateToken(claims, conf.App.Jwt.AuthKey)
}

func parseAuthToken(tokenStr string) (*authClaims, error) {
	claims, err := ParseToken(tokenStr, &authClaims{}, conf.App.Jwt.AuthKey)
	if err != nil {
		return nil, err
	}
	if claims, ok := claims.(*authClaims); ok {
		return claims, nil
	}
	return nil, err //理论上一定执行不到
}
