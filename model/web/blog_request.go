package web

type BlogCreateRequest struct {
	Title   string `validate:required,min=3,max=20`
	Content string `validate:required,min=20`
	UserId  string `validate:required`
}

type BlogUpdateRequest struct {
	Id      string `validate:required`
	Title   string `validate:required,min=3,max=20`
	Content string `validate:required,min=20`
}
