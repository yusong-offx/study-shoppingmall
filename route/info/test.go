package info

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/yusong-offx/myshoppingmall/database"
	"golang.org/x/crypto/bcrypt"
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
// @Param        login_info body IdPwd true "annoying"
// @Accept       json
// @Produce      json
// @Router       /test [post]
func Test(c *fiber.Ctx) error {
	idPwd := IdPwd{}
	err := c.BodyParser(&idPwd)
	if err != nil {
		return err
	}

	rows, err := DB.Query("select * from users where id=$1", idPwd.Id)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		user := Tests{}
		if err = rows.Scan(&user.Id, &user.Password, &user.Addr, &user.Phone_number, &user.Email); err != nil {
			return err
		}
		if err = bcrypt.CompareHashAndPassword(user.Password, []byte(idPwd.Pwd)); err != nil {
			return c.Status(401).SendString(err.Error())
		}
	}
	return c.Status(200).JSON("login")
}
