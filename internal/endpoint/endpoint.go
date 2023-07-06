package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {

	s := fmt.Sprintf("Days left until the new year: %d", e.s.DaysLeft())

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}

	return nil
}
