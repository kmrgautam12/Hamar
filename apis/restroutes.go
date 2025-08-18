package apis

import (
	"encoding/base64"
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterRoutes(e *echo.Echo) {
	e.POST("/signup", SignupUser)
	e.POST("/login", SignupUser)
}

func SignupUser(e echo.Context) error {
	var user UsersRequest
	err := e.Bind(&user)
	if err != nil {
		return err
	}
	encodedStr := base64.StdEncoding.EncodeToString([]byte(user.Passowrd))
	b, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf(encodedStr, encodedStr[:5])), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Passowrd = string(b)

}
