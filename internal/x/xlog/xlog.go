package xlog

import (
	"log/slog"
	"os"
	"sync"
)

var (
	log  *slog.Logger
	once sync.Once
)

func GetLogger() *slog.Logger {
	once.Do(func() {
		log = slog.New(slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				AddSource:   false,
				Level:       slog.LevelDebug,
				ReplaceAttr: nil,
			}))
		slog.SetDefault(log)
	})

	return log
}

func Err(err error) slog.Attr {
	return slog.Any("err", err)
}
