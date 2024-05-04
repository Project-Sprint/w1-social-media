package service

import (
	"context"

	"github.com/Project-Sprint/w1-social-media/model"
)

type CatService struct {
	catRepository CatRepository
}

func NewCatService(repo CatRepository) *CatService {

	return &CatService{catRepository: repo}
}

func (s *CatService) PostCat(ctx context.Context, cat model.Cats) error {

	_, err := s.catRepository.Insert(ctx, cat)

	if err != nil {
		return err
	}

	return nil
}