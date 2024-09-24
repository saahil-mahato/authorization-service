package handlers

import (
	"fmt"
	"net/http"

	v1 "buf.build/gen/go/permifyco/permify/protocolbuffers/go/base/v1"
	"github.com/gin-gonic/gin"
	"github.com/saahil-mahato/authorization-service/internal/models"
	"github.com/saahil-mahato/authorization-service/internal/service"
)

type DataHandler struct {
	dataService service.DataService
}

func NewDataHandler(dataService service.DataService) *DataHandler {
	return &DataHandler{dataService: dataService}
}

func (h *DataHandler) WriteAuthorizationData(c *gin.Context) {
	tenantID := c.Param("id")

	var payload models.AuthorizationDataPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	metaData := &v1.DataWriteRequestMetadata{
		SchemaVersion: payload.SchemaVersion,
	}

	entity := &v1.Entity{
		Type: payload.EntityType,
		Id:   payload.EntityID,
	}

	subject := &v1.Subject{
		Type:     payload.SubjectType,
		Id:       payload.SubjectID,
		Relation: payload.SubjectRelation,
	}

	authorizationData, err := h.dataService.WriteAuthorizationData(tenantID, payload.Relation, payload.Attribute, payload.Value, metaData, entity, subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to write authorization data for tenant %s and schema version %s. Error %s", tenantID, metaData.SchemaVersion, err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"authorization_data": authorizationData})
}
