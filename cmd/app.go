package cmd

import (
	"context"
	"law/conf"
	"law/model"
	"law/route"
	"law/service"
	"law/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() {
	initTimezone()
	c := make(chan os.Signal, 1)
	conf.Init()
	service.InitCacheManager()
	utils.InitLogger()
	defer utils.FlushLog()
	model.Init()
	//以上调用顺序不应随意改变,彼此存在依赖关系
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	format := `remote_ip:${remote_ip},host:${host},method:${method},user_agent:${user_agent},status:${status},error:${error},latency_human:${latency_human},uri:${uri}`
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: log.Logger, Format: format}))
	e.Use(utils.MidAuth)
	e.HTTPErrorHandler = customHTTPErrorHandler
	route.InitRouter(e)
	go func() {
		if err := e.Start(conf.App.Http.Address); err != nil {
			println(err.Error())
		}
	}()
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	switch s {
	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if e != nil {
			_ = e.Shutdown(ctx)
		}
	case syscall.SIGHUP:
	default:
	}
}

type CustomValidator struct {
	validator *validator.Validate
}

func customHTTPErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}
	message := he.Message
	if !c.Response().Committed {
		c.JSON(utils.Fail("系统错误", message))
	}
}
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func initTimezone() {
	loc, _ := time.LoadLocation("Asia/Shanghai") //加载时区
	time.Local = loc
}
