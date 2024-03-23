package ports

import (
	"fmt"
	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
	"net/http"
)

var _ ServerInterface = (*HttpServer)(nil)

type HttpServer struct {
}

func (h *HttpServer) GetBookById(c echo.Context, id openapi_types.UUID) error {
	return c.JSON(http.StatusOK, Book{Title: fmt.Sprintf("book %v", id)})
}

func (h *HttpServer) RegisterWith(e *echo.Echo) {
	g := e.Group("api/v1")
	RegisterHandlers(g, h)
}
