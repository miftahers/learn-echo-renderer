package model

import "time"

type User struct {
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" form:"deletedAt"`
	UUID      string    `json:"uuid" form:"uuid" gorm:"primaryKey"`
	Username  string    `json:"username" form:"username"`
	Password  string    `json:"password" form:"password"`
}
