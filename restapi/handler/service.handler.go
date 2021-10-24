package handler

import (
	"golang-restapi/model"
	"golang-restapi/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateServiceRequest -- create service request
func CreateServiceRequest(c *gin.Context) {
	userID := uint64(c.MustGet("userID").(float64))
	var serviceRequest model.Service
	c.Bind(&serviceRequest)
	repository.CreateServiceRequest(&serviceRequest, userID)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success to create service request",
		"data": gin.H{
			"requestID":   serviceRequest.RequestID,
			"status":      serviceRequest.Status,
			"vesselName":  serviceRequest.VesselName,
			"serviceType": serviceRequest.ServiceType,
			"dataAgent":   serviceRequest.DataAgent,
			"cargo":       serviceRequest.Cargo,
			"etd":         serviceRequest.ETD,
			"eta":         serviceRequest.ETA,
		},
	})
}

// GetServices -- Get all services
func GetServices(c *gin.Context) {
	userID := uint64(c.MustGet("userID").(float64))
	var services []model.Service = repository.GetServices(userID)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success to get all services",
		"data":    services,
	})
}
