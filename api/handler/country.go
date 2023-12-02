package handler

import (
	"essy_travel/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCountry godoc
// @ID create_country
// @Router /country [POST]
// @Summary Create Country
// @Description Create Country
// @Tags Country
// @Accept json
// @Produce json
// @Param object body models.CreateCountry true "CreateCountryRequestBody"
// @Success 200 {object} Response{data=models.Country} "CountryBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCountry(c *gin.Context) {
	var country = models.CreateCountry{}
	err := c.ShouldBindJSON(&country)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	resp, err := h.strg.Country().Create(country)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdCountry godoc
// @ID get_by_id_country
// @Router /country/{id} [GET]
// @Summary Get By Id Country
// @Description Get By Id Country
// @Tags Country
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.Country} "CountryBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CountryGetById(c *gin.Context) {
	id := c.Param("id")
	resp, err := h.strg.Country().GetById(models.CountryPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, 500, "Country does not exist: "+err.Error())
		return
	}
	if resp == nil {
		handleResponse(c, http.StatusNoContent, "The data with id is not in the database...")
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListCountry godoc
// @ID get_list_country
// @Router /country [GET]
// @Summary Get List Country
// @Description Get List Country
// @Tags Country
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success 200 {object} Response{data=models.GetListCountryResponse} "GetListCountryResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CountryGetList(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	resp, err := h.strg.Country().GetList(models.GetListCountryRequest{Offset: offset, Limit: limit})
	if err != nil {
		handleResponse(c, 500, "Country does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateCountry godoc
// @ID update_country
// @Router /country/{id} [PUT]
// @Summary Update Country
// @Description Update Country
// @Tags Country
// @Accept json
// @Produce json
// @Param id path string true "CountryPrimaryKey_ID"
// @Param object body models.UpdateCountry true "UpdateCountryBody"
// @Success 200 {object} Response{data=string} "Updated Country"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CountryUpdate(c *gin.Context) {
	var country = models.UpdateCountry{}
	err := c.ShouldBindJSON(&country)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	id := c.Param("id")
	country.Id = id
	resp, err := h.strg.Country().Update(country)
	if err != nil {
		handleResponse(c, 500, "Country does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteCountry godoc
// @ID delete_country
// @Router /country/{id} [DELETE]
// @Summary Delete Country
// @Description Delete Country
// @Tags Country
// @Accept json
// @Produce json
// @Param id path string true "DeleteCountryPath"
// @Success 200 {object} Response{data=string} "Deleted Country"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CountryDelete(c *gin.Context) {
	id := c.Param("id")
	_, err := h.strg.Country().Delete(models.CountryPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, 500, "Country does not delete: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Deleted:")
}
