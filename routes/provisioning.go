package routes

import (
	"cloudcrafter/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterProvisioningRoutes(router *gin.Engine, handler *handlers.ProvisioningHandler) {
	router.POST("/resources", handler.CreateResourceHandler)
	router.DELETE("/resources/:provider/:id", handler.DeleteResourceHandler)
	router.GET("/resources/:provider/:id", handler.GetResourceHandler)
}
