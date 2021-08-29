package model

import "time"

type Todo struct {
	ID              uint      `json:"id"`
	TaskOwnerID     uint      `gorm:"index" json:"taskOwnerId"`
	TaskAssigneeID  uint      `gorm:"index" json:"taskAssignerId"`
	TaskName        string    `json:"taskName"`
	TaskDescription string    `json:"taskDescription"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
