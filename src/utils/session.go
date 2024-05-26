package utils

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func GenerateSessionID() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func CreateSessionCookie(c echo.Context, value string, expires time.Time) {
	cookie := new(http.Cookie)
	cookie.Name = "sid"
	cookie.Value = value
	cookie.Expires = expires
	cookie.Secure = true
	cookie.HttpOnly = true
	c.SetCookie(cookie)
}

func DeleteSessionCookie(c echo.Context) {
	c.SetCookie(&http.Cookie{
		Name:    "sid",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
}
