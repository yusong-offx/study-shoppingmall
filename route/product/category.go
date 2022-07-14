package product

import (
	"container/list"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yusong-offx/myshoppingmall/utils"
)

type Category struct {
	Cur  string `json:"current cate"`
	Prev string `json:"previous cate"`
}

var CategoryMap = map[string]string{}
var CategoryTree = map[string][]string{}

func MakeCategoryMap() error {
	rows, err := utils.DB.Query("SELECT * FROM categories")
	if err != nil {
		return err
	}
	for rows.Next() {
		var cur, prev string
		err = rows.Scan(&cur, &prev)
		if err != nil {
			return err
		}
		CategoryMap[cur] = prev
		CategoryTree[prev] = append(CategoryTree[prev], cur)
	}
	return nil
}

func MakeCategorySlice(cur string) []string {
	categorySlice := []string{}
	for {
		categorySlice = append(categorySlice, cur)
		prev, ok := CategoryMap[cur]
		if !ok {
			break
		}
		cur = prev
	}
	return categorySlice
}

// @Tags		Category
// @Summary		refresh category
// @Description	remake category_map in server
// @Router		/product/category/refresh [get]
// @Security	ApiKeyAuth
// @Param		Authorization header string true "Authorization"
func CategoryRefresh(c *fiber.Ctx) error {
	for k := range CategoryMap {
		delete(CategoryMap, k)
	}
	if err := MakeCategoryMap(); err != nil {
		return utils.ErrorReqeustJSON(err, 500, c)
	}
	return c.Status(200).JSON(utils.RequestJSON{
		Success: true,
		Msg:     "Refresh Category",
	})
}

// @Tags		Category
// @Summary		add category
// @Description	add category
// @Param		Cate body Category true "prev - cur"
// @Router		/product/category/add [post]
// @Security	ApiKeyAuth
// @Param		Authorization header string true "Authorization"
func CategoryPost(c *fiber.Ctx) error {
	// Parse body (json)
	cate := Category{}
	err := c.BodyParser(&cate)
	if err != nil {
		return utils.ErrorReqeustJSON(err, 500, c)
	}

	// Insert DB (to lower case)
	rows, err := utils.DB.Query("INSERT INTO categories (cur, prev) VALUES ($1, $2)", strings.ToLower(cate.Cur), strings.ToLower(cate.Prev))
	if err != nil {
		return utils.ErrorReqeustJSON(err, 500, c)
	}
	defer rows.Close()

	return c.Status(201).JSON(utils.RequestJSON{
		Success: true,
		Msg:     "CREATE",
	})
}

func CategorySearch(c *fiber.Ctx) error {
	search_name := c.Params("cate")
	sqlConstrain := []string{search_name}
	q := list.New()
	q.PushBack(search_name)

	for e := q.Front(); e != nil; e = e.Next() {
		sliceV, ok := CategoryTree[e.Value.(string)]
		if ok {
			for _, v := range sliceV {
				q.PushBack(v)
				sqlConstrain = append(sqlConstrain, v)
			}
		}

	}

	rows, err := utils.DB.Query(fmt.Sprintf("SELECT * FROM items WHERE category IN('%s')", strings.Join(sqlConstrain, "', '")))
	if err != nil {
		return utils.ErrorReqeustJSON(err, 500, c)
	}
	datas := []Item{}
	for rows.Next() {
		e := Item{}
		err := rows.Scan(&e.Id, &e.Name, &e.Stock, &e.Price, &e.Content, &e.Photo, &e.Vender, &e.Category)
		if err != nil {
			return utils.ErrorReqeustJSON(err, 500, c)
		}
		datas = append(datas, e)
	}
	return c.Status(fiber.StatusOK).JSON(utils.RequestJSON{
		Success: true,
		Body:    datas,
	})
}
