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
	r.Add(http.MethodPost, "/user/get_user_info", controller.GetUserInfo)
}
