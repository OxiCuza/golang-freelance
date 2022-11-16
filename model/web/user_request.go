package web

type UserCreateRequest struct {
	Name  string `validate:"required,max=20,min=3"`
	Email string `validate:"required,email"`
}

type UserUpdateRequest struct {
	Id   string `validate:"required"`
	Name string `validate:"required,max=20,min=3"`
}
