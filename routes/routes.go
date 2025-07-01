package routes

import (
	"cloudcrafter/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handler *handlers.ProvisioningHandler) {
	RegisterHealthRoutes(router)
	RegisterProvisioningRoutes(router, handler)
}
