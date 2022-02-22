package main

import (
	"fmt"
	"group-project2/configs"
	_AuthController "group-project2/deliveries/controllers/auth"
	_UserController "group-project2/deliveries/controllers/user"
	"group-project2/deliveries/routes"
	_AuthRepo "group-project2/repositories/auth"
	_UserRepo "group-project2/repositories/user"
	"group-project2/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	authRepo := _AuthRepo.New(db)
	userRepo := _UserRepo.New(db)

	ac := _AuthController.New(authRepo)
	uc := _UserController.New(userRepo)

	e := echo.New()

	routes.RegisterPaths(e, ac, uc)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
