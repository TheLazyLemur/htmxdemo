package main

import (
	"context"
	"database/sql"
	db_setup "htmxdemo/db"
	"htmxdemo/db/queries"
	"htmxdemo/handlers"
	"htmxdemo/handlers/middleware"
	"htmxdemo/service"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"

	_ "embed"

	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	dbc *sql.DB
	e   *echo.Echo
	h   *handlers.Handlers
}

func main() {
	a := app{}

	a.setup()
	defer a.close()

	a.mountRoutes()
	a.startServer()
}

func (a *app) setup() {
	a.e = echo.New()
	a.e.HideBanner = true

	dbc := db_setup.Setup()

	a.dbc = dbc

	queries := queries.New()
	o := service.New(dbc, queries)
	a.h = handlers.New(o)
}

func (a *app) mountRoutes() {
	a.e.Use(middleware.SimulateNetworkLatency)
	a.e.Use(middleware.CheckAuth)
	a.e.Use(middleware.SetPath)

	a.e.GET("/", a.h.HandleHomePage)
	a.e.GET("/transactions", a.h.HandleTransactions)
	a.e.POST("/transactions", a.h.HandleInsertTransaction)
	a.e.POST("/hx/transactions/search", a.h.HandleTransactionsSearch)
	a.e.PUT("/hx/transactions", a.h.HandleUpdateTransaction)
}

func (a *app) startServer() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		if err := a.e.Start(":3000"); err != nil && err != http.ErrServerClosed {
			a.e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.e.Shutdown(ctx); err != nil {
		a.e.Logger.Fatal(err)
	}
}

func (a *app) close() error {
	return a.dbc.Close()
}
