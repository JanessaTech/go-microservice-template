package dto

type AccountDTO struct {
	ID       string `json:"id"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
