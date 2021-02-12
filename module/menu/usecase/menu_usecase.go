package usecase

import (
	"context"
	"time"

	"github.com/cpartogi/foodmenu/module/menu"
	"github.com/cpartogi/foodmenu/schema/request"
	"github.com/cpartogi/foodmenu/schema/response"
)

// AuthUsecase will create a usecase with its required repo
type MenuUsecase struct {
	menuRepo       menu.Repository
	contextTimeout time.Duration
}

// NewAuthUsecase will create new an contactUsecase object representation of auth.Usecase
func NewMenuUsecase(ar menu.Repository, timeout time.Duration) menu.Usecase {
	return &MenuUsecase{
		menuRepo:       ar,
		contextTimeout: timeout,
	}
}

func (u *MenuUsecase) MenuType(ctx context.Context) (dis []response.MenuType, err error) {
	resp := []response.MenuType{}

	mtype, err := u.menuRepo.MenuType(ctx)

	if err != nil {
		return resp, err
	}

	return mtype, err
}

func (u *MenuUsecase) MenuAdd(ctx context.Context, addm request.Menu) (mn response.MenuAdd, err error) {
	resp := response.MenuAdd{
		MenuTypeId:  addm.MenuTypeId,
		WartegId:    addm.WartegId,
		MenuName:    addm.MenuName,
		MenuDetail:  addm.MenuDetail,
		MenuPicture: addm.MenuPicture,
		MenuPrice:   addm.MenuPrice,
	}

	req := request.Menu{
		MenuTypeId:  addm.MenuTypeId,
		WartegId:    addm.WartegId,
		MenuName:    addm.MenuName,
		MenuDetail:  addm.MenuDetail,
		MenuPicture: addm.MenuPicture,
		MenuPrice:   addm.MenuPrice,
	}

	addmenu, err := u.menuRepo.MenuAdd(ctx, req)

	if err != nil {
		return resp, err
	}
	return addmenu, err
}

func (u *MenuUsecase) MenuDelete(ctx context.Context, menu_id string) (md response.MenuDelete, err error) {
	resp := response.MenuDelete{
		MenuId: menu_id,
	}

	delmenu, err := u.menuRepo.MenuDelete(ctx, menu_id)
	if err != nil {
		return resp, err
	}

	return delmenu, err
}

func (u *MenuUsecase) MenuUpdate(ctx context.Context, menu_id string, upm request.MenuUpdate) (mu response.MenuUpdate, err error) {
	resp := response.MenuUpdate{
		MenuId:      menu_id,
		MenuTypeId:  upm.MenuTypeId,
		WartegId:    upm.WartegId,
		MenuName:    upm.MenuName,
		MenuDetail:  upm.MenuDetail,
		MenuPicture: upm.MenuPicture,
		MenuPrice:   upm.MenuPrice,
	}

	req := request.MenuUpdate{
		MenuTypeId:  upm.MenuTypeId,
		WartegId:    upm.WartegId,
		MenuName:    upm.MenuName,
		MenuDetail:  upm.MenuDetail,
		MenuPicture: upm.MenuPicture,
		MenuPrice:   upm.MenuPrice,
	}

	upmenu, err := u.menuRepo.MenuUpdate(ctx, menu_id, req)

	if err != nil {
		return resp, err
	}

	return upmenu, err

}

func (u *MenuUsecase) MenuList(ctx context.Context, warteg_id string, menu_type_id int) (list []response.MenuList, err error) {
	resp := []response.MenuList{}

	menulist, err := u.menuRepo.MenuList(ctx, warteg_id, menu_type_id)

	if err != nil {
		return resp, err
	}

	return menulist, err

}

func (u *MenuUsecase) MenuDetail(ctx context.Context, menu_id string) (mnd response.MenuDetail, err error) {
	resp := response.MenuDetail{}

	mdetail, err := u.menuRepo.MenuDetail(ctx, menu_id)

	if err != nil {
		return resp, err
	}

	return mdetail, err
}
