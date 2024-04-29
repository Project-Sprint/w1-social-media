package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Project-Sprint/w1-social-media/common"
	"github.com/Project-Sprint/w1-social-media/delivery/rest"
	"github.com/Project-Sprint/w1-social-media/repository"
	"github.com/Project-Sprint/w1-social-media/service"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var PORT string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	PORT = os.Getenv("SERVER_PORT")
	if PORT == "" {
		PORT = "8080"
	}

}

func main() {
	db := common.ConnectToSQL()

	// define repository
	userRepository := repository.NewUserRepository(db)

	// define service
	userService := service.NewUserService(*userRepository)

	app := fiber.New()
	apiV1 := app.Group("/api/v1")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// api/v1/user handler
	userRoutes := apiV1.Group("/user")
	userRoutes.Post("/register", rest.UserRegisterHandler(userService))

	fmt.Println("Listening on Port: ", PORT)
	if err := app.Listen(fmt.Sprintf(":%s", PORT)); err != nil {
		log.Fatalf("Failed Listening Port: %s\n", PORT)
	}
}
