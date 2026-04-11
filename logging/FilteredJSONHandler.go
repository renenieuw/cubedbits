package logging

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
)

type FilteredJSONHandler struct {
	jSONHandler slog.Handler
	jSONDumpHandler slog.Handler
	dumpBytes *bytes.Buffer
	logFilter *LogFilter
	ignoreLogs bool
}

func NewFilteredJSONHandler(w io.Writer, opts *slog.HandlerOptions, logFilter *LogFilter) *FilteredJSONHandler {
	s := "Hello"
	buf := bytes.NewBufferString(s)
	fmt.Fprint(buf, ", World!")
	return &FilteredJSONHandler{
		jSONHandler: slog.NewJSONHandler(w,opts),
		jSONDumpHandler: slog.NewJSONHandler(buf,opts),
		logFilter: logFilter,
		dumpBytes: buf,
	}
}

func (h *FilteredJSONHandler) Enabled(context context.Context, level slog.Level) bool {
	return h.jSONHandler.Enabled(context, level)
}

func (h *FilteredJSONHandler) Handle(context context.Context, r slog.Record) error {
	r.Attrs(func(a slog.Attr) bool {
		if(a.Key == "Dump")	{
			h.dumpBytes.Reset()

			h.jSONDumpHandler.Handle(context,r)
			os.WriteFile("C:/Data/Logging/tictactoe/dump.json", h.dumpBytes.Bytes(), 0644)
			fmt.Println(h.dumpBytes)
		}
		return true
	})

	if(h.ignoreLogs == false) {
		return h.jSONHandler.Handle(context, r)
	}
	return nil;
}

func (h *FilteredJSONHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if(h.logFilter.HandleAttributes(attrs)) {
		var jsonHandler = h.jSONHandler.WithAttrs(attrs)
		var jsonDumpHandler = h.jSONDumpHandler.WithAttrs(attrs)
		return &FilteredJSONHandler{jSONHandler: jsonHandler, jSONDumpHandler: jsonDumpHandler, dumpBytes: h.dumpBytes, logFilter: h.logFilter, ignoreLogs: false }

	} else {
		var jsonHandler = h.jSONHandler.WithAttrs(attrs)
		var jsonDumpHandler = h.jSONDumpHandler.WithAttrs(attrs)
		return &FilteredJSONHandler{jSONHandler: jsonHandler, jSONDumpHandler: jsonDumpHandler, dumpBytes: h.dumpBytes, logFilter: h.logFilter, ignoreLogs: true }
	}
}

func (h *FilteredJSONHandler) WithGroup(name string) slog.Handler {
	fmt.Println("testgroup")
	return h.jSONHandler.WithGroup(name)
}
