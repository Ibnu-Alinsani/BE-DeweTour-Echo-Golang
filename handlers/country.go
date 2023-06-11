package handlers

import (
	countrydto "dumbmerch/dto/country"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repository"
	"fmt"
	"net/http"
	"strconv"

	// "github.com/golang-jwt/jwt/request"
	"github.com/labstack/echo/v4"
)

type CountryHandler struct {
	CountryRepository repository.CountryRepository
}

func HandlerCountry(CountryRepository repository.CountryRepository) *CountryHandler {
	return &CountryHandler{CountryRepository}
}

func (h *CountryHandler) GetAllCountry(c echo.Context) error {
	dataResponse, err := h.CountryRepository.GetAllCountry()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}

func (h *CountryHandler) GetCountryById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	dataResponse, err := h.CountryRepository.GetCountryById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}

func (h *CountryHandler) AddCountry(c echo.Context) error {
	request := new(countrydto.CountryRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	dataRequest := models.Country {
		Name: request.Name,
	}

	dataResponse, err := h.CountryRepository.AddCountry(dataRequest)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}

func (h *CountryHandler) EditCountry(c echo.Context) error {
	request := new(countrydto.CountryRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	country, err := h.CountryRepository.GetCountryById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	if request.Name != "" {
		country.Name = request.Name
	}

	dataResponse, err := h.CountryRepository.EditCountry(country)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}

func (h *CountryHandler) DeleteCountry(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id,"ini id")
	country, err := h.CountryRepository.GetCountryById(id)
	fmt.Println(country, "ini country")

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	dataResponse ,err := h.CountryRepository.DeleteCountry(country)
	fmt.Println(dataResponse, "ini data response")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
		Code: http.StatusInternalServerError,
		Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}