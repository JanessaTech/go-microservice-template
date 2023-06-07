package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hi-supergirl/go-microservice-template/handlers/services"
	"github.com/hi-supergirl/go-microservice-template/handlers/services/dto"
	"github.com/hi-supergirl/go-microservice-template/helper"
	"github.com/hi-supergirl/go-microservice-template/middlewares"
)

type AccountHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Me(c *gin.Context)
}

type accountHandler struct {
	accountService services.AccountService
	nextId         int
}

func NewAccountHandler(accountService services.AccountService) AccountHandler {
	return &accountHandler{accountService: accountService}
}

func (ah *accountHandler) Register(c *gin.Context) {
	var auth dto.AccountDTO
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encodedPassword, err := helper.EncodePassword(auth.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accDto := dto.AccountDTO{ID: ah.nextId, UserName: auth.UserName, Password: encodedPassword}
	ah.nextId = ah.nextId + 1
	savedAccDto, err := ah.accountService.Save(c.Request.Context(), accDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	savedAccDto.Password = ""

	c.JSON(http.StatusOK, gin.H{"savedAccDto": savedAccDto})
}

func (ah *accountHandler) Login(c *gin.Context) {
	var auth dto.AccountDTO
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accDto, err := ah.accountService.GetByName(c.Request.Context(), auth.UserName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := accDto.ValidatePassword(auth.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	jwt, err := helper.GenerateJWT(accDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

func (ah *accountHandler) Me(c *gin.Context) {
	curAccount, err := ah.getCurrentAccount(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	curAccount.Password = ""
	c.JSON(http.StatusOK, gin.H{"current account": curAccount})
}

func (ah *accountHandler) getCurrentAccount(c *gin.Context) (*dto.AccountDTO, error) {
	currentId, err := helper.GetCurrentAccountId(c)
	if err != nil {
		return nil, err
	}
	accDto, err := ah.accountService.GetById(c.Request.Context(), currentId)
	if err != nil {
		return nil, err
	}
	return accDto, nil
}

func AccountRoute(ah AccountHandler, c *gin.Engine) {
	api := c.Group("/api")

	api.Use()
	{
		api.POST("/account/register", ah.Register)
		api.POST("/account/login", ah.Login)
	}

	api.Use(middlewares.VerifyJwtToken())
	{
		api.GET("/account/me", ah.Me)
	}
}
