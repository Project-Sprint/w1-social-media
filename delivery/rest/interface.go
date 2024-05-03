package rest

import (
	"context"

	"github.com/Project-Sprint/w1-social-media/model"
)

type UserService interface {
	Register(ctx context.Context, user model.User) error
}

type MatchService interface{}
