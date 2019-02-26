package api

import (
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Service interface {
	WriteImage(w io.Writer) error
	OnClick(x, y int) error
}

type Handlers struct {
	svc Service
}

func New(svc Service) *Handlers {
	return &Handlers{
		svc: svc,
	}
}

func (handlers *Handlers) WriteImage(c echo.Context) error {

	w := c.Response()
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)

	return handlers.svc.WriteImage(w)
}

func (handlers *Handlers) OnClick(c echo.Context) error {
	x, _ := strconv.Atoi(c.QueryParam("x"))
	y, _ := strconv.Atoi(c.QueryParam("y"))

	if err := handlers.svc.OnClick(x, y); err != nil {
		return err
	}

	return c.String(http.StatusOK, "ok")
}
