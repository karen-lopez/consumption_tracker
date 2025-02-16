package handlers

import (
	"consumption_tracker/cmd/internal/core/ports"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ConsumptionHandler struct {
	Service ports.ConsumptionTrackerService
}

func NewConsumptionHandler(service ports.ConsumptionTrackerService) *ConsumptionHandler {
	return &ConsumptionHandler{Service: service}
}

// GetConsumption godoc
// @Summary Get energy consumption data
// @Description Get energy consumption data for a specific meter within a date range
// @Tags consumption
// @Accept json
// @Produce json
// @Param meter_id query int true "Meter ID"
// @Param start_date query string true "Start date in YYYY-MM-DD format"
// @Param end_date query string true "End date in YYYY-MM-DD format"
// @Param kind_period query string false "Period type (daily, weekly, monthly)"
// @Success 200 {object} dtos.ConsumptionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /consumption [get]
func (h *ConsumptionHandler) GetConsumption(c *gin.Context) {
	meterID, parseErr := strconv.Atoi(c.Query("meter_id"))
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	kindPeriod := c.Query("kind_period")

	if parseErr != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "meter_id is required"})
		return
	}
	consumptions, err := h.Service.GetConsumption(c.Request.Context(), meterID, startDate, endDate, kindPeriod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	if consumptions == nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "no consumption data found"})
		return
	}

	c.JSON(http.StatusOK, consumptions)
}

// ErrorResponse defines the structure for error responses
type ErrorResponse struct {
	Error string `json:"error"`
}
