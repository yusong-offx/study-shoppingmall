package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/yusong-offx/myshoppingmall/docs"
	"github.com/yusong-offx/myshoppingmall/route/product"
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
	if err := utils.Connect(); err != nil {
		log.Fatal(err.Error())
	}
	defer utils.DB.Close()

	// Make CategoryMap
	if err := product.MakeCategoryMap(); err != nil {
		log.Fatal(err.Error())
	}

	// Middleware
	// app.Use(csrf.New())
	app.Use(cors.New())
	app.Use(recover.New())

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/clear/bad/good", func(c *fiber.Ctx) error {
		c.ClearCookie()
		return c.SendString("clear")
	})
	app.Get("/test123", product.CategorySearch)

	// Package
	// Base
	app.Get("/clearcookies", utils.CookiesClear)

	// User
	G_user := app.Group("/user")
	G_user.Post("/login", user.UserLogin)
	G_user.Post("/signup", user.UserSignUp)
	G_user.Post("/:id/withdraw", utils.JWTmiddleware, user.UserDelete)
	G_user.Get("/:id/info", utils.JWTmiddleware, user.UserGet)
	G_user.Get("/:id/logout", utils.JWTmiddleware, utils.Logout)

	// Vender
	G_vender := app.Group("/vender")
	G_vender.Post("/login", vender.VenderLogin)
	G_vender.Post("/signup", vender.VenderSignUp)
	G_vender.Post("/:id/withdraw", utils.JWTmiddleware, vender.UserDelete)
	G_vender.Get("/:id/info", utils.JWTmiddleware, vender.VenderGet)
	G_vender.Get("/:id/logout", utils.JWTmiddleware, utils.Logout)

	// Product
	G_product := app.Group("/product")
	G_product.Get("/item/:id", product.ItemGet)

	G_product.Post("/category/add", product.CategoryPost)
	G_product.Get("/category/refresh", product.CategoryRefresh)
	G_product.Get("/category/:cate", product.CategorySearch)

	// G_product.Post("/item/add", product.ItemPost)
	// G_product.Get("/", product.ItemGet)

	// Open server
	log.Fatal(app.Listen("localhost:8080"))

}
