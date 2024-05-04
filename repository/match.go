package repository

import (
	"context"
	"database/sql"

	"github.com/Project-Sprint/w1-social-media/model"
)

type MatchRepository struct {
	db *sql.DB
}

func NewMatchRepository(db *sql.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

func (r *MatchRepository) Insert(ctx context.Context, in model.Match) (interface{}, error) {
	query := `
		INSERT INTO matchs (matchCatId, userCatId, userId, targetUserId, message, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6)

	`

	result, err := r.db.ExecContext(ctx, query, in.MatchCatId, in.UserCatId, in.UserId, in.TargetUserId, in.Message, in.CreatedAt)
	if err != nil {
		return nil, err
	}

	// query = `
	// 	UPDATE cats
	// 	SET hasMatched = true
	// 	WHERE id IN ($1, $2)
	// `

	// result, err := r.db.ExecContext(ctx, query, in.MatchCatId, in.UserCatId)
	// if err != nil {
	// 	return nil, err
	// }

	return result, nil
}

func (r *MatchRepository) FindCatById(ctx context.Context, catId int) (model.Cats, error) {
	query := `
		SELECT id, user_id, name, race, sex, ageInMonth, description, imageUrls, hasMatched, createdAt
		FROM cats
		WHERE id = $1
	`

	var cat model.Cats
	if err := r.db.QueryRowContext(ctx, query, catId).Scan(&cat.Id, &cat.User_id, &cat.Name, &cat.Race, &cat.Sex, &cat.AgeInMonth, &cat.Description, &cat.ImageUrls, &cat.HasMatched, &cat.CreatedAt); err != nil {
		return cat, err
	}

	return cat, nil
}

func (r *MatchRepository) Approve(ctx context.Context, in model.RequestMatchApprove) (interface{}, error) {
	// query := `
	// 	INSERT INTO matchs (matchCatId, userCatId, userId, targetUserId, message, createdAt)
	// 	VALUES ($1, $2, $3, $4, $5, $6)
	// `

	// result, err := r.db.ExecContext(ctx, query, in.MatchCatId, in.UserCatId, in.UserId, in.TargetUserId, in.Message, in.CreatedAt)
	// if err != nil {
	// 	return nil, err
	// }

	// // query = `
	// // 	UPDATE cats
	// // 	SET hasMatched = true
	// // 	WHERE id IN ($1, $2)
	// // `

	// // result, err := r.db.ExecContext(ctx, query, in.MatchCatId, in.UserCatId)
	// // if err != nil {
	// // 	return nil, err
	// // }

	// return result, nil
	return nil, nil
}
