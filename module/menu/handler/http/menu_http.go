package http

import (
	"strconv"

	"github.com/cpartogi/foodmenu/constant"
	"github.com/cpartogi/foodmenu/module/menu"
	"github.com/cpartogi/foodmenu/pkg/utils"
	"github.com/cpartogi/foodmenu/schema/request"
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
	router.GET("/menus/list", handler.MenuList)
	router.POST("/menu", handler.MenuAdd)
	router.DELETE("/menu/:menu_id", handler.MenuDelete)
	router.PUT("/menu/:menu_id", handler.MenuUpdate)
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
// Menu Type handles HTTP request for menu type
func (h *MenuHandler) MenuType(c echo.Context) error {

	ctx := c.Request().Context()
	bal, err := h.menuUsecase.MenuType(ctx)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessGetData, bal)

}

// MenuAdd godoc
// @Summary Add Menu
// @Description Add Menu
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param request body request.Menu true "Request Body"
// @Success 201 {object} response.SwaggerMenuAdd
// @Failure 400 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/menu [post]
// Menuadd handles HTTP request for add menu
func (h *MenuHandler) MenuAdd(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Menu{}

	//parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, map[string]interface{}{})
	}

	//validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, map[string]interface{}{})
	}

	reg, err := h.menuUsecase.MenuAdd(ctx, req)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.CreatedResponse(c, "Succes add menu", reg)
}

// MenuDelete godoc
// @Summary Delete Menu
// @Description Delete Menu
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param menu_id path string true "Menu Id"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/menu/{menu_id} [delete]
// MenuDelete handles HTTP request for delete menu
func (h *MenuHandler) MenuDelete(c echo.Context) error {
	ctx := c.Request().Context()
	menuId := c.Param("menu_id")

	_, err := h.menuUsecase.MenuDelete(ctx, menuId)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, "Success delete menu", map[string]interface{}{})
}

// MenuUpdate godoc
// @Summary Update Menu
// @Description Update Menu
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param menu_id path string true "Menu Id"
// @Param request body request.MenuUpdate true "Request Body"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/menu/{menu_id} [put]
// MenuUpdate handles HTTP request for update menu
func (h *MenuHandler) MenuUpdate(c echo.Context) error {
	ctx := c.Request().Context()
	menuId := c.Param("menu_id")
	req := request.MenuUpdate{}

	//parsing
	err := utils.ParsingParameter(c, &req)
	if err != nil {
		return utils.ErrorParsing(c, err, map[string]interface{}{})
	}

	//validate
	err = utils.ValidateParameter(c, &req)
	if err != nil {
		return utils.ErrorValidate(c, err, map[string]interface{}{})
	}

	reg, err := h.menuUsecase.MenuUpdate(ctx, menuId, req)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, "Succes update menu", reg)

}

// MenuList godoc
// @Summary  Menu list
// @Description Menu List
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param warteg_id query string true "warteg id"
// @Param menu_type_id query string true "menu type id"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 404 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/menus/list [get]
// MenuList handles HTTP request for menu list
func (h *MenuHandler) MenuList(c echo.Context) error {
	ctx := c.Request().Context()
	queryValues := c.Request().URL.Query()
	wartegId := queryValues.Get("warteg_id")
	menu_typeid := queryValues.Get("menu_type_id")

	menu_type_id, parseErr := strconv.Atoi(menu_typeid)

	if parseErr != nil {
		return utils.ErrorResponse(c, parseErr, map[string]interface{}{})
	}

	menu, err := h.menuUsecase.MenuList(ctx, wartegId, menu_type_id)
	if err != nil {
		return utils.ErrorResponse(c, err, map[string]interface{}{})
	}

	return utils.SuccessResponse(c, constant.SuccessGetData, menu)
}
