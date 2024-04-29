package repository

import (
	"context"
	"database/sql"

	"github.com/Project-Sprint/w1-social-media/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Insert(ctx context.Context, user model.User) error {
	return nil
}
