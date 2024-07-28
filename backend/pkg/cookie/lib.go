package cookie

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"log"
)

var DefaultMaxAge = MaxAgeMonth
var MaxAgeMonth = 30 * 24 * time.Hour
var MaxAgeYear = 365 * 24 * time.Hour

func Get(c echo.Context, key string) (string, error) {
	cookie, err := c.Cookie(key)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func GetUint(c echo.Context, key string) (uint, error) {
	v, err := Get(c, key)
	if err != nil {
		return 0, err
	}
	ret, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(ret), nil
}

func Set(c echo.Context, key, value string) {
	c.SetCookie(From(key, value))
}

func SetUint(c echo.Context, key string, value uint) {
	v := strconv.FormatUint(uint64(value), 10)
	Set(c, key, v)
}

// note: do not specify more than two ages.
func From(key, value string, age ...time.Duration) *http.Cookie {
	var maxAge time.Duration
	if len(age) == 0 {
		maxAge = DefaultMaxAge
	} else if len(age) == 1 {
		maxAge = age[0]
	} else {
		log.Fatalln("Don't specify more than two ages, what are you doing???")
	}
	return &http.Cookie{
		Name:     key,
		Value:    value,
		MaxAge:   int(maxAge.Seconds()),
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
	}
}
