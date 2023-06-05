package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct{}

func NewProductController() *ProductHandler {
	return &ProductHandler{}
}

func (pc *ProductHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ProductController": "GetAll"})
}
func (pc *ProductHandler) FindById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ProductController": "FindById"})

}

func (pc *ProductHandler) Add(c *gin.Context) {

}

func (pc *ProductHandler) Update(c *gin.Context) {

}
func (pc *ProductHandler) Delete(c *gin.Context) {

}

func ProductRoute(pc *ProductHandler, c *gin.Engine) {
	api := c.Group("/api")

	api.GET("/products", pc.GetAll)
	api.GET("/products/{id}", pc.FindById)
	api.POST("/products", pc.Add)
	api.PUT("/products", pc.Update)
	api.DELETE("/products/{id}", pc.Delete)
}
