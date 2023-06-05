package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ProductHandler": "GetAll"})
}
func (h *ProductHandler) FindById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ProductHandler": "FindById"})

}

func (h *ProductHandler) Add(c *gin.Context) {

}

func (h *ProductHandler) Update(c *gin.Context) {

}
func (h *ProductHandler) Delete(c *gin.Context) {

}

func Route2(h *ProductHandler, c *gin.Engine) {
	api := c.Group("/api")

	api.GET("/products", h.GetAll)
	api.GET("/products/{id}", h.FindById)
	api.POST("/products", h.Add)
	api.PUT("/products", h.Update)
	api.DELETE("/products/{id}", h.Delete)
}
