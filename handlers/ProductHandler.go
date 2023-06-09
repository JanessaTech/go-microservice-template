package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hi-supergirl/go-microservice-template/handlers/services"
	"github.com/hi-supergirl/go-microservice-template/handlers/services/dto"
	"github.com/hi-supergirl/go-microservice-template/logging"
	"github.com/hi-supergirl/go-microservice-template/middlewares"
	"go.uber.org/zap"
)

type ProductHandler interface {
	GetAll(c *gin.Context)
	Add(c *gin.Context)
	Delete(c *gin.Context)
}

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(logger *zap.Logger, productService services.ProductService) ProductHandler {
	return &productHandler{productService: productService}
}

func (ph *productHandler) GetAll(c *gin.Context) {
	execute(c, func(c *gin.Context) *Response {
		logger := logging.FromContext(c)
		logger.Infow("[productHandler]", "GetAll", "")
		products, err := ph.productService.GetAll(c.Request.Context())
		if err != nil {
			return NewResponse(http.StatusBadRequest, failed, err)
		}
		return NewResponse(http.StatusOK, success, products)
	})
}

func (ph *productHandler) Add(c *gin.Context) {
	execute(c, func(c *gin.Context) *Response {
		logger := logging.FromContext(c)
		logger.Infow("[productHandler]", "Add", "")
		var productDtO dto.ProductDTO
		if err := c.ShouldBindJSON(&productDtO); err != nil {
			return NewResponse(http.StatusBadRequest, failed, err)
		}
		savedProduct, err := ph.productService.Add(c.Request.Context(), productDtO)
		if err != nil {
			return NewResponse(http.StatusBadRequest, failed, err)
		}

		return NewResponse(http.StatusOK, success, savedProduct)
	})
}

func (ph *productHandler) Delete(c *gin.Context) {
	execute(c, func(c *gin.Context) *Response {
		logger := logging.FromContext(c)
		id := StringToUint(c.Param("id"))
		logger.Infow("[productHandler]", "Delete id = ", id)
		err := ph.productService.Delete(c.Request.Context(), id)
		if err != nil {
			return NewResponse(http.StatusBadRequest, failed, err)
		}
		return NewResponse(http.StatusOK, success, nil)
	})

}

func StringToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}

func ProductRoute(ph ProductHandler, logger *zap.Logger, c *gin.Engine) {
	api := c.Group("/api")
	api.Use(middlewares.JwtTokenMiddleWare(), middlewares.RequestTraceMiddleWare())
	api.GET("/products", ph.GetAll)
	api.POST("/products", ph.Add)
	api.DELETE("/products/:id", ph.Delete)
}
