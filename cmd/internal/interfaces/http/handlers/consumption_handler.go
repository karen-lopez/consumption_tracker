package handlers

import (
	"consumption_tracker/cmd/internal/application/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ConsumptionHandler struct {
	Service *services.EnergyConsumptionService
}

func NewConsumptionHandler(service *services.EnergyConsumptionService) *ConsumptionHandler {
	return &ConsumptionHandler{Service: service}
}

func (h *ConsumptionHandler) GetConsumption(c *gin.Context) {
	meterID, parseErr := strconv.Atoi(c.Query("meters_ids"))
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	kindPeriod := c.Query("kind_period")

	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "meter_id is required"})
		return
	}
	consumptions, err := h.Service.GetConsumption(c.Request.Context(), meterID, startDate, endDate, kindPeriod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if consumptions == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no consumption data found"})
		return
	}

	c.JSON(http.StatusOK, consumptions)
}
