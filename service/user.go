package service

import (
	"context"
	"net/http"
	"time"

	"github.com/Project-Sprint/w1-social-media/helpers"
	"github.com/Project-Sprint/w1-social-media/model"
	"github.com/gofiber/fiber/v2"
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

func (s *UserService) Login(c *fiber.Ctx) error {
	var dataFromBody model.User
	c.BodyParser(&dataFromBody)

	if dataFromBody.Email == "" || dataFromBody.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(model.HTTPResponse{
			Code:    http.StatusBadRequest,
			Message: "",
			Data:    nil,
			Error:   "Fill All Field",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	s.userRepository.Insert(ctx, dataFromBody)
	user, err := s.userRepository.FindOne(ctx, dataFromBody)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.HTTPResponse{
			Code:    http.StatusBadRequest,
			Message: "",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	checkPassword := helpers.ComparePassword([]byte(user.Password), []byte(dataFromBody.Password))
	if !checkPassword {
		return c.Status(http.StatusBadRequest).JSON(model.HTTPResponse{
			Code:    http.StatusBadRequest,
			Message: "",
			Data:    nil,
			Error:   "invalid email/password",
		})
	}

	accessToken := helpers.SignToken(uint(user.Id))

	return c.Status(http.StatusOK).JSON(model.HTTPResponse{
		Code:    http.StatusOK,
		Message: "Success Login",
		Data: struct {
			AccessToken string     `json:"accessToken"`
			User        model.User `json:"user"`
		}{
			AccessToken: accessToken,
			User:        user,
		},
		Error: "",
	})
}
