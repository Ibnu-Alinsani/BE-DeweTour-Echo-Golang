package handlers

import (
	authdto "dumbmerch/dto/auth"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/pkg/bcrypt"
	jwtToken "dumbmerch/pkg/jwt"
	"dumbmerch/repository"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	AuthRepository repository.AuthRepository
}

func HandlerAuth(AuthRepository repository.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(c echo.Context) error {
	request := new(authdto.AuthRequest)
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

	password, err := bcrypt.HashingPassword(request.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user := models.User {
		FullName: request.FullName,
		Email: request.Email, 
		Password: password,
		Phone: request.Phone,
		Address: request.Address,
		Role: "user",
	}

	fmt.Println(user)

	data, err := h.AuthRepository.Register(user)
	fmt.Println(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: data,
	})
}

// Login Auth

func (h *handlerAuth) Login(c echo.Context) error {
	request := new(authdto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user := models.User {
		Email: request.Email,
		Password: request.Password,
	}

	// check email
	user, err := h.AuthRepository.Login(user.Email)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// check password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	fmt.Println(isValid, "ini is valid")
	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour*2).Unix()
	
	token, err := jwtToken.GenerateToken(&claims)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := authdto.LoginResponse {
		Id : 		user.ID,
		FullName:   user.FullName,
		Email:    	user.Email,
		Role: 		user.Role,
		Token:    	token,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: loginResponse,
	})
}

func (h *handlerAuth) CheckAuth(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, _ := h.AuthRepository.CheckAuth(int(userId))


	user.Image = path_file + user.Image

	return c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: ConvertResponseUser(user),
	})
}
