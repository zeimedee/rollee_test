package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zeimedee/test2/internal/handlers"
	"github.com/zeimedee/test2/internal/services"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	service := services.NewWordService()
	wordhandler := handlers.NewWordHandler(service)

	api := router.Group("/service")
	{
		api.POST("/", wordhandler.StoreWord)
		api.GET("/", wordhandler.RetrieveWord)
	}
	return router
}
