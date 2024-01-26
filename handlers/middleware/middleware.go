package middleware

import (
	"context"
	"htmxdemo/types"
	"math/rand"
	"time"

	"github.com/labstack/echo/v4"
)

func SimulateNetworkLatency(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		time.Sleep(time.Duration(rand.Int63n(250)) * time.Millisecond)
		return next(c)
	}
}

func CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), types.MyDataKey, true)

		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}

func SetPath(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		route := c.Request().URL.Path

		ctx := context.WithValue(c.Request().Context(), types.MyRouteKey, route)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
