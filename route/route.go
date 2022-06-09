package route

import (
	"law/controller"
	"law/utils"

	"github.com/labstack/echo/v4"
)

var backenPrefix string = "/backend"

func InitRouter(e *echo.Echo) {
	e.Use(utils.BaseAuth)
	bg := e.Group(backenPrefix)
	bg.Use(utils.BackendAuth)
	///法律问题
	e.GET("/legal_issue/list", controller.LegalIssueContro.List)
	e.GET("/legal_issue/category_list", controller.LegalIssueContro.CategoryList)
	e.GET("/legal_issue/get", controller.LegalIssueContro.Get)
	e.POST("/legal_issue/favorite", controller.LegalIssueContro.Favorite)
	e.POST("/legal_issue/cancel_favorite", controller.LegalIssueContro.CancelFavorite)
	e.GET("/legal_issue/is_favorite", controller.LegalIssueContro.IsFavorite)
	//
	bg.GET("/legal_issue/get", controller.LegalIssueContro.Get)
	bg.GET("/legal_issue/list", controller.LegalIssueContro.List)
	bg.POST("/legal_issue/update", controller.LegalIssueContro.Update)
	bg.POST("/legal_issue/create", controller.LegalIssueContro.Create)
	bg.POST("/legal_issue/delete", controller.LegalIssueContro.Delete)
	///
	///咨询
	e.POST("/consultation/create", controller.ConsultationCreate)
	e.GET("/consultation/get", controller.ConsultationGet)
	e.GET("/consultation/list", controller.ConsultationList)
	e.POST("/consultation/set_status", controller.ConsultationSetStatus)
	e.POST("/consultation/add_reply", controller.ConsultationAddReply)
	e.GET("/consultation/list_reply", controller.ConsultationListReply)
	e.GET("/consultation/file_upload_auth", controller.ConsultationFileUploadAuth)
	//
	bg.POST("/consultation/set_status", controller.ConsultationSetStatus)
	bg.POST("/consultation/add_reply", controller.ConsultationAddReply)
	bg.GET("/consultation/list_reply", controller.ConsultationListReply)
	bg.GET("/consultation/get", controller.ConsultationGet)
	bg.GET("/consultation/file_upload_auth", controller.ConsultationBackendFileUploadAuth)
	bg.GET("/consultation/list", controller.ConsultationBackendList)
	///
	///侵权监测
	addCooperationRoute(e, bg, "/infringement_monitor/")
	///
	///我要维权
	addCooperationRoute(e, bg, "/rights_protection/")
	///
	///用户
	e.POST("/user/login", controller.Login)
	e.POST("/user/set_phone", controller.SetPhone)
	e.POST("/user/set_name_and_avatar_url", controller.SetNameAndAvatarUrl)
	e.GET("/user/get_user_info", controller.GetUserInfo)
	e.GET("/wx/notify", controller.WxNotify)
	e.POST("/wx/notify", controller.WxNotify)
	//
	bg.GET("/user/get_user_info", controller.GetUserInfo)
	bg.POST("/user/login", controller.BackendLogin)
	///
}

func addCooperationRoute(e *echo.Echo, bg *echo.Group, prefix string) {
	e.POST(prefix+"add", controller.CooperationController.Add)
	e.GET(prefix+"get", controller.CooperationController.Get)
	e.POST(prefix+"update_base_info", controller.CooperationController.UpdateBaseInfo)
	//
	bg.GET(prefix+"get", controller.CooperationController.BgGet)
	bg.POST(prefix+"set_deal_info", controller.CooperationController.SetDealInfo)
	bg.GET(prefix+"list", controller.CooperationController.BackendList)
}
