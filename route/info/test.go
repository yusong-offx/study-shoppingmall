package info

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
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
// @Router       /test [get]
func Test(c *fiber.Ctx) error {
	ck := new(fiber.Cookie)
	ck.Name = "my cookies"
	ck.Value = "answer"
	ck.Expires = time.Now().Add(time.Second)
	c.Cookie(ck)
	fmt.Println(c.Cookies("csrf_"), c.Cookies("my cookies"))
	return c.SendString("hello World!")
}

func Test1(c *fiber.Ctx) error {
	time.Sleep(3 * time.Second)
	return c.SendString("error")
}
