package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()

	InitializeRoutes(&r.RouterGroup)

	r.Run(":8080")
}
