package main

import (
	"os"

	"github.com/kataras/iris/v12"
)

type (
	request struct {
		MAC     string `json:"MAC"`
		IP      string `json:"IP"`
		Message string `json:"Message"`
	}
)

func main() {
	app := iris.New()
	app.Handle("POST", "/api/computers", addNewComputer)
	var port_number = os.Getenv("PORT")
	if port_number == "" {
		port_number = "8080"
	}
	app.Listen(":" + port_number)
}

func addNewComputer(ctx iris.Context) {
	var req request
	ctx.ReadJSON(&req)

	ctx.JSON(req)
}
