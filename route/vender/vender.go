package vender

import (
	"fmt"

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

type VenderInfo struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Addr         string `json:"addr"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
	Url          string `json:"url"`
}

// @Tags		Vender
// @Summary get	vender info
// @Description	get vender info
// @Parmas		ID
// @Router		/vender/{id}/info [get]
// @Security	ApiKeyAuth
// @param		Authorization header string true "Authorization"
func VenderGet(c *fiber.Ctx) error {
	rows, err := utils.DB.Query("SELECT id, name, addr, phone_number, email, url FROM venders WHERE id = $1", c.Params("id"))
	if err != nil {
		return utils.ErrorReqeustJSON(err, 500, c)
	}
	ret := VenderInfo{}
	if rows.Next() {
		err = rows.Scan(&ret.Id, &ret.Name, &ret.Addr, &ret.Phone_number, &ret.Email, &ret.Url)
		if err != nil {
			return utils.ErrorReqeustJSON(err, 500, c)
		}
	} else {
		return utils.ErrorReqeustJSON(fmt.Errorf("no user (id : %s)", c.Params("id")), 403, c)
	}
	return c.Status(200).JSON(ret)
}

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
	return utils.DBPost(utils.SqlFuncArgs{
		Rows: rows,
		Err:  err,
		Ctx:  c,
	})
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

func UserDelete(c *fiber.Ctx) error {
	if err := utils.DBDelete(utils.SqlFuncArgs{
		Args: []string{"venders", "id", c.Params("id")},
	}); err != nil {
		return utils.ErrorReqeustJSON(err, 500, c)
	}
	return c.Status(fiber.StatusOK).Redirect("/clearcookies")
}
