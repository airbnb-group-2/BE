package routes

import (
	"group-project2/deliveries/controllers/auth"
	"group-project2/deliveries/controllers/image"
	"group-project2/deliveries/controllers/room"
	"group-project2/deliveries/controllers/user"
	"group-project2/deliveries/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPaths(e *echo.Echo, ac *auth.AuthController, uc *user.UserController, rc *room.RoomController, ic *image.ImageController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	ar := e.Group("/login")
	ar.POST("", ac.Login())

	ur := e.Group("/users")
	ur.POST("", uc.Insert())
	urj := ur.Group("/jwt")
	urj.Use(middlewares.JWTMiddleware())
	urj.GET("", uc.GetUserByID())
	urj.PUT("", uc.Update())
	urj.PUT("/setrenter", uc.SetRenter())
	urj.DELETE("", uc.DeleteByID())

	rr := e.Group("/rooms")
	rr.GET("", rc.GetAllRooms())
	rr.GET("/:id", rc.GetRoomByID())
	rr.GET("", rc.GetRoomsByCity())
	rrj := rr.Group("/jwt")
	rrj.Use(middlewares.JWTMiddleware())
	rrj.POST("", rc.Insert())
	rrj.GET("", rc.GetRoomsByUserID())
	rrj.PUT("/:id", rc.Update())
	rrj.DELETE("/:id", rc.Delete())

	ir := e.Group("/images")
	ir.GET("", ic.GetImagesByRoomID())
	ir.GET("/:id", ic.GetImageByID())
	irj := ir.Group("/jwt")
	irj.Use(middlewares.JWTMiddleware())
	irj.POST("", ic.Insert())
	irj.PUT("", ic.Update())
	irj.DELETE("/:id", ic.DeleteImageByID())
	irj.DELETE("", ic.DeleteImageByRoomID())
}
