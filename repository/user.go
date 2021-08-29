package repository

import (
	model "iris-todos/models"
	util "iris-todos/utils"
)

func IsUserExisted(username string) bool {
	db := util.DatabaseConnect()

	results := db.Where("username = ?", username).First(&model.User{})

	return results.RowsAffected > 0
}

func FindByUsername(username string) model.User {
	db := util.DatabaseConnect()
	user := model.User{}
	db.Where("username = ?", username).First(&user)
	return user
}

func CreateUser(username string, hashedPassword string) (model.User, error) {
	db := util.DatabaseConnect()

	user := model.User{
		Username: username,
		Password: hashedPassword,
	}

	result := db.Create(&user)

	if result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}
