package route

import (
	"law/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRouter(r *echo.Router) {
	r.Add(http.MethodGet, "/legal_issue/list", controller.LegalIssueList)
	r.Add(http.MethodGet, "/legal_issue/category_list", controller.LegalIssueCategoryList)
	r.Add(http.MethodGet, "/legal_issue/get", controller.LegalIssueGet)
	r.Add(http.MethodPost, "/legal_issue/favorite", controller.LegalIssueFavorite)
	r.Add(http.MethodPost, "/legal_issue/cancel_favorite", controller.LegalIssueCancelFavorite)
	r.Add(http.MethodGet, "/legal_issue/is_favorite", controller.LegalIssueIsFavorite)
	r.Add(http.MethodPost, "/consultation/create", controller.ConsultationCreate)
	r.Add(http.MethodGet, "/consultation/get", controller.ConsultationGet)
	r.Add(http.MethodGet, "/consultation/list", controller.ConsultationList)
	r.Add(http.MethodPost, "/consultation/set_status", controller.ConsultationSetStatus)
	r.Add(http.MethodPost, "/consultation/add_reply", controller.ConsultationAddReply)
	r.Add(http.MethodGet, "/consultation/list_reply", controller.ConsultationListReply)
	r.Add(http.MethodGet, "/consultation/backend_list", controller.ConsultationBackendList)
	r.Add(http.MethodGet, "/consultation/file_upload_auth", controller.ConsultationFileUploadAuth)
	r.Add(http.MethodGet, "/consultation/reply_file_upload_auth", controller.ConsultationReplyFileUploadAuth)
	r.Add(http.MethodPost, "/infringement_detection/add", controller.InfringementDetectionAdd)
	r.Add(http.MethodPost, "/rights_protection/add", controller.RightsProtectionAdd)
	r.Add(http.MethodPost, "/detection_return_visit/update", controller.DetectionReturnVisitUpdate)
	r.Add(http.MethodGet, "/detection_return_visit/get", controller.DetectionReturnVisitGet)
	r.Add(http.MethodGet, "/infringement_detection/get", controller.InfringementDetectionGet)
	r.Add(http.MethodGet, "/infringement_detection/list", controller.InfringementDetectionList)
	r.Add(http.MethodPost, "/protection_return_visit/update", controller.ProtectionReturnVisitUpdate)
	r.Add(http.MethodGet, "/protection_return_visit/get", controller.ProtectionReturnVisitGet)
	r.Add(http.MethodGet, "/rights_protection/get", controller.RightsProtectionGet)
	r.Add(http.MethodGet, "/rights_protection/list", controller.RightsProtectionList)
	r.Add(http.MethodPost, "/user/login", controller.Login)
	r.Add(http.MethodPost, "/background/login", controller.BackgroundLogin)
	r.Add(http.MethodPost, "/user/set_phone", controller.SetPhone)
	r.Add(http.MethodPost, "/user/set_name_and_avatar_url", controller.SetNameAndAvatarUrl)
	r.Add(http.MethodGet, "/user/get_user_info", controller.GetUserInfo)
}
