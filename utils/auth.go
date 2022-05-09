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
		ctx.Set("uid", claims.Uid)
		return next(ctx)
	}
}

//后台接口权限验证
func BackendAuth(next echo.HandlerFunc) echo.HandlerFunc {
	bgUids := map[int]bool{}
	for _, bgAccountInfo := range *conf.App.BgAccounts {
		bgUids[bgAccountInfo.Uid] = true
	}
	return func(ctx echo.Context) error {
		if ctx.Request().URL.Path == conf.App.Jwt.BackendLoginPath {
			return next(ctx)
		}
		uid := ctx.Get("uid").(int)
		if _, ok := bgUids[uid]; !ok {
			return ctx.JSON(ErrJwt("非法访问后台接口。"))
		}
		return next(ctx)
	}
}

type authClaims struct {
	Uid int
	jwt.StandardClaims
}

func CreateAuthToken(uid int) (string, error) {
	nowSecond := int64(time.Now().Unix())
	expireAtSecond := nowSecond + int64(conf.App.Jwt.AuthLifetime)
	claims := &authClaims{
		Uid: uid,
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
	return nil, err
}
