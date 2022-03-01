package routes

import (
	"group-project2/deliveries/controllers/auth"
	"group-project2/deliveries/controllers/book"
	"group-project2/deliveries/controllers/image"
	paymentmethod "group-project2/deliveries/controllers/payment-method"
	"group-project2/deliveries/controllers/rating"
	"group-project2/deliveries/controllers/room"
	"group-project2/deliveries/controllers/user"
	"group-project2/deliveries/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPaths(e *echo.Echo, ac *auth.AuthController, uc *user.UserController, rc *room.RoomController, ic *image.ImageController, rtc *rating.RatingController, pmc *paymentmethod.PaymentMethodController, bc *book.BookController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	a := e.Group("/login")
	a.POST("", ac.Login())

	u := e.Group("/users")
	u.POST("", uc.Insert())
	uj := u.Group("/jwt")
	uj.Use(middlewares.JWTMiddleware())
	uj.GET("", uc.GetUserByID())
	uj.PUT("", uc.Update())
	uj.PUT("/setrenter", uc.SetRenter())
	uj.DELETE("", uc.DeleteByID())

	r := e.Group("/rooms")
	r.GET("/all", rc.GetAllRooms())
	r.GET("/:id", rc.GetRoomByID())
	r.GET("", rc.GetRoomsByUserID())
	r.GET("", rc.GetRoomsByCity())
	rj := r.Group("/jwt")
	rj.Use(middlewares.JWTMiddleware())
	rj.POST("", rc.Insert())
	rj.PUT("/:id", rc.Update())
	rj.DELETE("/:id", rc.Delete())

	i := e.Group("/images")
	i.GET("", ic.GetImagesByRoomID())
	i.GET("/:id", ic.GetImageByID())
	ij := i.Group("/jwt")
	ij.Use(middlewares.JWTMiddleware())
	ij.POST("", ic.Insert())
	ij.PUT("/:id", ic.Update())
	ij.DELETE("/:id", ic.DeleteImageByID())
	ij.DELETE("/delete", ic.DeleteImageByRoomID())

	rt := e.Group("/ratings")
	rt.GET("", rtc.GetRatingsByRoomID())
	rt.POST("", rtc.Insert(), middlewares.JWTMiddleware())

	pm := e.Group("/payment-methods")
	pm.GET("", pmc.Get())
	pmj := pm.Group("/jwt")
	pmj.Use(middlewares.JWTMiddleware())
	pmj.POST("", pmc.Insert())
	pmj.DELETE("/:id", pmc.Delete())

	b := e.Group("/books")
	b.Use(middlewares.JWTMiddleware())
	b.POST("", bc.Insert())
	b.GET("/check-status/:id", bc.GetStatusID())
	b.GET("/user-books", bc.GetAllBooksByUserID())
	b.GET("/user-histories", bc.GetBookHistoryByUserID())
	b.PUT("/set-paid/:id", bc.SetPaid())
	b.PUT("/set-cancel/:id", bc.SetCancel())
	b.PUT("/set-checkin/:id", bc.SetCheckInTime())
	b.PUT("/set-checkout/:id", bc.SetCheckOutTime())

	// m := e.Group("/midtrans")
	// m.GET("", midtranspay.Notification())
}
