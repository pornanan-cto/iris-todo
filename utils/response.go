package util

import "github.com/kataras/iris/v12"

func Response(data interface{}, ctx iris.Context) {
	if ctx.GetStatusCode() >= 400 {
		response := iris.Map{
			"message": data,
			"status":  ctx.GetStatusCode(),
		}
		ctx.JSON(response)
	} else {
		response := iris.Map{
			"result": data,
			"status": ctx.GetStatusCode(),
		}
		ctx.JSON(response)
	}
}
