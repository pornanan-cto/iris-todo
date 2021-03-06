package main

import (
	controller "iris-todos/controllers"
	util "iris-todos/utils"
	"log"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	config, err := util.LoadConfigs(".")
	if err != nil {
		log.Fatal("Cannot load config.", err)
	}

	util.DatabaseConnect()

	userAPI := app.Party("/auth")
	{
		userAPI.Post("", controller.CreateUser)
		userAPI.Post("/login", controller.Login)
	}

	protectedAPI := app.Party("/")
	{
		protectedAPI.Use(util.VerifyTokenHandler())
		protectedAPI.Get("/me", controller.Me)

		protectedAPI.Get("/todos", controller.GetTodos)
		protectedAPI.Post("/todos", controller.CreateTodo)
		protectedAPI.Delete("/todos/{taskID:uint}", controller.RemoveTodo)
		protectedAPI.Put("/todos/{taskID:uint}/assignee", controller.AddAssignee)
		protectedAPI.Delete("/todos/{taskID:uint}/assignee", controller.RemoveAssignee)
		protectedAPI.Put("/todos/{taskID:uint}/status", controller.ChangeTodoStatus)
	}

	app.Listen(":" + config.APP_PORT)
}
