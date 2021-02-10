package usecase

import (
	"context"
	"time"

	"github.com/cpartogi/foodmenu/module/menu"
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
