package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TestHandler struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// var testInstance *TestHandler

func TestGetData(c echo.Context) error {
	return c.String(http.StatusOK, "dsffsd")
}

func (h *TestHandler) TestGetDataInStruct(c echo.Context) error {
	user := Init("1", "M9nood")
	fmt.Println("user", user)
	return c.JSON(http.StatusOK, user)
}

func Init(
	id string,
	name string) TestHandler {
	testInstance := TestHandler{
		Id:   id,
		Name: name,
	}
	return testInstance
}
