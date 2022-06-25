package product

import (
	"github.com/gofiber/fiber/v2"
	// . "github.com/yusong-offx/myshoppingmall/database"
)

type Item struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
	Price    int    `json:"price"`
	Content  string `json:"content"`
	Photo    []byte `json:"photo"`
	Vender   string `json:"vender"`
	Category string `json:"category"`
}

// @Tags		Product
// @Summary		Get Items Info
// @Param		name path string true "product name"
// @Param		stock path string true "product stock"
// @Router		/product/{name}/{stock} [get]
func ItemGet(c *fiber.Ctx) error {
	a := c.Params("name")
	b := c.Params("stock")
	return c.SendString(a + b + "Good")
}

// func ItemPost(c *fiber.Ctx) error {
// 	// Parse body json
// 	itemInfo := Item{}
// 	err := c.BodyParser(&itemInfo)

// 	rows, err := DB.Query("INSERT INTO items () VALUES ()")
// }
