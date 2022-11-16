package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-freelance/helper"
	"golang-freelance/model/domain"
)

type BlogRepositoryImpl struct {
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
	query := "SELECT id, title, content, user_id FROM blog_posts"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var blogs []domain.BlogPost
	for rows.Next() {
		blog := domain.BlogPost{}
		rows.Scan(&blog.Id, &blog.Title, &blog.Content, &blog.UserId)
		blogs = append(blogs, blog)
	}

	return &blogs
}
