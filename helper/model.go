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
