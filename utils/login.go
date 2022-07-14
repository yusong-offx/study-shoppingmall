package utils

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginInfo struct {
	Id  string `json:"id"`
	Pwd string `json:"password"`
}

func LoginGetfunc(rows *sql.Rows, args ...any) any {
	var hashed []byte
	login, _ := args[0].(LoginInfo)
	if rows.Next() {
		if err := rows.Scan(&hashed); err != nil {
			return err
		}
		if err := bcrypt.CompareHashAndPassword(hashed, []byte(login.Pwd)); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("NO USER (ID : %s )", login.Id)
	}
	return hashed
}

func Login(table string, c *fiber.Ctx) error {
	// Parse body (json)
	login := LoginInfo{}
	err := c.BodyParser(&login)
	if err != nil {
		return ErrorReqeustJSON(err, 500, c)
	}

	// Get password and compare
	sqlCMD := SqlFuncArgs{
		Utilfunc: LoginGetfunc,
		Args:     login,
		Ctx:      c,
	}
	sqlCMD.Rows, sqlCMD.Err = DB.Query("SELECT password FROM "+table+" WHERE id = $1", login.Id)
	_, err = DBGet(sqlCMD)
	if err != nil {
		return ErrorReqeustJSON(err, 500, c)
	}

	// Gernerate jwt
	expiretime := time.Now().Add(30 * time.Minute)
	token, err := GenerateToken(AuthClaims{
		Id:        login.Id,
		UserType:  table,
		ExpiredAt: expiretime.Unix(),
		Uuid:      uuid.Must(uuid.NewRandom()),
	})
	if err != nil {
		return ErrorReqeustJSON(err, 500, c)
	}
	c.Cookie(&fiber.Cookie{
		Name:     "Authorization",
		Value:    token,
		Expires:  expiretime,
		HTTPOnly: true,
		SameSite: "lax",
	})

	return c.Status(200).JSON(RequestJSON{
		Success: true,
		Msg:     "LOGIN SUCCESS",
	})
}

func CookiesClear(c *fiber.Ctx) error {
	c.ClearCookie()
	return c.Status(fiber.StatusOK).JSON(RequestJSON{
		Success: true,
		Msg:     "Successfully logout",
	})
}

func Logout(c *fiber.Ctx) error {
	return c.Redirect("/clearcookies")
}
