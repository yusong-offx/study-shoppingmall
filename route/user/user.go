package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusong-offx/myshoppingmall/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           string `json:"id"`
	Password     string `json:"password"`
	Addr         string `json:"addr"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
}

type UserInfo struct {
	Id           string `json:"id"`
	Addr         string `json:"addr"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
}

// @Tags 		User
// @Summary 	Get User Info
// @Description	Get User Info from DB
// @Param		id path string true "User id"
// @Security	ApiKeyAuth
// @param		Authorization header string true "Authorization"
// @Router		/user/info/{id} [get]
func UserGet(c *fiber.Ctx) error {
	// Get data from DB
	rows, err := utils.DB.Query("SELECT id, addr, phone_number, email FROM users WHERE id = $1", c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(utils.RequestJSON{
			Success: false,
			Msg:     err.Error(),
		})
	}
	data := UserInfo{}
	if rows.Next() {
		err = rows.Scan(&data.Id, &data.Addr, &data.Phone_number, &data.Email)
		if err != nil {
			return c.Status(500).JSON(utils.RequestJSON{
				Success: false,
				Msg:     err.Error(),
			})
		}
	} else {
		return c.SendStatus(fiber.StatusNoContent)
	}
	return c.JSON(utils.RequestJSON{
		Success: true,
		Body:    data,
	})
}

// @Tags			User
// @Summary			Sign up
// @Description		Post user info for sign up into DB
// @Param			INFO body User true "User infomation"
// @Router			/user/signup [post]
func UserSignUp(c *fiber.Ctx) error {
	// Parse body(json)
	user := &User{}
	err := c.BodyParser(user)
	if err != nil {
		return utils.ErrorReqeustJSON(err, 500, c)
	}

	// Bcrypt password
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return utils.ErrorReqeustJSON(err, 500, c)
	}

	// Post Data
	rows, err := utils.DB.Query(
		"INSERT INTO users VALUES ($1, $2, $3, $4, $5)",
		user.Id, pwdHash, user.Addr, user.Phone_number, user.Email)
	return utils.DBPost(rows, err, c)
}

// @Tags         User
// @Summary      Login
// @Description  Check ID and Password
// @Description  Password secure by bcrypt
// @Param        ID&PWD body utils.LoginInfo true "ID and Password"
// @Router       /user/login [post]
func UserLogin(c *fiber.Ctx) error {
	return utils.Login("users", c)
}
