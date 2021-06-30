package routing

import (
	"fmt"

	"github.com/iamthe1whoknocks/bazilikgroup_test_task/pkg/handler"
	"github.com/labstack/echo/v4"
)

//Run - init and run web-server
func Run(host, port string) error {
	e := echo.New()

	e.HideBanner = true
	e.Debug = true

	handler := handler.New()

	api := e.Group("/api")
	api.GET("/add", handler.Add)
	api.GET("/sub", handler.Sub)
	api.GET("/mul", handler.Mul)
	api.GET("/div", handler.Div)

	return e.Start(fmt.Sprintf("%s:%s", host, port))
}
