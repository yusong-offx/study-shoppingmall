package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/yusong-offx/myshoppingmall/database"
	_ "github.com/yusong-offx/myshoppingmall/docs"
	"github.com/yusong-offx/myshoppingmall/route/info"
	"github.com/yusong-offx/myshoppingmall/route/login"
	"github.com/yusong-offx/myshoppingmall/route/signup"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()

	// Allow CORS
	app.Use(cors.New())

	// Database connect
	err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.DB.Close()

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Post("/test", info.Test)

	// Path
	app.Post("/login", login.Post)
	app.Post("/signup", signup.Post)

	// Open server
	log.Fatal(app.Listen("localhost:8080"))
}
