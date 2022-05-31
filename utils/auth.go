package utils

import (
	"law/conf"
	"law/enum"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var bgUids map[int]bool
var nonAuthPath map[string]bool

func initPublicVar() {
	if bgUids == nil {
		bgUids = map[int]bool{}
		for _, bgAccountInfo := range *conf.App.BgAccounts {
			bgUids[bgAccountInfo.Uid] = true
		}
	}
	if nonAuthPath == nil {
		nonAuthPath = map[string]bool{}
		for _, path := range conf.App.Jwt.NonAuthPath {
			nonAuthPath[path] = true
		}
	}
}

var IsBackend enum.YesOrNo = enum.NO

// midAuth 登录认证中间件
func MidAuth(next echo.HandlerFunc) echo.HandlerFunc {
	initPublicVar()
	return func(ctx echo.Context) error {
		IsBackend = enum.YES
		if _, ok := nonAuthPath[ctx.Request().URL.Path]; ok {
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
	initPublicVar()
	return func(ctx echo.Context) error {
		ctx.Set("is_backend", enum.YES)
		if _, ok := nonAuthPath[ctx.Request().URL.Path]; ok {
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
	jwt.RegisteredClaims
}

func CreateAuthToken(uid int) (string, error) {
	now := time.Now().Local()
	expireAt := now.Add(time.Second * time.Duration(conf.App.Jwt.AuthLifetime))
	claims := &authClaims{
		Uid: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
			NotBefore: jwt.NewNumericDate(now),
		},
	}
	return CreateToken(claims, conf.App.Jwt.AuthKey)
}

func parseAuthToken(tokenStr string) (*authClaims, error) {
	claims, err := ParseAndValidToken(tokenStr, &authClaims{}, conf.App.Jwt.AuthKey)
	if err != nil {
		return nil, err
	}
	if claims, ok := claims.(*authClaims); ok {
		return claims, nil
	}
	return nil, err
}
