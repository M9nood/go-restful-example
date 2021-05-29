package controllers

import (
	"net/http"

	"m9hub/services"

	"github.com/labstack/echo/v4"
)

func GetAllEventsHandler(c echo.Context) error {
	eventService := services.NewEventService()
	events, err := eventService.GetAllEventsService()
	if err != nil {
		return c.JSON(http.StatusOK, err)
	}
	return c.JSON(http.StatusOK, events)
}
