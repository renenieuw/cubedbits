package main

import (
	"log"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/logdyhq/logdy-core/logdy"
)

type Logger struct {
	logdy   logdy.Logdy
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	l.handler.ServeHTTP(w, r)

	// If this is a request to Logdy backend, ignore it
	if strings.HasPrefix(r.URL.Path, l.logdy.Config().HttpPathPrefix) {
		return
	}

	l.logdy.Log(logdy.Fields{
		"ua":     r.Header.Get("user-agent"),
		"method": r.Method,
		"path":   r.URL.Path,
		"query":  r.URL.RawQuery,
		"time":   time.Since(start),
	})

	l.logdy.Log(logdy.Fields{
		"method": "blah",
	})

	slog.Info("test")
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	d = d + 1

	ebitenutil.DebugPrint(screen, "Hello, starss!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

var d int

func main() {
	d = 0
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	logger := logdy.InitializeLogdy(logdy.Config{
		HttpPathPrefix: "/_logdy-ui",
	}, mux)

	go func() {
		addr := ":8082"
		log.Printf("server is listening at %s", addr)
		log.Fatal(http.ListenAndServe(addr, &Logger{logdy: logger, handler: mux}))
	}()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Starss")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
