package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusong-offx/myshoppingmall/database"
	"golang.org/x/crypto/bcrypt"
)

type Userinfo struct {
	Id           string `json:"id"`
	Password     []byte `json:"password"`
	Addr         string `json:"addr"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
}

type loginInfo struct {
	Id  string `json:"id"`
	Pwd string `json:"password"`
}

// Auth 구현
func LoginGet(c *fiber.Ctx) error {
	return c.SendString("temporary")
}

// @Tags         User
// @Summary      Login
// @Description  check ID and Password
// @Param        login_info body loginInfo true "annoying"
// @Router       /login [post]
func LoginPost(c *fiber.Ctx) error {
	idpwd := loginInfo{}
	err := c.BodyParser(&idpwd)
	if err != nil {
		return err
	}

	rows, err := database.DB.Query("select password from users where id=$1", idpwd.Id)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		var pwd []byte
		if err = rows.Scan(&pwd); err != nil {
			return err
		}
		if err = bcrypt.CompareHashAndPassword(pwd, []byte(idpwd.Pwd)); err != nil {
			return c.Status(401).SendString(err.Error())
		}
	}
	return c.Status(200).JSON("login success")
}
