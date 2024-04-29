package service

import (
	"context"

	"github.com/Project-Sprint/w1-social-media/model"
)

type UserRepository interface {
	Insert(ctx context.Context, user model.User) error
}