package repository

import (
	"context"
	"database/sql"
	"golang-freelance/model/domain"
)

type BlogRepository interface {
	Save(ctx context.Context, tx *sql.Tx, post *domain.BlogPost) *domain.BlogPost
	Update(ctx context.Context, tx *sql.Tx, post *domain.BlogPost) *domain.BlogPost
	Delete(ctx context.Context, tx *sql.Tx, post *domain.BlogPost)
	FindById(ctx context.Context, tx *sql.Tx, blogId string) (*domain.BlogPost, error)
	FindAll(ctx context.Context, tx *sql.Tx) *[]domain.BlogPost
}
