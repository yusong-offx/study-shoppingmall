package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/yusong-offx/myshoppingmall/docs"
	"github.com/yusong-offx/myshoppingmall/route/user"
	"github.com/yusong-offx/myshoppingmall/route/vender"
	"github.com/yusong-offx/myshoppingmall/utils"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description OAuth protects our entity endpoints

// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	// Fiber start
	app := fiber.New()

	// Database connect
	err := utils.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer utils.DB.Close()

	// Middleware
	// app.Use(csrf.New())
	app.Use(cors.New())
	app.Use(recover.New())

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Package
	// User
	G_user := app.Group("/user")
	G_user.Post("/login", user.UserLogin)
	G_user.Post("/signup", user.UserSignUp)
	G_user.Get("/info/:id", utils.JWTmiddleware, user.UserGet)

	// Vender
	G_vender := app.Group("/vender")
	G_vender.Post("/login", vender.VenderLogin)
	G_vender.Post("/signup", vender.VenderSignUp)

	// Open server
	log.Fatal(app.Listen("localhost:8080"))

}
