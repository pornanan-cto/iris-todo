package controller

import (
	"fmt"
	"iris-todos/repository"
	util "iris-todos/utils"

	"github.com/kataras/iris/v12"
)

func GetTodos(ctx iris.Context) {
	claims := util.GetClaims(ctx)
	status := ctx.URLParamDefault("status", "all")
	userTodos := repository.GetRelatedTodos(claims.ID, status)

	ctx.StatusCode(iris.StatusOK)
	util.Response(userTodos, ctx)
}

func CreateTodo(ctx iris.Context) {
	claims := util.GetClaims(ctx)

	taskName := ctx.PostValue("taskName")
	taskDescription := ctx.PostValue("taskDescription")

	todo, err := repository.NewTodo(claims.ID, taskName, taskDescription)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		util.Response(err.Error(), ctx)
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	util.Response(todo, ctx)
}

func AddAssignee(ctx iris.Context) {
	claims := util.GetClaims(ctx)

	taskID, _ := ctx.Params().GetUint("taskID")
	assigneeID, _ := ctx.PostValueInt("assigneeId")

	if !repository.IsOwner(claims.ID, taskID) {
		ctx.StatusCode(iris.StatusBadRequest)
		util.Response("only task owner to manage this task.", ctx)
		return
	}

	updatedTodo := repository.AddAssignee(taskID, uint(assigneeID))
	ctx.StatusCode(iris.StatusOK)
	util.Response(updatedTodo, ctx)
}

func RemoveTodo(ctx iris.Context) {
	claims := util.GetClaims(ctx)
	taskID, _ := ctx.Params().GetUint("taskID")

	if !repository.IsOwner(claims.ID, taskID) {
		ctx.StatusCode(iris.StatusBadRequest)
		util.Response("only task owner to manage this task.", ctx)
		return
	}

	repository.RemoveTodo(taskID)
	ctx.StatusCode(iris.StatusOK)
	util.Response(nil, ctx)
}

func ChangeTodoStatus(ctx iris.Context) {
	claims := util.GetClaims(ctx)

	taskID, _ := ctx.Params().GetUint("taskID")

	if !repository.IsRelevant(claims.ID, taskID) {
		ctx.StatusCode(iris.StatusBadRequest)
		util.Response("only task owner, assignee is allowed to manage this task.", ctx)
		return
	}

	status := ctx.PostValueTrim("status")

	updatedTodo := repository.ChangeTodoStatus(taskID, status)
	ctx.StatusCode(iris.StatusOK)
	util.Response(updatedTodo, ctx)
}

func RemoveAssignee(ctx iris.Context) {
	claims := util.GetClaims(ctx)

	taskID, _ := ctx.Params().GetUint("taskID")

	if !repository.IsOwner(claims.ID, taskID) {
		ctx.StatusCode(iris.StatusBadRequest)
		util.Response("only task owner to manage this task.", ctx)
		return
	}

	updatedTodo := repository.AddAssignee(taskID, uint(0))

	fmt.Println("Push notification!")

	ctx.StatusCode(iris.StatusOK)
	util.Response(updatedTodo, ctx)
}
