package handler

import (
	"essy_travel/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCity godoc
// @ID create_city
// @Router /city [POST]
// @Summary Create City
// @Description Create City
// @Tags City
// @Accept json
// @Produce json
// @Param object body models.CreateCity true "CreateCityRequestBody"
// @Success 200 {object} Response{data=models.City} "CityBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCity(c *gin.Context) {
	var city = models.CreateCity{}
	err := c.ShouldBindJSON(&city)
	if err != nil {
		c.JSON(400, "ShouldBindJSON err:"+err.Error())
		return
	}

	resp, err := h.strg.City().Create(city)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Does not create"+err.Error())
		return
	}
	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdCity godoc
// @ID get_by_id_city
// @Router /city/{id} [GET]
// @Summary Get By Id City
// @Description Get By Id City
// @Tags City
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.City} "GetListCityResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CityGetById(c *gin.Context) {
	var id = c.Param("id")

	resp, err := h.strg.City().GetById(models.CityPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, 500, "City does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// GetListCity godoc
// @ID get_list_city
// @Router /city [GET]
// @Summary Get List City
// @Description Get List City
// @Tags City
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success 200 {object} Response{data=models.GetListCityResponse} "GetListCityResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CityGetList(c *gin.Context) {
	offset, err := h.getIntegerOrDefaultValue(c.Query("offset"), 0)
	if err != nil {
		handleResponse(c, 400, "invalid offset")
		return
	}

	limit, err := h.getIntegerOrDefaultValue(c.Query("limit"), 0)
	if err != nil {
		handleResponse(c, 400, "invalid limit")
		return
	}

	resp, err := h.strg.City().GetList(models.GetListCityRequest{
		Offset: int(offset),
		Limit:  int(limit),
	})
	if err != nil {
		handleResponse(c, 500, "City does not exist: "+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateCity godoc
// @ID update_city
// @Router /city/{id} [PUT]
// @Summary Update City
// @Description Update City
// @Tags City
// @Accept json
// @Produce json
// @Param id path string true "PrimaryKey"
// @Param object body models.UpdateCity true "UpdateCityBody"
// @Success 200 {object} Response{data=models.City} "Updated City"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CityUpdate(c *gin.Context) {
	var city = models.UpdateCity{}
	err := c.ShouldBindJSON(&city)

	if err != nil {
		handleResponse(c, http.StatusBadRequest, "Error while json decoding"+err.Error())
		return
	}
	id := c.Param("id")
	city.Id = id
	resp, err := h.strg.City().Update(city)
	if err != nil {
		handleResponse(c, 500, "City does not update: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteCity godoc
// @ID update_city
// @Router /city/{id} [DELETE]
// @Summary Delete City
// @Description Delete City
// @Tags City
// @Accept json
// @Produce json
// @Param id path string true "DeleteCityPath"
// @Success 200 {object} Response{data=string} "Deleted City"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CityDelete(c *gin.Context) {
	id := c.Param("id")

	_, err := h.strg.City().Delete(models.CityPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, 500, "City does not delete: "+err.Error())
		return
	}

	handleResponse(c, http.StatusAccepted, "Deleted:")
}
