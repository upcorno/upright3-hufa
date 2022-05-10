package controller

import (
	"law/model"
	"law/service"
	"law/utils"

	"github.com/labstack/echo/v4"
)

type wxCredential struct {
	Code string `json:"code" form:"code" query:"code" validate:"required,min=1"`
}

func Login(ctx echo.Context) error {
	credential := &wxCredential{}
	if err := utils.BindAndValidate(ctx, credential); err != nil {
		return ctx.JSON(utils.ErrIpt("参数解析失败！", err.Error()))
	}
	token, err := service.User.Login(credential.Code)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("登录失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", map[string]string{"token": token}))
}

func SetPhone(ctx echo.Context) error {
	credential := &wxCredential{}
	if err := utils.BindAndValidate(ctx, credential); err != nil {
		return ctx.JSON(utils.ErrIpt("参数解析失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	err := service.User.SetPhone(uid, credential.Code)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("设置手机号失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", nil))
}

type nameAndAvatarUrl struct {
	NickName  string `json:"nick_name" form:"nick_name" query:"nick_name" validate:"required"`
	AvatarUrl string `json:"avatar_url" form:"avatar_url" query:"avatar_url" validate:"required,url"`
}

func SetNameAndAvatarUrl(ctx echo.Context) error {
	nameAndUrl := &nameAndAvatarUrl{}
	if err := utils.BindAndValidate(ctx, nameAndUrl); err != nil {
		return ctx.JSON(utils.ErrIpt("参数解析失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	err := service.User.SetNameAndAvatarUrl(uid, nameAndUrl.NickName, nameAndUrl.AvatarUrl)
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
	if !has {
		return ctx.JSON(utils.ErrIpt("用户 id 不存在！", uid))
	}
	return ctx.JSON(utils.Succ("success", user))
}

type accountAndPassWord struct {
	Account  string `json:"account" form:"account" query:"account" validate:"required"`
	Password string `json:"password" form:"password" query:"password" validate:"required"`
}

func BackendLogin(ctx echo.Context) error {
	accountAndPassWord := &accountAndPassWord{}
	if err := utils.BindAndValidate(ctx, accountAndPassWord); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析失败！", err.Error()))
	}
	token, err := service.User.BackendLogin(accountAndPassWord.Account, accountAndPassWord.Password)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("登录失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", map[string]string{"token": token}))
}
