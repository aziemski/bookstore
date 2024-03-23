package xlog

import "log/slog"

func Err(err error) slog.Attr {
	return slog.Any("err", err)
}
