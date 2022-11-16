package web

type BlogCreateRequest struct {
	Title   string `json:"title" validate:"required,min=3,max=20"`
	Content string `json:"content" validate:"required,min=20"`
	UserId  string `json:"user_id" validate:"required"`
}

type BlogUpdateRequest struct {
	Id      string `validate:"required"`
	Title   string `validate:"required,min=3,max=20"`
	Content string `validate:"required,min=20"`
}
