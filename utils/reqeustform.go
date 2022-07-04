package utils

import "github.com/gofiber/fiber/v2"

type RequestJSON struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Body    interface{} `json:"body"`
}

func ErrorReqeustJSON(err error, httpcode int, c *fiber.Ctx) error {
	return c.Status(httpcode).JSON(RequestJSON{
		Success: false,
		Msg:     err.Error(),
	})
}
