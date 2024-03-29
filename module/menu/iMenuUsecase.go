package menu

import (
	"context"

	"github.com/cpartogi/foodmenu/schema/request"
	"github.com/cpartogi/foodmenu/schema/response"
)

// Usecase is
type Usecase interface {
	MenuType(ctx context.Context) (mt []response.MenuType, err error)
	MenuAdd(ctx context.Context, addm request.Menu) (mn response.MenuAdd, err error)
	MenuDelete(ctx context.Context, menu_id string) (md response.MenuDelete, err error)
	MenuUpdate(ctx context.Context, menu_id string, upm request.MenuUpdate) (mu response.MenuUpdate, err error)
	MenuList(ctx context.Context, warteg_id, menu_type_id, menu_name string) (list []response.MenuList, err error)
	MenuDetail(ctx context.Context, menu_id string) (mnd response.MenuDetail, err error)
}
