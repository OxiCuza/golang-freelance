package domain

type BlogPost struct {
	BaseModel
	Title   string
	Content string
	UserId  string
}
