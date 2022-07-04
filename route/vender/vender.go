package vender

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusong-offx/myshoppingmall/utils"
	"golang.org/x/crypto/bcrypt"
)

type Vender struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	Addr         string `json:"addr"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
	Url          string `json:"url"`
}

// func VenderGet(c *fiber.Ctx) error {
// 	rows, err := database.DB.Query("SELECT * FROM venders where id = $1", c.Params())
// 	c.Qu
// }

// func VenderPost(c *fiber.Ctx) error {
// 	err := database.DB.Query("INSERT INTO venders VALUES ($1, $2, $3, $4, $5, $6)")
// }

// @Tags		Vender
// @Summary 	vender sign up
// @Description	vender sign up
// @Param		INFO body Vender true "vender info"
// @Router		/vender/signup [post]
func VenderSignUp(c *fiber.Ctx) error {
	// Parse body(json)
	vender := &Vender{}
	err := c.BodyParser(vender)
	if err != nil {
		return utils.ErrorReqeustJSON(err, 500, c)
	}

	// Bcrypt password
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(vender.Password), 10)
	if err != nil {
		return utils.ErrorReqeustJSON(err, 500, c)
	}

	// Post Data
	rows, err := utils.DB.Query(
		"INSERT INTO venders VALUES ($1, $2, $3, $4, $5, $6)",
		vender.Id, vender.Name, pwdHash, vender.Addr, vender.Phone_number, vender.Email)
	return utils.DBPost(rows, err, c)
}

// @Tags         Vender
// @Summary      Login
// @Description  Check ID and Password
// @Description  Password secure by bcrypt
// @Param        ID&PWD body utils.LoginInfo true "ID and Password"
// @Router       /vender/login [post]
func VenderLogin(c *fiber.Ctx) error {
	return utils.Login("venders", c)
}
