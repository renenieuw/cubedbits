package main

import (
	"image/color"
	"log/slog"
	"maps"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mlange-42/ark/ecs"
	ga "github.com/renenieuw/cubedbits/assets"
	"github.com/renenieuw/cubedbits/games/tictactoe/assets"
	ts "github.com/renenieuw/cubedbits/games/tictactoe/states"
	"github.com/renenieuw/cubedbits/loader"
	"github.com/renenieuw/cubedbits/logging"
	"github.com/renenieuw/cubedbits/resources"
	st "github.com/renenieuw/cubedbits/states"
)

const (
	gameWidth  = 720
	gameHeight = 600
)

var (
	Assets map[string]interface{}
)

type Game struct {
	world        *ecs.World
	stateMachine st.StateMachine
}

func (g *Game) Update() error {
	g.stateMachine.Update(g.world)
	return nil
}

var (
	col         color.RGBA
	gopherImage *ebiten.Image
)

func (g *Game) Draw(screen *ebiten.Image) {
	var touchIDs []ebiten.TouchID

	touchIDs = inpututil.AppendJustPressedTouchIDs(touchIDs[:0])
	if len(touchIDs) != 0 {
		slog.Debug("touchIds: ", "touchIDs", touchIDs)
    }

	//	op := &ebiten.DrawImageOptions{}

	//	screen.DrawImage(gopherImage, op)
	g.stateMachine.Draw(g.world, screen)

	// slog.Info(fmt.Sprintf("%s%d", "Drawing game", gopherImage.Bounds().Max.X))
	// ebitenutil.DrawRect(screen, 11, 12, settings.Scale, settings.Scale, particleData.Color)
	//

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	logFile, err := os.OpenFile("c:/temp/logs/tictactoe.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	handlerOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
		AddSource: true,
	}
	contexts := make(map[string]bool)
	contexts["Loader.*"] = true
	logFilter := logging.LogFilter{ Contexts: contexts }
	mainlogger := slog.New(logging.NewFilteredJSONHandler(logFile,handlerOpts, &logFilter))


	slog.SetDefault(mainlogger)
	logger := slog.Default().With("Context","Main")

	ga.Assets = map[string]func(string)[]byte{
        "cubedBits": ga.GetAsset,
        "tictactoe": assets.GetAsset,
    }

	logger.Debug("Initializing world")


	logger.Debug("Making new world")
	w := ecs.NewWorld()
	logger.Debug("Initializing resources")
	r := resources.InitResources()
	ecs.AddResource(w, r)

	dataGameEngine := string(ga.Spritesheets[:])
	dataGame := assets.TictactoeJson[:]

	sse := loader.LoadSpriteSheetsFromString(dataGameEngine)
	ss := loader.LoadSpriteSheetsFromJson(dataGame, "tictactoe", "tictactoe.png")

	maps.Copy(sse, ss)

	r.ScreenDimensions = &resources.ScreenDimensions{Width: 640, Height: 480, Title: "TicTacToe"}
	r.SpriteSheets = &sse

	// Load fonts
	fonts := loader.LoadFonts("cubedBits/fonts.toml")
	ecs.AddResource(w, &fonts)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Starss")
	if err := ebiten.RunGame(&Game{w, st.Init(&ts.GameplayState{}, w)}); err != nil {
		slog.Error("error", "err", err)
	}
}
