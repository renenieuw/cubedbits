package logging


import (
	"context"
	"fmt"
	"io"
	"log/slog"
)

type FilteredJSONHandler struct {
	jSONHandler slog.Handler
	logFilter *LogFilter
	ignoreLogs bool
}

func NewFilteredJSONHandler(w io.Writer, opts *slog.HandlerOptions, logFilter *LogFilter) *FilteredJSONHandler {
	return &FilteredJSONHandler{
		jSONHandler: slog.NewJSONHandler(w,opts),
		logFilter: logFilter,
	}
}

func (h *FilteredJSONHandler) Enabled(context context.Context, level slog.Level) bool {
	return h.jSONHandler.Enabled(context, level)
}

func (h *FilteredJSONHandler) Handle(context context.Context, r slog.Record) error {
	if(h.ignoreLogs == false) {
		return h.jSONHandler.Handle(context, r)
	}
	return nil;
}

func (h *FilteredJSONHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if(h.logFilter.HandleAttributes(attrs)) {
		var jsonHandler = h.jSONHandler.WithAttrs(attrs)
		return &FilteredJSONHandler{jSONHandler: jsonHandler, logFilter: h.logFilter, ignoreLogs: false }

	} else {
		var jsonHandler = h.jSONHandler.WithAttrs(attrs)
		return &FilteredJSONHandler{jSONHandler: jsonHandler, logFilter: h.logFilter, ignoreLogs: true }
	}
}

func (h *FilteredJSONHandler) WithGroup(name string) slog.Handler {
	fmt.Println("testgroup")
	return h.jSONHandler.WithGroup(name)
}
