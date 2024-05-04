package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Project-Sprint/w1-social-media/model"
)

type MatchService struct {
	matchRepository MatchRepository
}

var (
	ErrSameGender     = errors.New("same gender cat")
	ErrAlreadyMatched = errors.New("cat is taken/already matched")
	ErrSameOwner      = errors.New("same owner")
)

func NewMatchService(repo MatchRepository) *MatchService {
	return &MatchService{matchRepository: repo}
}

func (s *MatchService) PostMatch(ctx context.Context, body model.RequestMatch) error {

	// temp header, remove if auth already exists
	reqUserId := 3

	targetCat, err := s.matchRepository.FindCatById(ctx, body.MatchCatId)
	if err != nil {
		return err
	}

	ourCat, err := s.matchRepository.FindCatById(ctx, body.UserCatId)
	if err != nil {
		return err
	}

	switch {
	case ourCat.User_id != reqUserId:
		return sql.ErrNoRows

	case ourCat.Sex == targetCat.Sex:
		return ErrSameGender

	case ourCat.HasMatched && targetCat.HasMatched:
		return ErrAlreadyMatched

	case ourCat.User_id == targetCat.User_id:
		return ErrSameOwner
	}

	match := model.Match{
		MatchCatId:   targetCat.Id,
		UserCatId:    ourCat.Id,
		UserId:       ourCat.User_id,
		TargetUserId: targetCat.User_id,
		Message:      body.Message,
		CreatedAt:    time.Now(),
	}

	_, err = s.matchRepository.Insert(ctx, match)
	if err != nil {
		return err
	}

	return nil
}
