package controller

import (
	"law/model"
	"law/service"
	"law/utils"

	"github.com/labstack/echo/v4"
)

func Login(ctx echo.Context) error {
	code := &wxCode{}
	if err := ctx.Bind(code); err != nil {
		return ctx.JSON(utils.ErrIpt("参数解析失败！", err.Error()))
	}
	uid, err := service.Login(code.Code)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("登录失败！", err.Error()))
	}
	token, err := utils.CreateAuthToken(uid, ctx.RealIP())
	if err != nil {
		return ctx.JSON(utils.ErrIpt("token 生成失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", map[string]string{"token": token}))
}

type wxCode struct {
	Code string `json:"code"`
}

func SetPhone(ctx echo.Context) error {
	code := &wxCode{}
	if err := ctx.Bind(code); err != nil {
		return ctx.JSON(utils.ErrIpt("参数解析失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	err := service.SetPhone(uid, code.Code)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("设置手机号失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", nil))
}

type nameAndAvatarUrl struct {
	NickName  string `json:"nickName"`
	AvatarUrl string `json:"avatarUrl"`
}

func SetNameAndAvatarUrl(ctx echo.Context) error {
	nameAndUrl := &nameAndAvatarUrl{}
	if err := ctx.Bind(nameAndUrl); err != nil {
		return ctx.JSON(utils.ErrIpt("参数解析失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	err := service.SetNameAndAvatarUrl(uid, nameAndUrl.NickName, nameAndUrl.AvatarUrl)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("设置昵称和头像失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", nil))
}

func GetUserInfo(ctx echo.Context) error {
	uid := ctx.Get("uid").(int)
	user := model.User{Id: uid}
	has, err := user.Get()
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取用户信息失败！", err.Error()))
	}
	if has {
		return ctx.JSON(utils.ErrIpt("用户 id 不存在！", uid))
	}
	return ctx.JSON(utils.Succ("success", user))
}
