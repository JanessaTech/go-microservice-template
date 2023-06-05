package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hi-supergirl/go-microservice-template/middlewares"
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
	c.JSON(http.StatusOK, gin.H{"AccountHandler": "Logout"})
}
func (h *AccountHandler) Me(c *gin.Context) {

}

func Route1(h *AccountHandler, c *gin.Engine) {
	api := c.Group("/api")

	api.Use()
	{
		api.POST("/account/register", h.Register)
		api.POST("/account/login", h.Login)
	}

	api.Use(middlewares.VerifyJwtToken())
	{
		api.GET("/account/logout", h.Logout)
		api.GET("/account/me", h.Me)
	}

}
