package logging


import (
	"os"
)

type LogWriter struct {
	Path string
}

func (lw *LogWriter) Write(p []byte) (n int, err error) {

	logFile, err := os.OpenFile(lw.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logFile.Write(p)
	logFile.Close()

	return len(p), nil
}
