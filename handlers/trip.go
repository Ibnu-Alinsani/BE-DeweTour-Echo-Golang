package handlers

import (
	dto "dumbmerch/dto/result"
	tripdto "dumbmerch/dto/trip"
	"dumbmerch/models"
	"dumbmerch/repository"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type TripHandler struct {
	TripRepository repository.TripRepository
}

func HandlerTrip(TripRepository repository.TripRepository) *TripHandler {
	return &TripHandler{TripRepository}
}

var path_file = "http://localhost:5000/uploads/"

func (h *TripHandler) GetAllTrip(c echo.Context) error {
	dataResponse, err := h.TripRepository.GetAllTrip()

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

func (h *TripHandler) GetTripById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	trip, err := h.TripRepository.GetTripById(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}


	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: trip,
	})
}

func (h *TripHandler) CreateTrip(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	fmt.Println("iyeu photo urang", dataFile)	

	price, _ := strconv.Atoi(c.FormValue("price"))
	quota, _ := strconv.Atoi(c.FormValue("quota"))
	fmt.Println(quota)
	countryId, _ := strconv.Atoi(c.FormValue("country_id"))

	Trip := models.Trip{
		Title:          c.FormValue("title"),
		CountryId:      countryId,
		Accomodation:   c.FormValue("accomodation"),
		Transportation: c.FormValue("transportation"),
		Eat:            c.FormValue("eat"),
		Day:            c.FormValue("day"),
		Night:          c.FormValue("night"),
		DateTrip:       c.FormValue("dateTrip"),
		Price:          price,
		Quota:          quota,
		CurrentQuota: 	quota,
		Description:    c.FormValue("description"),
		Image:          dataFile,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	fmt.Println(Trip)

	validation := validator.New()
	err := validation.Struct(Trip)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	country, err := h.TripRepository.CreateTrip(Trip)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusOK,
			Message: err.Error(),
		})
	}
	
	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: country,
	})
}

func (h *TripHandler) DeleteTrip(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	trip, err := h.TripRepository.GetTripById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	deleteTrip, err := h.TripRepository.DeleteTrip(trip)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}	

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: deleteTrip,
	})
}

func (h *TripHandler) EditTrip(c echo.Context) error {
	var request tripdto.UpdateTrip


	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	trip, err := h.TripRepository.GetTripById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	dataFile := c.Get("dataFile").(string)
	
	countryId, _ := strconv.Atoi(c.FormValue("country_id"))
	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	quota, err := strconv.Atoi(c.FormValue("quota"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	title := c.FormValue("title")
	transportation := c.FormValue("transportation")
	accomodation := c.FormValue("accomodation")
	eat :=            c.FormValue("eat")
	day :=            c.FormValue("day")
	night :=          c.FormValue("night")
	dateTrip :=       c.FormValue("dateTrip")
	description :=    c.FormValue("description")
	
	getId, err := h.TripRepository.GetCountryByIdForTrip(countryId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	
	if title != "" {
		trip.Title = title
	}

	if request.CurrentQuota != 0 {
		trip.CurrentQuota = request.CurrentQuota
	}

	trip.Country = models.CountryResponse(getId)

	if accomodation != "" {
		trip.Accomodation = accomodation
	}
	if transportation != "" {
		trip.Transportation = transportation
	}
	if eat != "" {
		trip.Eat = eat
	}
	if day != "" {
		trip.Day = day
	}
	if night != "" {
		trip.Night = night
	}
	if dateTrip != "" {
		trip.DateTrip = dateTrip
	}
	if price != 0 {
		trip.Price = price
	}
	if quota != 0 {
		trip.Quota = quota
	}
	if description != "" {
		trip.Description = description
	}
	trip.Image = dataFile

	trip.UpdatedAt = time.Now()

	dataResponse, err := h.TripRepository.EditTrip(trip)
	fmt.Println(dataResponse)

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

func ConvertResponseTrip(t models.Trip) models.TripResponse {
	return models.TripResponse{
		Id:             0,
		Title:          t.Title,
		CountryId:      t.CountryId,
		Country:        t.Country,
		Accomodation:   t.Accomodation,
		Transportation: t.Transportation,
		Eat:            t.Eat,
		Day:            t.Day,
		Night:          t.Night,
		DateTrip:       t.DateTrip,
		Price:          t.Price,
		Quota:          t.Quota,
		Description:    t.Description,
	}
}