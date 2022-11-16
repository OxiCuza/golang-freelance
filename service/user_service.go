package service

import (
	"context"
	"golang-freelance/model/web"
)

type UserService interface {
	Save(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userId string)
	FindById(ctx context.Context, userId string) web.UserResponse
	FindAll(ctx context.Context) []web.UserResponse
}
