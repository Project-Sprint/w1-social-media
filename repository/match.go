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

func (r *MatchRepository) Approve(ctx context.Context, in model.Match) error {

	query := `
		DELETE FROM matchs
		WHERE ((matchCatId = $1 AND userCatId = $2) OR (matchCatId = $2 AND userCatId = $1))
		AND id != $3;
	`

	_, err := r.db.ExecContext(ctx, query, in.MatchCatId, in.UserCatId, in.Id)
	if err != nil {
		return err
	}

	query = `
		UPDATE cats
		SET hasMatched = true
		WHERE id IN ($1, $2)
	`

	_, err = r.db.ExecContext(ctx, query, in.MatchCatId, in.UserCatId)
	if err != nil {
		return err
	}

	return nil
}

func (r *MatchRepository) FindMatchById(ctx context.Context, in model.RequestMatchApprove) (model.Match, error) {
	query := `
		SELECT id, matchCatId, userCatId, userId, targetUserId, message, createdAt
		FROM matchs
		WHERE id = $1
	`

	var match model.Match
	if err := r.db.QueryRowContext(ctx, query, in.MatchId).Scan(&match.Id, &match.MatchCatId, &match.UserCatId, &match.UserId, &match.TargetUserId, &match.Message, &match.CreatedAt); err != nil {
		return match, err
	}

	return match, nil
}
