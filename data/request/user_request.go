package request

type CreateUserRequest struct {
	Username string `validate:"required,min=1,max=200" json:"username"`
	Password string `validate:"required,min=1,max=200" json:"password"`
	Email    string `validate:"required,email" json:"email"`
}

type UpdateUserRequest struct {
	Id       int    `validate:"required" json:"id"`
	Username string `validate:"required,min=1,max=200" json:"username"`
	Password string `validate:"required,min=1,max=200" json:"password"`
	Email    string `validate:"required,email" json:"email"`
}

type DeleteUserRequest struct {
	Id int `validate:"required" json:"id"`
}

type FindUserRequest struct {
	Id int `validate:"required" json:"id"`
}
