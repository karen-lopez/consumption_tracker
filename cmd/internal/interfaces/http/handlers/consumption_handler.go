package handlers

import (
	"consumption_tracker/cmd/internal/core/ports"
	"consumption_tracker/cmd/internal/interfaces/http/handlers/dtos"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ConsumptionHandler struct {
	Service ports.ConsumptionTrackerService
}

func NewConsumptionHandler(service ports.ConsumptionTrackerService) *ConsumptionHandler {
	return &ConsumptionHandler{Service: service}
}

// GetConsumption godoc
// @Summary Get energy consumption data
// @Description Get energy consumption data for specific meters within a date range
// @Tags consumption
// @Accept json
// @Produce json
// @Param meters_ids query string true "Comma-separated list of meter IDs"
// @Param start_date query string true "Start date in YYYY-MM-DD format"
// @Param end_date query string true "End date in YYYY-MM-DD format"
// @Param kind_period query string false "Period type (daily, weekly, monthly)"
// @Success 200 {object} dtos.ConsumptionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /consumption [get]
func (h *ConsumptionHandler) GetConsumption(c *gin.Context) {
	meterIDsParam := c.Query("meters_ids")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	kindPeriod := c.Query("kind_period")

	meterIDs, err := parseMeterIDs(meterIDsParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid meters_ids"})
		return
	}

	var allConsumptions []*dtos.ConsumptionResponse
	for _, meterID := range meterIDs {
		consumptions, err := h.Service.GetConsumption(c.Request.Context(), meterID, startDate, endDate, kindPeriod)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
			return
		}
		if consumptions != nil {
			allConsumptions = append(allConsumptions, consumptions)
		}
	}

	if len(allConsumptions) == 0 {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "no consumption data found"})
		return
	}

	combinedResponse := combineResponses(allConsumptions)
	c.JSON(http.StatusOK, combinedResponse)
}

func parseMeterIDs(meterIDsParam string) ([]int, error) {
	meterIDsStr := strings.Split(meterIDsParam, ",")
	meterIDs := make([]int, 0, len(meterIDsStr))
	for _, idStr := range meterIDsStr {
		id, err := strconv.Atoi(strings.TrimSpace(idStr))
		if err != nil {
			return nil, err
		}
		meterIDs = append(meterIDs, id)
	}
	return meterIDs, nil
}

type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

func combineResponses(responses []*dtos.ConsumptionResponse) *dtos.ConsumptionResponse {
	if len(responses) == 0 {
		return nil
	}

	combined := &dtos.ConsumptionResponse{
		Period:    responses[0].Period,
		DataGraph: []*dtos.MeterData{},
	}

	for _, response := range responses {
		combined.DataGraph = append(combined.DataGraph, response.DataGraph...)
	}

	return combined
}
