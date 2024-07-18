package mockctx

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MockContext struct {
	data   map[string]any
	cookie map[string]http.Cookie
	echo.Context
}

func New() MockContext {
	c := MockContext{}
	c.data = make(map[string]any)
	c.cookie = make(map[string]http.Cookie)
	return c
}

func (c MockContext) Get(key string) any {
	v, ok := c.data[key]
	if !ok {
		return nil
	}
	return v
}

func (c MockContext) Set(key string, value any) {
	c.data[key] = value
}

func (c MockContext) Cookie(key string) (*http.Cookie, error) {
	v, ok := c.cookie[key]
	if !ok {
		return nil, errors.New("cookie " + key + " not found")
	}
	return &v, nil
}

func (c MockContext) SetSimpleCookie(key string, value string) {
	c.cookie[key] = http.Cookie{
		Name:     key,
		Value:    value,
		HttpOnly: true,
		MaxAge:   5 * 60,
		SameSite: http.SameSiteLaxMode,
	}
}

func (c MockContext) SetCookie(cookie *http.Cookie) {
	key := cookie.Name
	c.cookie[key] = *cookie
}
