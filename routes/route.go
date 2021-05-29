package route

import (
	handler "m9hub/controllers"

	"github.com/labstack/echo/v4"
)

func RouterSetup() {
	e := echo.New()
	userHandler := handler.TestHandler{}
	e.GET("/", handler.TestGetData)
	e.GET("/test-struct", userHandler.TestGetDataInStruct)

	events := e.Group("/events")
	events.GET("", handler.GetAllEventsHandler)

	e.Logger.Fatal(e.Start(":3000"))
}
