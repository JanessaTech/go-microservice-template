package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hi-supergirl/go-microservice-template/middlewares"
)

type AccountHandler struct{}

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{}
}

func (ac *AccountHandler) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"AccountHandler": "Register"})
}

func (ac *AccountHandler) Login(c *gin.Context) {

}

func (ac *AccountHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"AccountHandler": "Logout"})
}
func (ac *AccountHandler) Me(c *gin.Context) {

}

func AccountRoute(ac *AccountHandler, c *gin.Engine) {
	api := c.Group("/api")

	api.Use()
	{
		api.POST("/account/register", ac.Register)
		api.POST("/account/login", ac.Login)
	}

	api.Use(middlewares.VerifyJwtToken())
	{
		api.GET("/account/logout", ac.Logout)
		api.GET("/account/me", ac.Me)
	}

}
