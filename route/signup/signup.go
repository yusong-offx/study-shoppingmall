package signup

import (
	"github.com/gofiber/fiber/v2"
	. "github.com/yusong-offx/myshoppingmall/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           string `json:"id"`
	Password     string `json:"password"`
	Addr         string `json:"addr"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
}

// @Tags			SignUp
// @Summary			Sign up
// @Description		Post user info
// @Param			Info body User true "User infomation"
// @Router			/signup [post]
func Post(c *fiber.Ctx) error {
	// Parse body(json)
	user := &User{}
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Bcrypt password
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Post Data
	rows, err := DB.Query("INSERT INTO users VALUES ($1, $2, $3, $4, $5)", user.Id, pwdHash, user.Addr, user.Phone_number, user.Email)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	return c.Status(201).JSON("Created")
}
