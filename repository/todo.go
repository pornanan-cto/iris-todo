package repository

import (
	model "iris-todos/models"
	util "iris-todos/utils"
)

func FindTaskByID(taskID uint) model.Todo {
	db := util.DatabaseConnect()
	todo := model.Todo{}
	db.Where("id = ?", taskID).First(&todo)
	return todo
}

func NewTodo(ownerID uint, taskName string, taskDescription string) (model.Todo, error) {
	db := util.DatabaseConnect()

	todo := model.Todo{
		TaskOwnerID:     ownerID,
		TaskName:        taskName,
		TaskDescription: taskDescription,
	}

	result := db.Create(&todo)

	if result.Error != nil {
		return model.Todo{}, result.Error
	}

	return todo, nil
}

func AddAssignee(taskID uint, assigneeID uint) model.Todo {
	db := util.DatabaseConnect()
	todo := FindTaskByID(taskID)
	todo.TaskAssigneeID = assigneeID
	db.Save(todo)
	return todo
}

type UserTodos struct {
	OwnedTasks    []model.Todo `json:"ownedTasks"`
	AssignedTasks []model.Todo `json:"assignedTasks"`
}

func GetRelatedTodos(userID uint) UserTodos {
	db := util.DatabaseConnect()
	userTodos := UserTodos{}

	db.Where("task_owner_id = ?", userID).Find(&userTodos.OwnedTasks)
	db.Where("task_assignee_id = ?", userID).Find(&userTodos.AssignedTasks)

	return userTodos
}
