package login

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/yusong-offx/myshoppingmall/database"
	"golang.org/x/crypto/bcrypt"
)

type userinfo struct {
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

// @Tags         login
// @Summary      login
// @Description  check id and password
// @Param        login_info body loginInfo true "annoying"
// @Accept       json
// @Produce      json
// @Router       /login [post]
func Post(c *fiber.Ctx) error {
	idpwd := loginInfo{}
	err := c.BodyParser(&idpwd)
	if err != nil {
		return err
	}

	rows, err := DB.Query("select * from users where id=$1", idpwd.Id)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		user := userinfo{}
		if err = rows.Scan(&user.Id, &user.Password, &user.Addr, &user.Phone_number, &user.Email); err != nil {
			return err
		}
		if err = bcrypt.CompareHashAndPassword(user.Password, []byte(idpwd.Pwd)); err != nil {
			return c.Status(401).SendString(err.Error())
		}
	}
	return c.Status(200).JSON("login")
}
