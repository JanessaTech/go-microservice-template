package dto

import "time"

type ProductDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
