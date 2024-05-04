package rest

import (
	"net/http"
	"time"

	"github.com/Project-Sprint/w1-social-media/model"
	"github.com/gofiber/fiber/v2"
)

func PostCatHandler(catInterface CatService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		
		var body model.PostCat
		c.BodyParser(&body)
		

		if body.Name == "" || body.Race == "" || body.Sex == "" || body.AgeInMonth <= 0 || body.Description == "" || len(body.ImageUrls) == 0  {

			return c.Status(http.StatusBadRequest).JSON(model.HTTPResponse{
				Message: "request doesn’t pass validation",
				Data:    nil,
			})
		}



		if len(body.Name) == 0 || len(body.Name) > 30 || 
		(body.Sex == "male" || body.Sex == "female") || 
		body.AgeInMonth <= 1 || body.AgeInMonth > 120082 || len(body.Description) <= 1 || len(body.Description) > 200 {
			return c.Status(http.StatusBadRequest).JSON(model.HTTPResponse{
				Message: "request doesn’t pass validation",
				Data:    nil,
			})
		}

		return c.Status(http.StatusCreated).JSON(model.HTTPResponse{
			Message: "success",
			Data: struct {
			Id int `json:"id"`
			CreatedAt time.Time `json:"createdAt"`
			}{
				Id: 1,
				CreatedAt: time.Now(),
			},
		})
	}
}

func GetCatHandler(catInterface CatService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, Get Cat!")
	}
}

func PutCatHandler(catInterface CatService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, Cat Put!")
	}
}

func DeleteCatHandler(catInterface CatService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, Cat Delete!")
	}
}
