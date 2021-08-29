package controller

import (
	"iris-todos/repository"
	util "iris-todos/utils"

	"github.com/kataras/iris/v12"
)

func CreateUser(ctx iris.Context) {
	username := ctx.PostValue("username")
	plainPassword := ctx.PostValue("password")

	hashedPassword := util.HashPassword(plainPassword)

	if repository.IsUserExisted(username) {
		ctx.StatusCode(iris.StatusBadRequest)
		util.Response("user is already existed.", ctx)
		return
	}

	user, err := repository.CreateUser(username, hashedPassword)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		util.Response(err.Error(), ctx)
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	util.Response(user, ctx)
}

func Login(ctx iris.Context) {
	username := ctx.PostValue("username")
	plainPassword := ctx.PostValue("password")

	user := repository.FindByUsername(username)

	if !util.ComparePassword(user.Password, plainPassword) {
		ctx.StatusCode(iris.StatusUnauthorized)
		util.Response("password is incorrect.", ctx)
		return
	}

	token, err := util.GenerateJWT(user)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		util.Response(err.Error(), ctx)
	}

	ctx.StatusCode(iris.StatusOK)
	util.Response(token, ctx)
}

func Me(ctx iris.Context) {
	claims := util.GetClaims(ctx)

	user := repository.FindByUsername(claims.Username)

	ctx.StatusCode(iris.StatusOK)
	util.Response(user, ctx)
}
