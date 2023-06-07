package model

import "time"

type Account struct {
	ID        uint      `gorm:"column:id"`
	UserName  string    `gorm:"column:user_name;default:JanessaTech"`
	Password  string    `gorm:"column:password;"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
