package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductHandler interface {
	GetAll(c *gin.Context)
	FindById(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type productHandler struct {
}

func NewProductController(logger *zap.Logger) ProductHandler {
	return &productHandler{}
}

func (pc *productHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ProductHandler": "GetAll"})
}
func (pc *productHandler) FindById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ProductHandler": "FindById"})

}

func (pc *productHandler) Add(c *gin.Context) {

}

func (pc *productHandler) Update(c *gin.Context) {

}
func (pc *productHandler) Delete(c *gin.Context) {

}

func ProductRoute(pc ProductHandler, logger *zap.Logger, c *gin.Engine) {
	api := c.Group("/api")

	api.GET("/products", pc.GetAll)
	api.GET("/products/{id}", pc.FindById)
	api.POST("/products", pc.Add)
	api.PUT("/products", pc.Update)
	api.DELETE("/products/{id}", pc.Delete)
}
