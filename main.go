package main

import (
	"net/http"
	"time"

	_menuHttpHandler "github.com/cpartogi/foodmenu/module/menu/handler/http"
	_menuRepo "github.com/cpartogi/foodmenu/module/menu/store"
	_menu "github.com/cpartogi/foodmenu/module/menu/usecase"

	_ "github.com/cpartogi/foodmenu/docs"
	appInit "github.com/cpartogi/foodmenu/init"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	log "go.uber.org/zap"
)

func init() {
	// Start pre-requisite app dependencies
	appInit.StartAppInit()
}

func main() {

	mysqlDb, err := appInit.ConnectToMySqlServer()
	if err != nil {
		log.S().Fatal(err)
	}

	// init router
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is healthy")
	})

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	// DI: Repository & Usecase
	menuRepo := _menuRepo.NewStore(mysqlDb.DB)

	menuUc := _menu.NewMenuUsecase(menuRepo, timeoutContext)

	// End of DI Stepss

	_menuHttpHandler.NewMenuHandler(e, menuUc)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// start serve
	e.Logger.Fatal(e.Start(viper.GetString("api.port")))
}
