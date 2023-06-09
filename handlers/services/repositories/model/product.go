package model

import "time"

type Product struct {
	ID        uint      `gorm:"column:id"`
	Name      string    `gorm:"column:name;"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
