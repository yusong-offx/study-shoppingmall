package info

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yusong-offx/myshoppingmall/database"
)

// Database data form
// @Discription Info struct
type Info struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	Ip_address string `json:"ip_address"`
}

func Get(c *fiber.Ctx) error {
	db := database.DB
	// Command sql
	rows, err := db.Query("SELECT * FROM mock_data")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	// Get data from database
	getData := struct {
		Datas []Info `json:"infos"`
	}{}
	for rows.Next() {
		data := Info{}
		err := rows.Scan(&data.Id, &data.First_name, &data.Last_name, &data.Email, &data.Gender, &data.Ip_address)
		if err != nil {
			return err
		}
		getData.Datas = append(getData.Datas, data)
	}

	// Return data by json
	return c.JSON(getData)
}

func Post(c *fiber.Ctx) error {
	db := database.DB

	// Body parse into struct
	data := &Info{}
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Insert data into database
	res, err := db.Query(
		"INSERT INTO mock_data values ($1, $2, $3, $4, $5, $6)",
		data.Id,
		data.First_name,
		data.Last_name,
		data.Email,
		data.Gender,
		data.Ip_address,
	)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer res.Close()

	return c.Status(201).JSON(data)
}

func Delete(c *fiber.Ctx) error {
	db := database.DB

	// Body parse into struct
	data := &Info{}
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Command sql
	res, err := db.Query("DELETE FROM mock_data Where id = $1", data.Id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer res.Close()
	return c.JSON("Deleted")
}
