package xmiddleware

import (
	"context"
	"log/slog"

	"github.com/aziemski/bookstore/internal/x/xlog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RequestLogger(log *slog.Logger) echo.MiddlewareFunc {
	mw := middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				log.LogAttrs(context.Background(), slog.LevelInfo, "req",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				log.LogAttrs(context.Background(), slog.LevelError, "req_err",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					xlog.Err(v.Error),
				)
			}
			return nil
		},
	})

	return mw
}
