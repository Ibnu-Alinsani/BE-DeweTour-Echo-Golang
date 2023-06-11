package handlers

import (
	dto "dumbmerch/dto/result"
	userdto "dumbmerch/dto/users"
	"dumbmerch/models"
	"dumbmerch/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepository repository.UserRepository
}

func HandlerUser(UserRepository repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepository}
}

// var path_file = "http://localhost:5000/uploads/"

// Create user or registration
func (h *UserHandler) CreateUser(c echo.Context) error {
	request := new(userdto.CreateUser)

	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user := models.User{
		FullName: request.FullName,
		Email: request.Email,
		Password: request.Password,
		Phone: request.Phone,
		Address: request.Address,
		Gender: "Non Binary",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	dataResponse, err := h.UserRepository.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: ConvertResponseUser(dataResponse),
	})
}

// Get all user
func (h *UserHandler) GetAllUser(c echo.Context) error {
	dataResponse, err := h.UserRepository.GetAllUser()

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

// get user by id
func (h *UserHandler) GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	
	dataResponse, err := h.UserRepository.GetUserById(id)
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: ConvertResponseUser(dataResponse),
	})
}

// Delete user
func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.UserRepository.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	dataResponse, err := h.UserRepository.DeleteUser(user)

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

// update user
func (h *UserHandler) UpdateUser(c echo.Context) error {
	dataFile := c.Get("dataFile").(string)
	request := new(userdto.UpdateUser)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.UserRepository.GetUserById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	
	fullname := c.FormValue("fullname");
	email := c.FormValue("email");
	phone := c.FormValue("phone");
	address := c.FormValue("address");

	if  fullname != "" {
		user.FullName = fullname
	}

	if  email != "" {
		user.Email = email
	}
	if  phone != "" {
		user.Phone = phone
	}
	if  address != "" {
		user.Address = address
	}

	user.Image = dataFile

	user.UpdatedAt = time.Now()

	dataResponse, err := h.UserRepository.UpdateUser(user)

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

func ConvertResponseUser(user models.User) userdto.UserResponse {
	return userdto.UserResponse{
		ID:       user.ID,
		Image: user.Image,
		FullName:     user.FullName,
		Email:    user.Email,
		// Password: user.Password,
		Gender: user.Gender,
		Role: user.Role,
		Phone : user.Phone,
		Address : user.Address,
		Transaction: user.Transaction,
	}
}