package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"kevinsheeran/walrus-backend/api"
	_ "kevinsheeran/walrus-backend/docs"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	gin.SetMode(gin.DebugMode)
	router.Use(gin.Recovery())
	register(router)
	return router
}

func register(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/api/v1/create-form", api.CreateForm)
	router.GET("/api/v1/form/:blobId", api.GetForm)
}
