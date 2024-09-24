package handlers

import (
	"fmt"
	"net/http"

	v1 "buf.build/gen/go/permifyco/permify/protocolbuffers/go/base/v1"
	"github.com/gin-gonic/gin"
	"github.com/saahil-mahato/authorization-service/internal/models"
	"github.com/saahil-mahato/authorization-service/internal/service"
)

type PermissionHandler struct {
	permissionService service.PermissionService
}

func NewPermissionHandler(permissionService service.PermissionService) *PermissionHandler {
	return &PermissionHandler{permissionService: permissionService}
}

func (h *PermissionHandler) CheckAccessControl(c *gin.Context) {
	tenantID := c.Param("id")

	var payload models.CheckAccessControlPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	metaData := &v1.PermissionCheckRequestMetadata{
		SchemaVersion: payload.SchemaVersion,
		SnapToken:     payload.SnapToken,
		Depth:         payload.Depth,
	}

	entity := &v1.Entity{
		Type: payload.EntityType,
		Id:   payload.EntityID,
	}

	subject := &v1.Subject{
		Type: payload.SubjectType,
		Id:   payload.SubjectID,
	}

	permission, err := h.permissionService.CheckAccessControl(tenantID, payload.Permission, metaData, entity, subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to check permission for tenant %s. Error %s", tenantID, err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"permission": permission})
}
