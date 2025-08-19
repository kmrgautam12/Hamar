package apis

import "Hamar/utils"

// ToDo : Through validation of request body
func ValidateSignupStruct(u utils.UsersRequest) bool {
	if len(u.UserName) < 5 || len(u.Passowrd) < 8 {
		return false
	}
	return true
}
