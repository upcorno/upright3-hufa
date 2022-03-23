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
}
