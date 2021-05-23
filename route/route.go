package route

import (
	"gorest/controllers"

	"github.com/labstack/echo/v4"
)

func RouterSetup() {
	e := echo.New()
	userHandler := controllers.TestHandler{}
	e.GET("/", controllers.TestGetData)
	e.GET("/test-struct", userHandler.TestGetDataInStruct)
	e.Logger.Fatal(e.Start(":1323"))
}
