package logging

import (
	"log/slog"

)

type LogFilter struct {
	Contexts map[string]bool
}

func (l LogFilter) HandleAttributes(attrs []slog.Attr) bool {
	for _, attr := range attrs {
		if(attr.Key == "Context"){
			for key, val := range l.Contexts {
				if(key == attr.Value.String()){
					return val;
				}
			}
		}
	}
	return false;
}
