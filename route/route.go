package route

import (
	"law/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRouter(r *echo.Router) {
	r.Add(http.MethodPost, "/user/login", controller.UserLogin)
	
}
