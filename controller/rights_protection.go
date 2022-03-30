package controller

import (
	"law/enum"
	"law/model"
	"law/utils"
	"time"

	"github.com/labstack/echo/v4"
)

//添加维权意向
func RightsProtectionAdd(ctx echo.Context) error {
	rightsProtection := &model.RightsProtection{}
	if err := ctx.Bind(rightsProtection); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	rightsProtection.CreatorUid = uid
	rightsProtection.CreateTime = int(time.Now().Unix())
	err := model.RightsProtectionAdd(rightsProtection)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("添加维权意向失败！", err.Error()))
	}
	err = model.ProtectionRetureVisitAdd(&model.ProtectionReturnVisit{
		CreatorUid: uid,
		ProtectionId: rightsProtection.Id,
		Classification: enum.NORETURN,
		CreateTime: int(time.Now().Unix()),
	})
	if err != nil {
		return ctx.JSON(utils.ErrIpt("添加默认回访失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", map[string]int{"rights_protection_id": rightsProtection.Id}))
}