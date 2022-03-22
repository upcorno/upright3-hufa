package controller

import (
	"law/conf"
	"law/model"
	"law/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func UserLogin(ctx echo.Context) error {
	tokenStr := ctx.QueryParam("token")
	if tokenStr == "" {
		return ctx.JSON(utils.ErrJwt("token不可为空"))
	}
	loginClaims, err := parseLoginToken(tokenStr)
	if err != nil {
		return ctx.JSON(utils.ErrJwt("token解析失败", err.Error()))
	}
	ysjUid := loginClaims.Uid
	user, has, err := model.UserGetByYsjUid(ysjUid)
	if err != nil {
		return ctx.JSON(utils.ErrJwt("用户信息查询失败", err.Error()))
	}
	if has && int(loginClaims.ExpiresAt) <= user.JwtLoginTokenExp {
		return ctx.JSON(utils.ErrJwt("该jwt已使用或已过期"))
	}
	if err := addOrUpdateUser(has, loginClaims); err != nil {
		return ctx.JSON(utils.ErrJwt("用户信息更新失败", err.Error()))
	}
	authTokenStr, err := utils.CreateAuthToken(loginClaims.Uid, ctx.RealIP())
	if err != nil {
		return ctx.JSON(utils.ErrJwt("创建auth token失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", authTokenStr))
}

func addOrUpdateUser(has bool, loginClaims *loginClaims) error {
	ysjUid := loginClaims.Uid
	user := &model.User{
		YsjUid:           ysjUid,
		Name:             loginClaims.Name,
		AvatarUrl:        loginClaims.AvatarUrl,
		Roles:            loginClaims.Roles,
		JwtLoginTokenExp: int(loginClaims.ExpiresAt),
	}
	if has {
		return model.UserUpdateByYsjUid(ysjUid, user)
	}
	return model.UserCreate(user)
}

func UserList(ctx echo.Context) error {
	roleStr := ctx.QueryParam("role")
	if roleStr != "" {
		usersList, err := model.UsersList(roleStr)
		if err != nil {
			return ctx.JSON(utils.ErrIpt("获取users list失败", err.Error()))
		}
		return ctx.JSON(utils.Succ("success", usersList))
	} else {
		usersList, err := model.AllUsersList()
		if err != nil {
			return ctx.JSON(utils.ErrIpt("获取all users list失败", err.Error()))
		}
		return ctx.JSON(utils.Succ("success", usersList))
	}
}

func UserInfo(ctx echo.Context) error {
	uid := ctx.Get("uid").(int)
	userInfo, err := model.UserInfo(uid)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取user info失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", userInfo))

}

type loginClaims struct {
	Uid       int
	Name      string
	AvatarUrl string
	Roles     []string
	jwt.StandardClaims
}

func parseLoginToken(tokenStr string) (*loginClaims, error) {
	claims, err := utils.ParseToken(tokenStr, &loginClaims{}, conf.App.Jwt.LoginKey) //conf.App.Jwt.AuthKey= youshangjiao
	if err != nil {
		return nil, err
	}
	return claims.(*loginClaims), nil
}

func CurrentUserUCurrencyAmount(ctx echo.Context) error {
	var amount int
	UCurrencyAmount := map[string]int{
		"u_currency_amount": amount,
	}
	return ctx.JSON(utils.Succ("success", UCurrencyAmount))

}
