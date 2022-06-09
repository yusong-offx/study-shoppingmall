package main

import (
	"log"

	"github.com/yusong-offx/myshoppingmall/db"

	"github.com/gofiber/fiber/v2"
)

type Info struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	Ip         string `json:"ip_address"`
}

func main() {
	app := fiber.New()

	db := db.Connect()
	defer db.Close()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Get("/info", func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT * FROM mock_data")
		if err != nil {
			c.Status(500)
			c.SendString(err.Error())
		}
		defer rows.Close()
		getData := []Info{}
		for rows.Next() {
			data := Info{}
			err := rows.Scan(&data.Id, &data.First_name, &data.Last_name, &data.Email, &data.Gender, &data.Ip)
			if err != nil {
				return err
			}
			getData = append(getData, data)
		}
		// b, _ := json.Unmarshal(a, &re)
		return c.JSON(getData)
	})

	app.Post("/info", func(c *fiber.Ctx) error {

	})
	log.Fatal(app.Listen(":8080"))
}
