package user

import "time"

type User struct {
	ID             int `gorm:"primaryKey"`
	UserName       string
	Email          string
	Password       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
