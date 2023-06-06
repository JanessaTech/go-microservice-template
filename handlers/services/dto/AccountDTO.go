package dto

import "golang.org/x/crypto/bcrypt"

type AccountDTO struct {
	ID       int    `json:"id"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (accountDTO *AccountDTO) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(accountDTO.Password), []byte(password))
}
