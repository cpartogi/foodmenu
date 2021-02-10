package store

import (
	"context"

	"github.com/cpartogi/foodmenu/constant"
	"github.com/cpartogi/foodmenu/schema/response"
)

const getMenuType = `-- name: MenuType
SELECT menu_type_id, menu_type_name FROM tb_menu_type ORDER BY menu_type_name
`

// Deposit balancelog is
func (q *Queries) MenuType(ctx context.Context) ([]response.MenuType, error) {
	rows, err := q.db.QueryContext(ctx, getMenuType)

	defer rows.Close()

	var y []response.MenuType
	var i response.MenuType

	c := 0
	for rows.Next() {
		_ = rows.Scan(
			&i.MenuTypeId,
			&i.MenuTypeName,
		)
		y = append(y, i)
		c++
	}

	//return not found
	if c == 0 {
		err = constant.ErrNotFound
	}

	return y, err
}
