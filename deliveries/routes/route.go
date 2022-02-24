package routes

import (
	"group-project2/deliveries/controllers/auth"
	"group-project2/deliveries/controllers/user"
	"group-project2/deliveries/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPaths(e *echo.Echo, ac *auth.AuthController, uc *user.UserController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	ar := e.Group("/login")
	ar.POST("", ac.Login())

	ur := e.Group("/users")
	ur.POST("", uc.Insert())
	ur.GET("", uc.GetUserByID(), middlewares.JWTMiddleware())
	ur.PUT("", uc.Update(), middlewares.JWTMiddleware())
	ur.PUT("/setrenter", uc.SetRenter(), middlewares.JWTMiddleware())
	ur.DELETE("", uc.DeleteByID(), middlewares.JWTMiddleware())

}
