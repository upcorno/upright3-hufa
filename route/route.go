package route

import (
	"law/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo, bg *echo.Group) {
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
	e.Add(http.MethodPost, "/infringement_monitor/add", controller.InfringementMonitorAdd)
	e.Add(http.MethodGet, "/infringement_monitor/get", controller.InfringementMonitorGet)
	e.Add(http.MethodPost, "/infringement_monitor/update_base_info", controller.InfringementMonitorUpdateBaseInfo)
	//
	bg.GET("/infringement_monitor/get", controller.InfringementMonitorBgGet)
	bg.POST("/infringement_monitor/set_deal_info", controller.InfringementMonitorSetDealInfo)
	bg.GET("/infringement_monitor/list", controller.InfringementMonitorBackendList)
	///
	///我要维权
	e.Add(http.MethodPost, "/rights_protection/add", controller.RightsProtectionAdd)
	e.Add(http.MethodGet, "/rights_protection/get", controller.RightsProtectionGet)
	e.Add(http.MethodPost, "/rights_protection/update_base_info", controller.RightsProtectionUpdateBaseInfo)
	//
	bg.GET("/rights_protection/get", controller.RightsProtectionBgGet)
	bg.POST("/rights_protection/set_deal_info", controller.RightsProtectionSetDealInfo)
	bg.GET("/rights_protection/list", controller.RightsProtectionBackendList)
	///
	///用户
	e.Add(http.MethodPost, "/user/login", controller.Login)
	e.Add(http.MethodPost, "/user/set_phone", controller.SetPhone)
	e.Add(http.MethodPost, "/user/set_name_and_avatar_url", controller.SetNameAndAvatarUrl)
	e.Add(http.MethodGet, "/user/get_user_info", controller.GetUserInfo)
	e.Add(http.MethodGet, "/wx/notify", controller.WxNotify)
	//
	bg.GET("/user/get_user_info", controller.GetUserInfo)
	bg.POST("/user/login", controller.BackendLogin)
	///
}
