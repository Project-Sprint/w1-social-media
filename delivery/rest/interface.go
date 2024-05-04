package rest

import (
	"context"

	"github.com/Project-Sprint/w1-social-media/model"
)

type UserService interface {
	Register(ctx context.Context, user model.User) error
}

type CatService interface {
	PostCat(ctx context.Context, body model.Cats) error
}

type MatchService interface{}
