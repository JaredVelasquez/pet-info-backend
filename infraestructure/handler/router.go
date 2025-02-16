package handler

import (
	"database/sql"
	user "pet-info-backend-docker/infraestructure/routers/user"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes registra las rutas en el enrutador Echo
func RegisterRoutes(e *echo.Echo, db *sql.DB) {
	user.RegisterUserRoutes(e, db)
}
