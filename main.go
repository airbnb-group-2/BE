package main

import (
	"fmt"
	"group-project2/configs"
	_AuthController "group-project2/deliveries/controllers/auth"
	_ImageController "group-project2/deliveries/controllers/image"
	_RatingController "group-project2/deliveries/controllers/rating"
	_RoomController "group-project2/deliveries/controllers/room"
	_UserController "group-project2/deliveries/controllers/user"
	"group-project2/deliveries/routes"
	_AuthRepo "group-project2/repositories/auth"
	_ImageRepo "group-project2/repositories/image"
	_RatingRepo "group-project2/repositories/rating"
	_RoomRepo "group-project2/repositories/room"
	_UserRepo "group-project2/repositories/user"
	awss3 "group-project2/services/aws-s3"
	"group-project2/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	authRepo := _AuthRepo.New(db)
	userRepo := _UserRepo.New(db)
	roomRepo := _RoomRepo.New(db)
	imageRepo := _ImageRepo.New(db)
	ratingRepo := _RatingRepo.New(db)

	awsSess := awss3.InitS3(config.S3_KEY, config.S3_SECRET, config.S3_REGION)

	ac := _AuthController.New(authRepo)
	uc := _UserController.New(userRepo)
	rc := _RoomController.New(roomRepo)
	ic := _ImageController.New(imageRepo, config, awsSess)
	rtc := _RatingController.New(ratingRepo)

	e := echo.New()

	routes.RegisterPaths(e, ac, uc, rc, ic, rtc)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.PORT)))
}
