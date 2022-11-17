package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"golang-freelance/helper"
	"golang-freelance/model/domain"
	"log"
)

type BlogRepositoryImpl struct {
	RedisClient helper.Redis
}

func NewBlogRepository(redisClient helper.Redis) BlogRepository {
	return &BlogRepositoryImpl{RedisClient: redisClient}
}

func (repository *BlogRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, blog *domain.BlogPost) *domain.BlogPost {
	query := "INSERT INTO blog_posts(id, title, content, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, query, blog.Id, blog.Title, blog.Content, blog.UserId, blog.CreatedAt, blog.UpdatedAt)
	helper.PanicIfError(err)

	return blog
}

func (repository *BlogRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, blog *domain.BlogPost) *domain.BlogPost {
	query := "UPDATE blog_posts SET title = ?, content = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, blog.Title, blog.Content, blog.UpdatedAt, blog.Id)
	helper.PanicIfError(err)

	return blog
}

func (repository *BlogRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, blog *domain.BlogPost) {
	query := "DELETE FROM blog_posts WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, blog.Id)
	helper.PanicIfError(err)
}

func (repository *BlogRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, blogId string) (*domain.BlogPost, error) {
	query := "SELECT id, title, content, user_id FROM blog_posts WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, blogId)
	helper.PanicIfError(err)
	defer rows.Close()

	blog := domain.BlogPost{}
	if rows.Next() {
		err := rows.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.UserId)
		helper.PanicIfError(err)
		return &blog, nil
	}

	return nil, errors.New("blog post not found")
}

func (repository *BlogRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) *[]domain.BlogPost {
	var blogs []domain.BlogPost

	result, err := repository.RedisClient.Get(ctx, "blog")
	if err != nil {
		query := "SELECT id, title, content, user_id FROM blog_posts"
		rows, err := tx.QueryContext(ctx, query)
		helper.PanicIfError(err)
		defer rows.Close()

		for rows.Next() {
			blog := domain.BlogPost{}
			rows.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.UserId)
			blogs = append(blogs, blog)
		}

		err = repository.RedisClient.Set(ctx, "blog", blogs)
		helper.PanicIfError(err)

		log.Println("set operation success")

		return &blogs
	}

	log.Println("get operation success")

	err = json.Unmarshal([]byte(result), &blogs)
	helper.PanicIfError(err)

	return &blogs
}
