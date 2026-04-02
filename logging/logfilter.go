package logging

import (
	"log/slog"
	"strings"
)

type LogFilter struct {
	Contexts map[string]bool
}

func (l LogFilter) HandleAttributes(attrs []slog.Attr) bool {
	for _, attr := range attrs {
		if(attr.Key == "Context"){
			for filter, val := range l.Contexts {
				if(strings.HasSuffix(filter,"*")) {
					filter = strings.Replace(filter, "*", "", -1)
					if strings.HasPrefix(attr.Value.String(), filter) {
						return val
					}
				} else {
					if(filter == attr.Value.String()) {
						return val
					}
				}



				// if(key == attr.Value.String()){
				// 	return val;
				// }
			}
		}
	}
	return false;
}
