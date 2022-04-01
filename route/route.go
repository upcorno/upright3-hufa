package route

import (
	"law/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRouter(r *echo.Router) {
	r.Add(http.MethodPost, "/consultation/create", controller.ConsultationCreate)
	r.Add(http.MethodPost, "/consultation_record/create", controller.ConsultationRecordCreate)
	r.Add(http.MethodGet, "/consultation_record_list/get", controller.ConsultationRecordListGet)
	r.Add(http.MethodGet, "/legal_issue/list", controller.LegalIssueList)
	r.Add(http.MethodPost, "/user/login", controller.Login)
	r.Add(http.MethodPost, "/user/set_phone", controller.SetPhone)
	r.Add(http.MethodPost, "/user/set_name_and_avatar_url", controller.SetNameAndAvatarUrl)
	r.Add(http.MethodGet, "/user/get_user_info", controller.GetUserInfo)
	r.Add(http.MethodPost, "/consultation_status/set", controller.ConsultationStatusSet)
	r.Add(http.MethodGet, "/legal_issue/get", controller.LegalIssueGet)
	r.Add(http.MethodPost, "/favorite/add", controller.FavoriteAdd)
	r.Add(http.MethodPost, "/favorite/cancel", controller.FavoriteCancel)
	r.Add(http.MethodGet, "/favorite/get", controller.IssueIsFavorite)
	r.Add(http.MethodGet, "/favorite/list", controller.FavoriteList)
	r.Add(http.MethodGet, "/consultation/list", controller.ConsultationList)
	r.Add(http.MethodGet, "/consultation/file_upload_auth", controller.ConsultationFileUploadAuth)
	r.Add(http.MethodGet, "/consultation_record/file_upload_auth", controller.ConsultationRecordFileUploadAuth)
	r.Add(http.MethodGet, "/category_list/get", controller.CategoryListGet)
	r.Add(http.MethodGet, "/legal_issue_list/by_category", controller.LegalIssueListByCategory)
	r.Add(http.MethodGet, "/consultation/get", controller.ConsultationGet)
	r.Add(http.MethodPost, "/infringement_detection/add", controller.InfringementDetectionAdd)
	r.Add(http.MethodPost, "/rights_protection/add", controller.RightsProtectionAdd)
	r.Add(http.MethodPost, "/detection_return_visit/update", controller.DetectionReturnVisitUpdate)
	r.Add(http.MethodGet, "/detection_return_visit/get", controller.DetectionReturnVisitGet)
	r.Add(http.MethodGet, "/infringement_detection/get", controller.InfringementDetectionGet)
	r.Add(http.MethodGet, "/infringement_detection/list", controller.InfringementDetectionList)
	r.Add(http.MethodPost, "/protection_return_visit/update", controller.ProtectionReturnVisitUpdate)
	r.Add(http.MethodGet, "/protection_return_visit/get", controller.ProtectionReturnVisitGet)
}
