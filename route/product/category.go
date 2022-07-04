package product

import (
	"github.com/yusong-offx/myshoppingmall/database"
)

var CategoryMap map[string]string

func MakeCategoryMap() error {
	rows, err := database.DB.Query("SELECT * FROM categories")
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
	}
	return nil
}

func MakeCategorySlice(cur string) []string {
	categorySlice := []string{}
	categorySlice = append(categorySlice, cur)
	for {
		cur, ok := CategoryMap[cur]
		if !ok {
			break
		}
		categorySlice = append(categorySlice, cur)
	}
	return categorySlice
}
