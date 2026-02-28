package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/ark/ecs"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/assets"
	ts "remapit.visualstudio.com/cubedbits/cubedbitsengine/games/ticktacktoe/states"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/loader"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/resources"
	st "remapit.visualstudio.com/cubedbits/cubedbitsengine/states"
)

const (
	gameWidth  = 720
	gameHeight = 600
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
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(gopherImage, op)
	g.stateMachine.Draw(g.world, screen)

	// slog.Info(fmt.Sprintf("%s%d", "Drawing game", gopherImage.Bounds().Max.X))
	// ebitenutil.DrawRect(screen, 11, 12, settings.Scale, settings.Scale, particleData.Color)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {

	img, _, err := image.Decode(bytes.NewReader(assets.Background))
	if err != nil {
		log.Fatal(err)
	}

	gopherImage = ebiten.NewImageFromImage(img)

	r := resources.ScreenDimensions{Width: 640, Height: 480, Title: "TickTackToe"}

	w := ecs.NewWorld()
	ecs.AddResource(w, &r)

	spriteSheets := loader.LoadSpriteSheets("../../assets/metadata/spritesheets/spritesheets.toml")
	slog.Info(fmt.Sprintf("%d", spriteSheets.SpriteSheets["background"].Sprites[0].Width))
	slog.Info(fmt.Sprintf("%d", len(spriteSheets.SpriteSheets)))

	ecs.AddResource(w, &spriteSheets)

	// Load fonts
	fonts := loader.LoadFonts("../../assets/metadata/fonts/fonts.toml")
	ecs.AddResource(w, &fonts)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Starss")
	if err := ebiten.RunGame(&Game{w, st.Init(&ts.MenuState{}, w)}); err != nil {
		log.Fatal(err)
	}
}
