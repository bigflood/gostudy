package api

import (
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type Service interface {
	WriteImage(w io.Writer) error
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