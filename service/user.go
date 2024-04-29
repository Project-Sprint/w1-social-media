package service

import (
	"context"

	"github.com/Project-Sprint/w1-social-media/model"
)

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (s *UserService) Register(ctx context.Context, user model.User) error {
	return nil
}
