package http

import (
	"github.com/cpartogi/foodmenu/constant"
	"github.com/cpartogi/foodmenu/module/menu"
	"github.com/cpartogi/foodmenu/pkg/utils"
	"github.com/labstack/echo/v4"
)

// AuthHandler  represent the httphandler for auth
type MenuHandler struct {
	menuUsecase menu.Usecase
}

// NewAuthHandler will initialize the contact/ resources endpoint
func NewMenuHandler(e *echo.Echo, us menu.Usecase) {
	handler := &MenuHandler{
		menuUsecase: us,
	}

	router := e.Group("/v1")
	router.GET("/menus/typelist", handler.MenuType)
}

// Menu Type godoc
// @Summary Menu Type
// @Description Menu Type
// @Tags Menu
// @Accept  json
// @Produce  json
// @Success 200 {object} response.SwaggerMenuType
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/menus/typelist [get]
// Disburselog handles HTTP request for disbursement history
func (h *MenuHandler) MenuType(c echo.Context) error {

	ctx := c.Request().Context()
	bal, err := h.menuUsecase.MenuType(ctx)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessGetData, bal)

}
