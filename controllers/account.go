package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct{}

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{}
}

func (h *AccountHandler) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"AccountHandler": "Register"})
}

func (h *AccountHandler) Login(c *gin.Context) {

}

func (h *AccountHandler) Logout(c *gin.Context) {

}
func (h *AccountHandler) Me(c *gin.Context) {

}

func Route1(h *AccountHandler, c *gin.Engine) {
	api := c.Group("/api")

	api.POST("/account/register", h.Register)
	api.POST("/account/login", h.Login)
	api.GET("/account/logout", h.Logout)
	api.GET("/account/me", h.Me)
}
