package user

import (
	"database/sql"
	service "pet-info-backend-docker/aplication/user"
	handler "pet-info-backend-docker/infraestructure/handler/user"
	repository "pet-info-backend-docker/infraestructure/repository/user"

	"github.com/labstack/echo/v4"
)

// RegisterUserRoutes registra las rutas de usuario en el enrutador Echo
func RegisterUserRoutes(e *echo.Echo, db *sql.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewHandler(userService)

	e.POST("/users/create", userHandler.CreateUser)
	e.GET("/users/get-by", userHandler.GetUser)
}
