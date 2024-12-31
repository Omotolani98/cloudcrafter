package routes

import (
	"cloudcrafter/pkg/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all the routes for the application
func RegisterRoutes(router *gin.Engine, handler *handlers.ProvisioningHandler) {
	RegisterHealthRoutes(router)
	RegisterProvisioningRoutes(router, handler)
}
