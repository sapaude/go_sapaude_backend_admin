package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sapaude/go-shims/shim"
)

func logRouter(e *echo.Echo) error {
	// ---------------------------------------------------
	// 核心：遍历并打印路由信息
	// ---------------------------------------------------
	fmt.Println("---------------------------------------------------")
	fmt.Println("Registered Echo Routes:")
	fmt.Println("---------------------------------------------------")

	for i, route := range e.Routes() {
		fmt.Printf("route[%d] => %s\n", i, shim.ToJsonString(route, false))
	}
	fmt.Println("---------------------------------------------------")

	return nil
}
