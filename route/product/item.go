package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusong-offx/myshoppingmall/utils"
)

// . "github.com/yusong-offx/myshoppingmall/database"

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
func ItemPost(c *fiber.Ctx) error {
	// Body parse(json)
	item := Item{}
	err := c.BodyParser(&item)
	if err != nil {
		return utils.ErrorReqeustJSON(err, 500, c)
	}
	rows, err := utils.DB.Query(
		"INSERT INTO name, stock, price, content, photo, vender, category VALUES ($1, $2, $3, $4, $5, $6, $7)",
		item.Name, item.Stock, item.Price, item.Content, item.Photo, item.Vender, item.Category)
	if err != nil {
		return utils.ErrorReqeustJSON(err, 500, c)
	}
	defer rows.Close()

	return c.Status(fiber.StatusCreated).JSON(utils.RequestJSON{
		Success: true,
		Msg:     "created",
	})
}

// @Tags		Product
// @Summary		Get Items Info
// @Param		name path string true "product name"
// @Param		stock path string true "product stock"
// @Router		/product/item/add [post]
func ItemGet(c *fiber.Ctx) error {
	// Parse body json
	item := Item{}

	rows, err := utils.DB.Query("INSERT ")
}
