package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusong-offx/myshoppingmall/database"
)

// @Tags 		User
// @Summary 	Get User Info
// @Description	Get User Info from DB
// @Param		id path string true "User id"
// @Router		/{id} [get]
func UserGet(c *fiber.Ctx) error {
	// Get data from DB (Userinfo struct define in login.go)
	rows, err := database.DB.Query("SELECT id, addr, phone_number, email FROM users WHERE id = $1", c.Params("id"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	data := Userinfo{}
	if rows.Next() {
		err = rows.Scan(&data.Id, &data.Addr, &data.Phone_number, &data.Email)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
	}
	return c.JSON(data)
}
