package web

type UserCreateRequest struct {
	Name  string `json:"name" validate:"required,max=20,min=3"`
	Email string `json:"email" validate:"required,email"`
}

type UserUpdateRequest struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,max=20,min=3"`
}
