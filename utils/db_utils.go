package utils

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type SqlFuncArgs struct {
	Rows     *sql.Rows
	Err      error
	Utilfunc func(*sql.Rows, ...any) any
	Args     any
	Ctx      *fiber.Ctx
}

func DBPost(args SqlFuncArgs) error {
	if args.Err != nil {
		return ErrorReqeustJSON(args.Err, 500, args.Ctx)
	}
	defer args.Rows.Close()

	return args.Ctx.Status(201).JSON(RequestJSON{
		Success: true,
		Msg:     args.Ctx.Method(),
	})
}

func DBGet(args SqlFuncArgs) (any, error) {
	if args.Err != nil {
		return nil, args.Err
	}
	defer args.Rows.Close()

	out := args.Utilfunc(args.Rows, args.Args)
	errOut, ok := out.(error)
	if ok {
		return nil, errOut
	}
	return out, nil
}

func DBDelete(args SqlFuncArgs) error {
	if args.Err != nil {
		return args.Err
	}
	v, ok := args.Args.([]string)
	if !ok {
		return errors.New("variable error")
	}
	rows, err := DB.Query(fmt.Sprintf("DELETE FROM %s WHERE %s = $1", v[0], v[1]), v[2])
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func DBUpdate(args SqlFuncArgs) error {
	if args.Err != nil {
		return args.Err
	}

	return nil
}
