package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	info "github.com/yusong-offx/myshoppingmall/route/info"
)

func main() {
	app := fiber.New()

	// "/"
	local := app.Group("/")
	local.Get("/info", info.Info_Get)
	local.Post("/info", info.Info_Post)

	// Open server
	log.Fatal(app.Listen("localhost:8080"))
}
