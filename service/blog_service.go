package service

import (
	"context"
	"golang-freelance/model/web"
)

type BlogService interface {
	Save(ctx context.Context, request web.BlogCreateRequest) web.BlogResponse
	Update(ctx context.Context, request web.BlogUpdateRequest) web.BlogResponse
	Delete(ctx context.Context, blogId string)
	FindById(ctx context.Context, blogId string) web.BlogResponse
	FindAll(ctx context.Context) []web.BlogResponse
}
