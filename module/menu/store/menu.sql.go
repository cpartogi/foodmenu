package store

import (
	"context"
	"database/sql"

	"github.com/cpartogi/foodmenu/constant"
	"github.com/cpartogi/foodmenu/schema/request"
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

const addMenu = `-- name: AddMenu :one
INSERT INTO tb_menu (
	menu_id,
    menu_type_id,
    warteg_id,
    menu_name,
    menu_detail,
    menu_picture,
	menu_price
) VALUES (
    uuid(),
    ?,
    ?,
    ?,
	?,
	?,
	?
)
`

func (q *Queries) MenuAdd(ctx context.Context, addm request.Menu) (mn response.MenuAdd, err error) {
	result, err := q.db.ExecContext(ctx, addMenu,
		addm.MenuTypeId,
		addm.WartegId,
		addm.MenuName,
		addm.MenuDetail,
		addm.MenuPicture,
		addm.MenuPrice,
	)

	if err != nil {
		return
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return
	}

	if rows != 1 {
		return
	}

	i := response.MenuAdd{
		MenuTypeId:  addm.MenuTypeId,
		WartegId:    addm.WartegId,
		MenuName:    addm.MenuName,
		MenuDetail:  addm.MenuDetail,
		MenuPicture: addm.MenuPicture,
		MenuPrice:   addm.MenuPrice,
	}

	return i, err
}

const deleteMenu = `-- name: DeleteMenu :one
DELETE FROM tb_menu WHERE menu_id = ?
`

func (q *Queries) MenuDelete(ctx context.Context, menu_id string) (md response.MenuDelete, err error) {
	result, err := q.db.ExecContext(ctx, deleteMenu, menu_id)

	if err != nil {
		return
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return
	}

	if rows != 1 {
		err = constant.ErrNotFound
	}

	i := response.MenuDelete{
		MenuId: menu_id,
	}

	return i, err
}

const updateMenu = `-- name: UpdateMenu :one
UPDATE tb_menu SET menu_type_id=?, warteg_id=?, menu_name=?, menu_detail=?, menu_picture=?, menu_price=?, updated_date=CURRENT_TIMESTAMP(3) WHERE menu_id = ?
`

func (q *Queries) MenuUpdate(ctx context.Context, menu_id string, upm request.MenuUpdate) (mu response.MenuUpdate, err error) {
	result, err := q.db.ExecContext(ctx, updateMenu,
		upm.MenuTypeId,
		upm.WartegId,
		upm.MenuName,
		upm.MenuDetail,
		upm.MenuPicture,
		upm.MenuPrice,
		menu_id,
	)

	if err != nil {
		return
	}

	rows, err := result.RowsAffected()

	if rows != 1 {
		err = constant.ErrNotFound
	}

	i := response.MenuUpdate{
		MenuId:      menu_id,
		MenuTypeId:  upm.MenuTypeId,
		WartegId:    upm.WartegId,
		MenuName:    upm.MenuName,
		MenuDetail:  upm.MenuDetail,
		MenuPicture: upm.MenuPicture,
		MenuPrice:   upm.MenuPrice,
	}

	return i, err
}

const listMenu = `-- name: ListMenu
SELECT a.menu_type_name, b.menu_name, b.menu_detail, b.menu_picture, b.menu_price FROM tb_menu_type a, tb_menu b WHERE a.menu_type_id=b.menu_type_id AND b.warteg_id=? AND b.menu_type_id=? ORDER BY b.menu_name
`

func (q *Queries) MenuList(ctx context.Context, warteg_id string, menu_type_id int) (list []response.MenuList, err error) {
	rows, err := q.db.QueryContext(ctx, listMenu, warteg_id, menu_type_id)

	if err != nil {
		return
	}

	defer rows.Close()

	var y []response.MenuList
	var i response.MenuList

	c := 0

	for rows.Next() {
		_ = rows.Scan(
			&i.MenuTypeName,
			&i.MenuName,
			&i.MenuDetail,
			&i.MenuPicture,
			&i.MenuPrice,
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

const getMenuDetail = `-- name: MenuDetail :one
SELECT b.menu_id, a.menu_type_name, b.warteg_id, b.menu_name, b.menu_detail, b.menu_picture, b.menu_price FROM tb_menu_type a, tb_menu b
WHERE a.menu_type_id=b.menu_type_id AND b.menu_id = ?
`

func (q *Queries) MenuDetail(ctx context.Context, menu_id string) (mnd response.MenuDetail, err error) {
	row := q.db.QueryRowContext(ctx, getMenuDetail, menu_id)
	var i response.MenuDetail
	err = row.Scan(
		&i.MenuId,
		&i.MenuTypeName,
		&i.WartegId,
		&i.MenuName,
		&i.MenuDetail,
		&i.MenuPicture,
		&i.MenuPrice,
	)

	if err == sql.ErrNoRows {
		err = constant.ErrNotFound
	}

	return i, err
}
