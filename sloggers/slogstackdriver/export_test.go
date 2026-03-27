package slogstackdriver

import (
	logpbtype "google.golang.org/genproto/googleapis/logging/type"

	"dev.latticeruntime.com/slog/v3"
)

func Sev(level slog.Level) logpbtype.LogSeverity {
	return sev(level)
}
