package main

import (
	"image/color"
	"log"
	"maps"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/ark/ecs"
	ga "github.com/renenieuw/cubedbits/assets"
	"github.com/renenieuw/cubedbits/games/ticktacktoe/assets"
	ts "github.com/renenieuw/cubedbits/games/ticktacktoe/states"
	"github.com/renenieuw/cubedbits/loader"
	"github.com/renenieuw/cubedbits/resources"
	st "github.com/renenieuw/cubedbits/states"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
		log.Printf("touchIds: %v", touchIDs)
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

	ga.Assets = map[string]func(string)[]byte{
        "cubedBits": ga.GetAsset,
        "ticktacktoe": assets.GetAsset,
    }

	w := ecs.NewWorld()
	r := resources.InitResources()
	ecs.AddResource(w, r)



	dataGameEngine := string(ga.Spritesheets[:])
	dataGame := string(assets.Spritesheets[:])

	// sse := loader.LoadSpriteSheets("../../assets/metadata/spritesheets/spritesheets.toml")
//	ss := loader.LoadSpriteSheets("assets/metadata/spritesheets/spritesheets.toml")

	sse := loader.LoadSpriteSheetsFromString(dataGameEngine)
	ss := loader.LoadSpriteSheetsFromString(dataGame)



	maps.Copy(sse, ss)

	r.ScreenDimensions = &resources.ScreenDimensions{Width: 640, Height: 480, Title: "TickTackToe"}
	r.SpriteSheets = &sse

	//	r := resources.ScreenDimensions{Width: 640, Height: 480, Title: "TickTackToe"}



	// Load fonts
	fonts := loader.LoadFonts("cubedBits/fonts.toml")
	ecs.AddResource(w, &fonts)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Starss")
	if err := ebiten.RunGame(&Game{w, st.Init(&ts.GameplayState{}, w)}); err != nil {
		log.Fatal(err)
	}
}
