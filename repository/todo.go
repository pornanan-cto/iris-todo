package repository

import (
	model "iris-todos/models"
	util "iris-todos/utils"
)

const (
	IN_PROGRESS = "in_progress"
	NOT_STARTED = "not_started"
	COMPLETED   = "complted"
	ALL         = "all"
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
		Status:          NOT_STARTED,
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

func GetRelatedTodos(userID uint, status string) UserTodos {
	db := util.DatabaseConnect()
	userTodos := UserTodos{}

	ownedTasks := db.Where("task_owner_id = ?", userID)
	if status != ALL {
		ownedTasks.Where("status = ?", status)
	}
	ownedTasks.Find(&userTodos.OwnedTasks)

	assignedTasks := db.Where("task_assignee_id = ?", userID).Find(&userTodos.AssignedTasks)
	if status != ALL {
		assignedTasks.Where("status = ?", status)
	}
	assignedTasks.Find(&userTodos.AssignedTasks)

	return userTodos
}

func IsOwner(userID uint, taskID uint) bool {
	todo := FindTaskByID(taskID)

	return todo.TaskOwnerID == userID
}

func RemoveTodo(taskID uint) {
	db := util.DatabaseConnect()
	todo := FindTaskByID(taskID)
	db.Delete(todo)
}
