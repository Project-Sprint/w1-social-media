package middleware

import (
	"context"
	"net/http"

	"github.com/Project-Sprint/w1-social-media/common"
	"github.com/Project-Sprint/w1-social-media/helpers"
	"github.com/Project-Sprint/w1-social-media/model"
	"github.com/Project-Sprint/w1-social-media/repository"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	checkToken, err := helpers.VerifyToken(authHeader)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(model.HTTPResponse{
			Code:    http.StatusUnauthorized,
			Message: "",
			Data:    nil,
			Error:   "Invalid Token",
		})
	}

	var tempUser model.User

	idFloat, ok := checkToken["id"].(float64)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(model.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: "",
			Data:    nil,
			Error:   "Error Converting User Id",
		})
	}

	tempUser.Id = int(idFloat)

	db := common.ConnectToSQL()

	// define repository
	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.FindById(context.Background(), tempUser)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(model.HTTPResponse{
			Code:    http.StatusUnauthorized,
			Message: "",
			Data:    nil,
			Error:   "Invalid Token",
		})
	}

	c.Locals("user", user) //set header

	return c.Next()
}
