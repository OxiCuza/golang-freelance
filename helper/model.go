package helper

import (
	"golang-freelance/model/domain"
	"golang-freelance/model/web"
)

func ToUserResponse(user *domain.User) web.UserResponse {
	return web.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}

func ToUserResponses(users *[]domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range *users {
		userResponses = append(userResponses, ToUserResponse(&user))
	}

	return userResponses
}

func ToBlogPostResponse(blog *domain.BlogPost) web.BlogResponse {
	return web.BlogResponse{
		Id:      blog.Id,
		Title:   blog.Title,
		Content: blog.Content,
		UserId:  blog.UserId,
	}
}

func ToBlogPostResponses(blogs *[]domain.BlogPost) []web.BlogResponse {
	var blogResponses []web.BlogResponse
	for _, blog := range *blogs {
		blogResponses = append(blogResponses, ToBlogPostResponse(&blog))
	}

	return blogResponses
}
