package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-freelance/helper"
	"golang-freelance/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (userRepository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user *domain.User) *domain.User {
	query := "INSERT INTO users(id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, query, user.Id, user.Name, user.Email, user.CreatedAt, user.CreatedAt)
	helper.PanicIfError(err)

	return user
}

func (userRepository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user *domain.User) *domain.User {
	query := "UPDATE users SET name = ?, updated_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, user.Name, user.UpdatedAt, user.Id)
	helper.PanicIfError(err)

	return user
}

func (userRepository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user *domain.User) {
	query := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, user.Id)
	helper.PanicIfError(err)
}

func (userRepository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId string) (*domain.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		helper.PanicIfError(err)
		return &user, nil
	}

	return nil, errors.New("user not found")
}

func (userRepository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) *[]domain.User {
	query := "SELECT id, name, email FROM users"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		rows.Scan(&user.Id, &user.Name, &user.Email)
		users = append(users, user)
	}

	return &users
}
