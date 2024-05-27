package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
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

func HasSessionCookie(c echo.Context) bool {
	cookie, err := c.Cookie("sid")
	if err != nil {
		return false
	}
	return cookie.Value != ""
}

func DeleteSessionCookie(c echo.Context) {
	c.SetCookie(&http.Cookie{
		Name:    "sid",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
}

func GetSessionTimeOutTime() (time.Time, error) {

	sessionTimeout := os.Getenv("SESSION_TIMEOUT")
	sessionTimeoutUnits := os.Getenv("SESSION_TIMEOUT_UNITS")

	if sessionTimeout == "" || sessionTimeoutUnits == "" {
		return time.Time{}, errors.New("variables SESSION_TIMEOUT and SESSION_TIMEOUT_UNITS must be set")
	}

	timeout, err := strconv.Atoi(sessionTimeout)
	if err != nil {
		return time.Time{}, err
	}
	timeoutUnits := strings.ToUpper(sessionTimeoutUnits)
	var timeoutDuration time.Duration

	switch timeoutUnits {
	case "M":
		timeoutDuration = time.Minute
	case "H":
		timeoutDuration = time.Hour
	default:
		return time.Time{}, errors.New("invalid SESSION_TIMEOUT_UNITS, valid values are M or H")
	}

	return time.Now().Add(time.Duration(timeout) * timeoutDuration), nil
}
