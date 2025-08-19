package database

import (
	"Hamar/utils"
	"fmt"

	"github.com/google/uuid"
)

func CreateUserInDB(u utils.UsersRequest) {

	schema := fmt.Sprintf("INSERT INTO users (user_id,username,password_hash,access_level) VALUES (%s,%s,%s,%b)", uuid.New(), u.UserName, u.Passowrd, u.AccessLevel)

}
