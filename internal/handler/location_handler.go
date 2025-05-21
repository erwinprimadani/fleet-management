package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/erwinprimadani/fleet-management/internal/models"
	"github.com/erwinprimadani/fleet-management/internal/usecase"
	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	Usecase *usecase.LocationUsecase
}

func NewLocationHandler(u *usecase.LocationUsecase) *LocationHandler {
	return &LocationHandler{Usecase: u}
}

func (h *LocationHandler) GetLatestLocation(c *gin.Context) {
	id := c.Param("vehicle_id")
	loc, err := h.Usecase.GetLatest(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, loc)
}

func (h *LocationHandler) GetLocationHistory(c *gin.Context) {
	id := c.Param("vehicle_id")
	start, _ := strconv.ParseInt(c.Query("start"), 10, 64)
	end, _ := strconv.ParseInt(c.Query("end"), 10, 64)

	startInRFC := time.Unix(start, 0).Format(time.RFC3339)
	endInRFC := time.Unix(end, 0).Format(time.RFC3339)

	history, err := h.Usecase.GetHistory(id, startInRFC, endInRFC)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, history)
}

// SendLatestLocation handles the POST request to send the latest location of a vehicle
func (h *LocationHandler) SendLatestLocation(c *gin.Context) {
	var location models.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid request body: %v", err),
		})
		return
	}

	if location.Latitude == 0 && location.Longitude == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "latitude and longitude are required",
		})
		return
	}

	if location.Timestamp == "" {
		// Set current timestamp if not provided
		location.Timestamp = time.Unix(time.Now().Unix(), 0).Format(time.RFC3339)
	}

	err := h.Usecase.SendDataToMQTT(location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to publish location: %v", err),
		})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
		"message": "Location sent successfully",
		"data":    location,
	})
}

func (h *LocationHandler) Healthcheck(c *gin.Context) {

	c.JSON(http.StatusOK, "pong")
}
