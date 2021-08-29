package model

import "time"

type User struct {
	ID        uint      `json:"id"`
	Username  string    `gorm:"index" json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
