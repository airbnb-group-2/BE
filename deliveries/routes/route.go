package routes

import (
	"group-project2/deliveries/controllers/auth"
	"group-project2/deliveries/controllers/image"
	"group-project2/deliveries/controllers/rating"
	"group-project2/deliveries/controllers/room"
	"group-project2/deliveries/controllers/user"
	"group-project2/deliveries/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPaths(e *echo.Echo, ac *auth.AuthController, uc *user.UserController, rc *room.RoomController, ic *image.ImageController, rtc *rating.RatingController) {
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
}
