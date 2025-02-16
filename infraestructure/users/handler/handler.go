package handler

import (
	"net/http"
	"pet-info-backend-docker/infraestructure/handler/response"
	"pet-info-backend-docker/infraestructure/users/model"
	"pet-info-backend-docker/infraestructure/users/service"

	"github.com/labstack/echo/v4"
)

// Handler representa un controlador para manejar solicitudes HTTP
type Handler struct {
	Service *service.UserService
}

// NewHandler crea una nueva instancia de Handler
func NewHandler(service *service.UserService) *Handler {
	return &Handler{Service: service}
}

// CreateUser maneja la solicitud para crear un nuevo usuario
func (handler *Handler) CreateUser(echo echo.Context) error {
	var user model.User
	if err := echo.Bind(&user); err != nil {
		return echo.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Invalid input"})
	}

	if err := handler.Service.CreateUser(&user); err != nil {
		return echo.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to create user"})
	}

	return echo.JSON(http.StatusCreated, response.SuccessResponse{Message: "User created successfully", Data: user})
}

// GetUser maneja la solicitud para obtener un usuario por email o username
func (handler *Handler) GetUser(echo echo.Context) error {
	var user model.User
	if err := echo.Bind(&user); err != nil {
		return echo.JSON(http.StatusBadRequest, response.ErrorResponse{Error: "Invalid input"})
	}

	users, err := handler.Service.GetUser(user.Email, user.Username)
	if err != nil {
		return echo.JSON(http.StatusInternalServerError, response.ErrorResponse{Error: "Failed to query user"})
	}

	if len(users) == 0 {
		return echo.JSON(http.StatusNotFound, response.ErrorResponse{Error: "User not found"})
	}

	return echo.JSON(http.StatusOK, response.SuccessResponse{Message: "Users found", Data: users})
}
