package middlewares

import (
	"errors"
	"group-project2/configs"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateToken(ID uint, isRenter bool) (string, error) {
	if ID < 1 {
		return "", errors.New("user_id tidak valid")
	}
	data := jwt.MapClaims{}
	data["id"] = ID
	data["isRenter"] = isRenter
	data["expired"] = time.Now().Add(time.Hour * 1).Unix()
	data["authorized"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString([]byte(configs.JWT_SECRET))
}

func ExtractTokenUserID(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		data := user.Claims.(jwt.MapClaims)
		id := data["id"].(int)
		return id
	}
	return 0
}
