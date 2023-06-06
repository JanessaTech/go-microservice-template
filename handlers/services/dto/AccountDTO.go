package dto

import "errors"

type AccountDTO struct {
	ID       int    `json:"id"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (accountDTO *AccountDTO) ValidatePassword(password string) error {
	if accountDTO.Password != password {
		return errors.New("password is not matched")
	}
	return nil
}
