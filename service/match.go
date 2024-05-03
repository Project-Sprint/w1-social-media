package service

import (
	"context"

	"github.com/Project-Sprint/w1-social-media/model"
)

type MatchService struct {
	matchRepository MatchRepository
}

func NewMatchService(repo MatchRepository) *MatchService {
	return &MatchService{matchRepository: repo}
}

func (s *MatchService) PostMatch(ctx context.Context, user model.User) error {
	return nil
}
