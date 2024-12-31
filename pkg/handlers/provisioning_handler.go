package handlers

import (
	"multicloud-provisioner/pkg/models"
	"multicloud-provisioner/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProvisioningHandler handles resource provisioning requests
type ProvisioningHandler struct {
	provisioningService *services.ProvisioningService
}

// NewProvisioningHandler creates a new provisioning handler
func NewProvisioningHandler(
	provisioningService *services.ProvisioningService,
) *ProvisioningHandler {
	return &ProvisioningHandler{
		provisioningService: provisioningService,
	}
}

// CreateResourceHandler handles requests to create new resources
func (h *ProvisioningHandler) CreateResourceHandler(c *gin.Context) {
	var config models.Configuration
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Call the service to create resources
	provisionedResources, err := h.provisioningService.CreateResource(config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"provisionedResources": provisionedResources})
}

// DeleteResourceHandler handles requests to delete a resource
func (h *ProvisioningHandler) DeleteResourceHandler(c *gin.Context) {
	resourceID := c.Param("id")
	providerName := c.Param("provider")

	// Call the service to delete the resource
	err := h.provisioningService.DeleteResource(providerName, resourceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Resource deleted successfully"})
}

// GetResourceHandler handles requests to retrieve resource details
func (h *ProvisioningHandler) GetResourceHandler(c *gin.Context) {
	resourceID := c.Param("id")
	providerName := c.Param("provider")

	// Call the service to retrieve resource metadata
	resourceMetadata, err := h.provisioningService.GetResource(providerName, resourceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"resource": resourceMetadata})
}
