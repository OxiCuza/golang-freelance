package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang-freelance/exception"
	"golang-freelance/helper"
	"golang-freelance/model/domain"
	"golang-freelance/model/web"
	"golang-freelance/repository"
	"time"
)

type BlogServiceImpl struct {
	BlogRepository repository.BlogRepository
	db             *sql.DB
	Validate       *validator.Validate
}

func NewBlogService(blogRepository repository.BlogRepository, db *sql.DB, validate *validator.Validate) BlogService {
	return &BlogServiceImpl{BlogRepository: blogRepository, db: db, Validate: validate}
}

func (service *BlogServiceImpl) Save(ctx context.Context, request web.BlogCreateRequest) web.BlogResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	blog := &domain.BlogPost{
		Id:        uuid.NewString(),
		Title:     request.Title,
		Content:   request.Content,
		UserId:    request.UserId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	blogPost := service.BlogRepository.Save(ctx, tx, blog)

	return helper.ToBlogPostResponse(blogPost)
}

func (service *BlogServiceImpl) Update(ctx context.Context, request web.BlogUpdateRequest) web.BlogResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	blog, err := service.BlogRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	blog.Title = request.Title
	blog.Content = request.Content
	blog.UpdatedAt = time.Now()

	blogPost := service.BlogRepository.Update(ctx, tx, blog)

	return helper.ToBlogPostResponse(blogPost)
}

func (service *BlogServiceImpl) Delete(ctx context.Context, blogId string) {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	blog, err := service.BlogRepository.FindById(ctx, tx, blogId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.BlogRepository.Delete(ctx, tx, blog)
}

func (service *BlogServiceImpl) FindById(ctx context.Context, blogId string) web.BlogResponse {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	blog, err := service.BlogRepository.FindById(ctx, tx, blogId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToBlogPostResponse(blog)
}

func (service *BlogServiceImpl) FindAll(ctx context.Context) []web.BlogResponse {
	tx, err := service.db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	blog := service.BlogRepository.FindAll(ctx, tx)

	return helper.ToBlogPostResponses(blog)
}
