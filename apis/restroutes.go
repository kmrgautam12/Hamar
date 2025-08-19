package apis

import (
	"Hamar/database"
	"Hamar/utils"
	"encoding/base64"
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterRoutes(e *E) {
	e.Cho.POST("/signup", SignupUser)
	e.Cho.POST("/login", SignupUser)
}

func SignupUser(e echo.Context) error {
	var user utils.UsersRequest
	err := e.Bind(&user)
	if err != nil {
		return err
	}
	ValidateSignupStruct(user)
	encodedStr := base64.StdEncoding.EncodeToString([]byte(user.Passowrd))
	b, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf(encodedStr, encodedStr[:5])), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Passowrd = string(b)
	database.CreateUserInDB(user)
	return nil
}
