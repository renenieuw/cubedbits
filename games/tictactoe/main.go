package main

import (
	"fmt"
	"image/color"
	"log/slog"
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

	g.stateMachine.Draw(g.world, screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	logFile, err := os.OpenFile("c:/data/logging/tictactoe/default.json", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}



	handlerOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
		AddSource: true,
	}
	var contexts []logging.Context
	contexts = append(contexts, logging.Context{ Name: "*", Enabled: true })
	logFilter := logging.LogFilter{ Contexts: contexts }
	mainlogger := slog.New(logging.NewFilteredJSONHandler(logFile, "C:/Data/Logging/tictactoe/dump.json",   handlerOpts, &logFilter))


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
	ss, _ := loader.LoadSpriteSheetsFromJson(dataGame, "tictactoe", "tictactoe", "tictactoe.png")

	v := resources.Textures { }
	v.AddSpritesheet(sse)
	v.AddSpritesheet(ss)
	resources.SetDefault(&v)

	vv := resources.Default()



	// maps.Copy(sse, ss)

	logger.Debug("Loaded spritesheetssss", "Object", sse)
	logger.Debug("Dump", "Dump", vv.Spritesheets)


	r.ScreenDimensions = &resources.ScreenDimensions{Width: 640, Height: 480, Title: "TicTacToe"}
	r.SpriteSheets = &sse

	spr := vv.GetSpriteRender("Tiles/O.png")
	logger.Debug("GetTexture found it" + spr.SpriteGroup)
	logger.Debug(fmt.Sprintf( "GetTexture found it %d", spr.SpriteNumber))


	// Load fonts
	fonts := loader.LoadFonts("cubedBits/fonts.toml")
	ecs.AddResource(w, &fonts)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Starss")
	if err := ebiten.RunGame(&Game{w, st.Init(&ts.GameplayState{}, w)}); err != nil {
		slog.Error("error", "err", err)
	}
}
