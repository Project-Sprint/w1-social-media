package repository

import (
	"context"
	"database/sql"
	"errors"

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

func (r *UserRepository) FindOne(ctx context.Context, in model.User) (model.User, error) {
	query := `
		SELECT id, name, role, email, password
		FROM "user"
		WHERE email = $1
	`

	var user model.User

	err := r.db.QueryRowContext(ctx, query, in.Email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return user, errors.New("invalid email/password")
	}

	return user, nil
}

func (r *UserRepository) FindById(ctx context.Context, in model.User) (model.User, error) {
	query := `
		SELECT id, name, email, password
		FROM "user"
		WHERE id = $1
	`

	var user model.User

	err := r.db.QueryRowContext(ctx, query, in.Id).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
