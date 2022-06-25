package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/yusong-offx/myshoppingmall/database"
	_ "github.com/yusong-offx/myshoppingmall/docs"

	"github.com/yusong-offx/myshoppingmall/route/info"
	"github.com/yusong-offx/myshoppingmall/route/product"
	"github.com/yusong-offx/myshoppingmall/route/user"
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
	// Fiber start
	app := fiber.New()

	// Database connect
	err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.DB.Close()

	// Middleware
	app.Use(cors.New())
	app.Use(csrf.New())
	app.Use(recover.New())

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/test", info.Test)
	app.Get("/test1", info.Test1)

	// Package
	// User
	app.Post("/login", user.LoginPost)
	app.Post("/signup", user.SignUpPost)
	app.Get("/:id", user.UserGet)

	// Product
	app.Get("/product/:name/:stock", product.ItemGet)

	// Open server
	log.Fatal(app.Listen("localhost:8080"))
}
