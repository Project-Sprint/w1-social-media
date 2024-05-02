package repository

import (
	"context"
	"database/sql"

	"github.com/Project-Sprint/w1-social-media/model"
)

type MatchRepository struct {
	db *sql.DB
}

func NewMatchRepository(db *sql.DB) *MatchRepository{
	return &MatchRepository{db: db}
}

func (r *MatchRepository) Insert(ctx context.Context, user model.Match) error {
	return nil
}