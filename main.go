package main

import (
	"fmt"
	"group-project2/configs"
	_AuthController "group-project2/deliveries/controllers/auth"
	_BookController "group-project2/deliveries/controllers/book"
	_ImageController "group-project2/deliveries/controllers/image"
	_PMController "group-project2/deliveries/controllers/payment-method"
	_RatingController "group-project2/deliveries/controllers/rating"
	_RoomController "group-project2/deliveries/controllers/room"
	_UserController "group-project2/deliveries/controllers/user"
	"group-project2/deliveries/routes"
	_AuthRepo "group-project2/repositories/auth"
	_BookRepo "group-project2/repositories/book"
	_ImageRepo "group-project2/repositories/image"
	_PMRepo "group-project2/repositories/payment-method"
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
	pmRepo := _PMRepo.New(db)
	bookRepo := _BookRepo.New(db)

	awsSess := awss3.InitS3(config.S3_KEY, config.S3_SECRET, config.S3_REGION)

	ac := _AuthController.New(authRepo)
	uc := _UserController.New(userRepo)
	rc := _RoomController.New(roomRepo)
	ic := _ImageController.New(imageRepo, config, awsSess)
	rtc := _RatingController.New(ratingRepo)
	pmc := _PMController.New(pmRepo)
	bc := _BookController.New(bookRepo)

	e := echo.New()

	routes.RegisterPaths(e, ac, uc, rc, ic, rtc, pmc, bc)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.PORT)))
}

/* 	c := midtranspay.InitConnection()
midtranspay.CreateTransaction(c)
w := httptest.NewRecorder()
r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{\n    \"masked_card\": \"451111-1117\",\n    \"bank\": \"bca\",\n    \"eci\": \"06\",\n    \"channel_response_code\": \"7\",\n    \"channel_response_message\": \"Denied\",\n    \"transaction_time\": \"2021-06-08 15:49:54\",\n    \"gross_amount\": \"100000.00\",\n    \"currency\": \"IDR\",\n      \"payment_type\": \"credit_card\",\n    \"signature_key\": \"76fe68ed1b7040c7c329356c1cd47819be3ccb8b056376ff3488bfa9af1db52a85ded0501b2dab1de56e5852982133a9ef7a47c54222abbe72288c2c4f591a71\",\n    \"status_code\": \"202\",\n    \"transaction_id\": \"36f3687e-05d4-4879-a428-fd6d1ffb786e\",\n    \"transaction_status\": \"deny\",\n    \"fraud_status\": \"challenge\",\n    \"status_message\": \"Success, transaction is found\",\n    \"merchant_id\": \"G812785002\",\n    \"card_type\": \"credit\"\n}"))
midtranspay.Notification(w, r, c)
*/
