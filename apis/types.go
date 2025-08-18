package apis

type UsersRequest struct {
	UserName    string `json:"username" validate:"required,min=5,max=128"`
	Passowrd    string `json:"-" validate:"required min=8"`
	AccessLevel int8   `json:"access_level" validate:"required"`
	Email       string `json:"email" validate:"omitempty,email"`
}
