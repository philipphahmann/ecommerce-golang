package router

import (
	"ecommerce/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.RouterGroup) {
	basePath := "/v1/product/"

	v1 := r.Group(basePath)
	{
		v1.GET("get/:id", handlers.GetProduct)
		v1.POST("create/", handlers.CreateProduct)
		v1.PATCH("update/:id", handlers.UpdateProduct)
		v1.DELETE("delete/:id", handlers.DeleteProduct)
	}
}
