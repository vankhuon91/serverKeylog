package main

import "github.com/kataras/iris/v12"

type (
	request struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}

	response struct {
		ID      uint64 `json:"id"`
		Message string `json:"message"`
	}
)

func main() {
	app := iris.New()
	app.Handle("GET", "/users", updateUser)
	app.Listen(":8080")
}

func updateUser(ctx iris.Context) {
	id, _ := ctx.Params().GetUint64("id")

	var req request

	resp := response{
		ID:      id,
		Message: req.Firstname + " updated successfully",
	}
	ctx.JSON(resp)
}
