package repository

import (
	"context"
	"database/sql"

	"github.com/Project-Sprint/w1-social-media/model"
)

type CatRepository struct {
	db *sql.DB
}

func NewCatRepository(db *sql.DB) *CatRepository{
	return &CatRepository{db: db}
}

func (r *CatRepository) Insert(ctx context.Context, user model.Cats) error {
	return nil
}