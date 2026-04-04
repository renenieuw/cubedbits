package logging

import (
	"log/slog"
	"strings"
)

type Context struct {
	Name string
	Enabled bool
}

type LogFilter struct {
	Contexts []Context
}

func (l LogFilter) HandleAttributes(attrs []slog.Attr) bool {
	for _, attr := range attrs {
		if(attr.Key == "Context"){
			for _, val := range l.Contexts {
				if(strings.HasSuffix(val.Name,"*")) {
					filter := strings.Replace(val.Name, "*", "", -1)
					if strings.HasPrefix(attr.Value.String(), filter) {
						return val.Enabled
					}
				} else {
					if(val.Name == attr.Value.String()) {
						return val.Enabled
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
