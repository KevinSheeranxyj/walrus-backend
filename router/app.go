package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"kevinsheeran/walrus-backend/api"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/register", api.CreateForm)
	return router
}

func register(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
