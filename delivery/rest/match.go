package rest

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/Project-Sprint/w1-social-media/model"
	"github.com/Project-Sprint/w1-social-media/service"
	"github.com/gofiber/fiber/v2"
)

func PostMatchHandler(matchInterface MatchService) fiber.Handler {

	return func(c *fiber.Ctx) error {
		var body model.RequestMatch
		c.BodyParser(&body)

		if body.Message == "" {
			return c.Status(http.StatusBadRequest).JSON(model.HTTPResponse{
				Message: "message should be filled",
				Data:    nil,
			})
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := matchInterface.PostMatch(ctx, body)
		if err != nil {
			var (
				errResp    = errors.New("internal server error")
				resp       model.HTTPResponse
				statusCode = 500
			)

			switch {
			case err == sql.ErrNoRows:
				errResp = errors.New("data not found")
				statusCode = 404

			case err == service.ErrSameGender || err == service.ErrAlreadyMatched || err == service.ErrSameOwner:
				errResp = err
				statusCode = 400
			}

			resp = model.HTTPResponse{
				Message: errResp.Error(),
				Data:    nil,
			}

			return c.Status(statusCode).JSON(resp)
		}

		return c.Status(http.StatusCreated).JSON(model.HTTPResponse{
			Message: "success matching",
			Data:    nil,
		})
	}
}

func ApproveMatchHandler(matchInterface MatchService) fiber.Handler {

	return func(c *fiber.Ctx) error {
		var body model.RequestMatchApprove
		c.BodyParser(&body)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := matchInterface.ApproveMatch(ctx, body); err != nil {

			if err == sql.ErrNoRows {
				return c.Status(http.StatusNotFound).JSON(model.HTTPResponse{
					Message: err.Error(),
					Data:    nil,
				})
			}

			return c.Status(http.StatusInternalServerError).JSON(model.HTTPResponse{
				Message: err.Error(),
				Data:    nil,
			})
		}

		return c.Status(http.StatusOK).JSON(model.HTTPResponse{
			Message: "success approving",
			Data:    nil,
		})
	}
}

func RejectMatchHandler(matchInterface MatchService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, Match Reject!")
	}
}
