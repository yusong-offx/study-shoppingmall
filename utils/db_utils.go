package utils

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type SqlFuncArgs struct {
	Rows     *sql.Rows
	Err      error
	Utilfunc func(*sql.Rows, ...any) any
	Args     any
	Ctx      *fiber.Ctx
}

func DBPost(rows *sql.Rows, err error, c *fiber.Ctx) error {
	if err != nil {
		return ErrorReqeustJSON(err, 500, c)
	}
	defer rows.Close()

	return c.Status(201).JSON(RequestJSON{
		Success: true,
		Msg:     c.Method(),
	})
}

func DBGet(args SqlFuncArgs) (any, error) {
	if args.Err != nil {
		return nil, args.Err
	}
	defer args.Rows.Close()

	out := args.Utilfunc(args.Rows, args.Args)
	wrapOut, ok := out.(error)
	if ok {
		return nil, wrapOut
	}
	return out, nil
}
