package route

import (
	"law/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRouter(r *echo.Router) {
	///法律问题
	r.Add(http.MethodGet, "/legal_issue/list", controller.LegalIssueList)
	r.Add(http.MethodGet, "/legal_issue/category_list", controller.LegalIssueCategoryList)
	r.Add(http.MethodGet, "/legal_issue/get", controller.LegalIssueGet)
	r.Add(http.MethodPost, "/legal_issue/favorite", controller.LegalIssueFavorite)
	r.Add(http.MethodPost, "/legal_issue/cancel_favorite", controller.LegalIssueCancelFavorite)
	r.Add(http.MethodGet, "/legal_issue/is_favorite", controller.LegalIssueIsFavorite)
	///
	///咨询
	r.Add(http.MethodPost, "/consultation/create", controller.ConsultationCreate)
	r.Add(http.MethodGet, "/consultation/get", controller.ConsultationGet)
	r.Add(http.MethodGet, "/consultation/list", controller.ConsultationList)
	r.Add(http.MethodPost, "/consultation/set_status", controller.ConsultationSetStatus)
	r.Add(http.MethodPost, "/consultation/add_reply", controller.ConsultationAddReply)
	r.Add(http.MethodGet, "/consultation/list_reply", controller.ConsultationListReply)
	r.Add(http.MethodGet, "/consultation/file_upload_auth", controller.ConsultationFileUploadAuth)
	r.Add(http.MethodPost, "/backend/consultation/set_status", controller.ConsultationSetStatus)
	r.Add(http.MethodPost, "/backend/consultation/add_reply", controller.ConsultationAddReply)
	r.Add(http.MethodGet, "/backend/consultation/list_reply", controller.ConsultationListReply)
	r.Add(http.MethodGet, "/backend/consultation/get", controller.ConsultationGet)
	r.Add(http.MethodGet, "/backend/consultation/file_upload_auth", controller.ConsultationBackendFileUploadAuth)
	r.Add(http.MethodGet, "/backend/consultation/list", controller.ConsultationBackendList)
	///
	///侵权监测
	r.Add(http.MethodPost, "/infringement_monitor/add", controller.InfringementMonitorAdd)
	r.Add(http.MethodGet, "/infringement_monitor/get", controller.InfringementMonitorGet)
	r.Add(http.MethodPost, "/infringement_monitor/update_base_info", controller.InfringementMonitorUpdateBaseInfo)
	r.Add(http.MethodPost, "/backend/infringement_monitor/get", controller.InfringementMonitorBgGet)
	r.Add(http.MethodPost, "/backend/infringement_monitor/set_deal_info", controller.InfringementMonitorSetDealInfo)
	r.Add(http.MethodGet, "/backend/infringement_monitor/backend_list", controller.InfringementMonitorBackendList)
	///
	///我要维权
	r.Add(http.MethodPost, "/rights_protection/add", controller.RightsProtectionAdd)
	r.Add(http.MethodGet, "/rights_protection/get", controller.RightsProtectionGet)
	r.Add(http.MethodPost, "/rights_protection/update_base_info", controller.RightsProtectionUpdateBaseInfo)
	r.Add(http.MethodGet, "/backend/rights_protection/get", controller.RightsProtectionBgGet)
	r.Add(http.MethodPost, "/backend/rights_protection/set_deal_info", controller.RightsProtectionSetDealInfo)
	r.Add(http.MethodGet, "/backend/rights_protection/list", controller.RightsProtectionBackendList)
	///
	///用户
	r.Add(http.MethodPost, "/user/login", controller.Login)
	r.Add(http.MethodPost, "/user/set_phone", controller.SetPhone)
	r.Add(http.MethodPost, "/user/set_name_and_avatar_url", controller.SetNameAndAvatarUrl)
	r.Add(http.MethodGet, "/user/get_user_info", controller.GetUserInfo)
	r.Add(http.MethodGet, "/backend/user/get_user_info", controller.GetUserInfo)
	r.Add(http.MethodPost, "/backend/user/login", controller.BackendLogin)
	///
}
