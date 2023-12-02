package handler

import (
	"essy_travel/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateAirport godoc
// @ID create_airport
// @Router /airport [POST]
// @Summary Create Airport
// @Description Create Airport
// @Tags Airport
// @Accept json
// @Produce json
// @Param object body models.CreateAirport true "CreateAirportRequestBody"
// @Success 200 {object} Response{data=models.Airport} "AirportBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateAirport(c *gin.Context) {
	var airport = models.CreateAirport{}
	err := c.ShouldBindJSON(&airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	resp, err := h.strg.Airport().Create(airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdAirport godoc
// @ID get_by_id_airport
// @Router /airport/{id} [GET]
// @Summary Get By Id Airport
// @Description Get By Id Airport
// @Tags Airport
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.Airport} "AirportBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) AirportGetById(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.strg.Airport().GetById(models.AirportPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, 500, "Airport does not exist: "+err.Error())
		return
	}
	if resp == nil {
		handleResponse(c, http.StatusNoContent, "The data with id is not in the database...")
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListAirport godoc
// @ID get_list_airport
// @Router /airport [GET]
// @Summary Get List Airport
// @Description Get List Airport
// @Tags Airport
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success 200 {object} Response{data=models.GetListAirportResponse} "GetListAirportResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) AirportGetList(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	resp, err := h.strg.Airport().GetList(models.GetListAirportRequest{Offset: offset, Limit: limit})
	if err != nil {
		handleResponse(c, 500, "Airport does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateAirport godoc
// @ID update_airport
// @Router /airport/{id} [PUT]
// @Summary Update Airport
// @Description Update Airport
// @Tags Airport
// @Accept json
// @Produce json
// @Param id path string true "AirportPrimaryKey_ID"
// @Param object body models.UpdateAirport true "UpdateAirportBody"
// @Success 200 {object} Response{data=string} "Updated Airport"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) AirportUpdate(c *gin.Context) {
	var airport = models.UpdateAirport{}
	err := c.ShouldBindJSON(&airport)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	id := c.Param("id")
	airport.Id = id
	resp, err := h.strg.Airport().Update(airport)
	if err != nil {
		handleResponse(c, 500, "Airport does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteAirport godoc
// @ID delete_airport
// @Router /airport/{id} [DELETE]
// @Summary Delete Airport
// @Description Delete Airport
// @Tags Airport
// @Accept json
// @Produce json
// @Param id path string true "DeleteAirportPath"
// @Success 200 {object} Response{data=string} "Deleted Airport"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) AirportDelete(c *gin.Context) {
	id := c.Param("id")
	_, err := h.strg.Airport().Delete(models.AirportPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, 500, "Airport does not delete: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Deleted:")
}
