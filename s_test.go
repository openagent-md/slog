package slog_test

import (
	"bytes"
	"testing"

	"github.com/openagent-md/slog/v3"
	"github.com/openagent-md/slog/v3/internal/assert"
	"github.com/openagent-md/slog/v3/internal/entryhuman"
	"github.com/openagent-md/slog/v3/sloggers/sloghuman"
)

func TestStdlib(t *testing.T) {
	t.Parallel()

	b := &bytes.Buffer{}
	l := slog.Make(sloghuman.Sink(b)).With(
		slog.F("hi", "we"),
	)
	stdlibLog := slog.Stdlib(bg, l, slog.LevelInfo)
	stdlibLog.Println("stdlib")

	et, rest, err := entryhuman.StripTimestamp(b.String())
	assert.Success(t, "strip timestamp", err)
	assert.False(t, "timestamp", et.IsZero())
	assert.Equal(t, "entry", " [info]  stdlib: stdlib  hi=we\n", rest)
}
