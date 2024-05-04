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

func (r *CatRepository) Insert(ctx context.Context, in model.Cats) (int, error) {
	
	query := `
				INSERT INTO cats (user_id, name, race, sex, ageinmonth, description, imageurls, hasmatched, createdAt)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
				RETURNING id`
	var cat_id int
	err := r.db.QueryRowContext(ctx, query, in.User_id, in.Name, in.Race, in.Sex, in.AgeInMonth, in.Description, in.ImageUrls, in.HasMatched, in.CreatedAt).Scan(&cat_id)
	
	
	if err != nil {
		return 0, err
	}

	return cat_id, nil
}