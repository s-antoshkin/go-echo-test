package app

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/s-antoshkin/go-echo-test/internal/endpoint"
	mw "github.com/s-antoshkin/go-echo-test/internal/middleware"
	"github.com/s-antoshkin/go-echo-test/internal/service"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()
	a.e = endpoint.New(a.s)
	a.echo = echo.New()

	a.echo.Use(mw.RoleCheck)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Println("Server running...")

	if err := a.echo.Start(":8080"); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
