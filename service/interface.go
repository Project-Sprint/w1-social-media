package service

import (
	"context"

	"github.com/Project-Sprint/w1-social-media/model"
)

type UserRepository interface {
	Insert(ctx context.Context, user model.User) error
	FindOne(ctx context.Context, user model.User) (model.User, error)
	FindById(ctx context.Context, user model.User) (model.User, error)
}

type MatchRepository interface {
	Insert(ctx context.Context, in model.Match) (interface{}, error)
	FindCatById(ctx context.Context, catId int) (model.Cats, error)
	Approve(ctx context.Context, in model.Match) error
	FindMatchById(ctx context.Context, in model.RequestMatchApprove) (model.Match, error)
}
