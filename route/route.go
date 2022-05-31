package route

import (
	"law/controller"
	"law/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

var backenPrefix string = "/backend"

func InitRouter(e *echo.Echo) {
	e.Use(utils.MidAuth)
	bg := e.Group(backenPrefix)
	bg.Use(utils.BackendAuth)
	///法律问题
	e.Add(http.MethodGet, "/legal_issue/list", controller.LegalIssueList)
	e.Add(http.MethodGet, "/legal_issue/category_list", controller.LegalIssueCategoryList)
	e.Add(http.MethodGet, "/legal_issue/get", controller.LegalIssueGet)
	e.Add(http.MethodPost, "/legal_issue/favorite", controller.LegalIssueFavorite)
	e.Add(http.MethodPost, "/legal_issue/cancel_favorite", controller.LegalIssueCancelFavorite)
	e.Add(http.MethodGet, "/legal_issue/is_favorite", controller.LegalIssueIsFavorite)
	///
	///咨询
	e.Add(http.MethodPost, "/consultation/create", controller.ConsultationCreate)
	e.Add(http.MethodGet, "/consultation/get", controller.ConsultationGet)
	e.Add(http.MethodGet, "/consultation/list", controller.ConsultationList)
	e.Add(http.MethodPost, "/consultation/set_status", controller.ConsultationSetStatus)
	e.Add(http.MethodPost, "/consultation/add_reply", controller.ConsultationAddReply)
	e.Add(http.MethodGet, "/consultation/list_reply", controller.ConsultationListReply)
	e.Add(http.MethodGet, "/consultation/file_upload_auth", controller.ConsultationFileUploadAuth)
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
	e.Add(http.MethodPost, "/user/login", controller.Login)
	e.Add(http.MethodPost, "/user/set_phone", controller.SetPhone)
	e.Add(http.MethodPost, "/user/set_name_and_avatar_url", controller.SetNameAndAvatarUrl)
	e.Add(http.MethodGet, "/user/get_user_info", controller.GetUserInfo)
	e.Add(http.MethodGet, "/wx/notify", controller.WxNotify)
	e.Add(http.MethodPost, "/wx/notify", controller.WxNotify)
	//
	bg.GET("/user/get_user_info", controller.GetUserInfo)
	bg.POST("/user/login", controller.BackendLogin)
	///
}

func addCooperationRoute(e *echo.Echo, bg *echo.Group, prefix string) {
	e.Add(http.MethodPost, prefix+"add", controller.CooperationController.Add)
	e.Add(http.MethodGet, prefix+"get", controller.CooperationController.Get)
	e.Add(http.MethodPost, prefix+"update_base_info", controller.CooperationController.UpdateBaseInfo)
	//
	bg.GET(prefix+"get", controller.CooperationController.BgGet)
	bg.POST(prefix+"set_deal_info", controller.CooperationController.SetDealInfo)
	bg.GET(prefix+"list", controller.CooperationController.BackendList)
}
