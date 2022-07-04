package info

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yusong-offx/myshoppingmall/route"
)

type Tests struct {
	Id           string `json:"id"`
	Password     []byte `json:"password"`
	Addr         string `json:"addr"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
}
type IdPwd struct {
	Id  string `json:"id"`
	Pwd string `json:"password"`
}

// @Tags         Test
// @Summary      Test List
// @Description  Test get list
// @Produce      json
// @Router       /test [get]
func Test(c *fiber.Ctx) error {
	fmt.Println(c.GetReqHeaders()["Authorization"])
	fmt.Println()
	fmt.Println(c.Cookies("Authorization"))
	fmt.Println()
	fmt.Println(c.GetRespHeaders())

	return c.JSON(route.RequestJSON{
		Success: false,
		Msg:     "godd",
	})
}

// @Tags         Test
// @Summary      Test List
// @Description  Test get list
// @Router       /test1 [get]
func Test1(c *fiber.Ctx) error {
	fmt.Println(c.IPs(), c.IP())
	return c.SendString("good")
}
