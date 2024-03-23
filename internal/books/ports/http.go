package ports

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

var _ ServerInterface = (*HTTPServer)(nil)

type HTTPServer struct {
}

func (h *HTTPServer) GetBookByID(c echo.Context, id openapi_types.UUID) error {
	return c.JSON(http.StatusOK, Book{Title: fmt.Sprintf("book %v", id)})
}

func (h *HTTPServer) RegisterWith(e *echo.Echo) {
	g := e.Group("api/v1")
	RegisterHandlers(g, h)
}
