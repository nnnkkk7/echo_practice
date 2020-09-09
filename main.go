package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type Response struct {
	Name    string `json:"name"`
	Message string `json:message`
	Status  string `json:status`
}

func main() {
	e := echo.New()

	e.POST("/send", sendMessage)
	e.Logger.Fatal(e.Start(":1323"))
}

func sendMessage(c echo.Context) error {
	m := new(Message)
	if error := c.Bind(m); error != nil {
		return error
	}
	r := new(Response)
	r.Name = m.Name
	r.Message = m.Message
	r.Status = "success"
	return c.JSON(http.StatusOK, r)
}
