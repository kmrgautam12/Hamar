package apis

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type E struct {
	Cho *echo.Echo
	DB  *sql.DB
}
