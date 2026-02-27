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
	gc "remapit.visualstudio.com/cubedbits/cubedbitsengine/components"
	tc "remapit.visualstudio.com/cubedbits/cubedbitsengine/games/ticktacktoe/components"
	ts "remapit.visualstudio.com/cubedbits/cubedbitsengine/games/ticktacktoe/states"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/loader"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/math"
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

	col = color.RGBA{0x80, 0x80, 0x80, 0xff}
	r := resources.ScreenDimensions{Width: 640, Height: 480, Title: "TickTackToe"}

	w := ecs.NewWorld()
	ecs.AddResource(w, &r)

	spriteSheets := loader.LoadSpriteSheets("../../assets/metadata/spritesheets/spritesheets.toml")
	slog.Info(fmt.Sprintf("%d", spriteSheets["background"].Sprites[0].Width))
	slog.Info(fmt.Sprintf("%d", len(spriteSheets)))
	ecs.AddResource(w, &spriteSheets)

	spriteSheet := spriteSheets["game"]

	mapper := ecs.NewMap1[gc.SpriteRender](w)
	mapper2 := ecs.NewMap2[gc.SpriteRender, gc.Transform](w)
	mapper3 := ecs.NewMap1[tc.Tile](w)

	for range 30 {
		_ = mapper.NewEntity(

			&gc.SpriteRender{
				SpriteSheet:  &spriteSheet,
				SpriteNumber: 3,
				Options:      ebiten.DrawImageOptions{},
			},
		)
	}

	mapper2.NewEntity(

		&gc.SpriteRender{
			SpriteSheet:  &spriteSheet,
			SpriteNumber: 2,
			Options:      ebiten.DrawImageOptions{},
		},
		&gc.Transform{Translation: math.Vector2{X: 133, Y: 220}},
	)

	mapper3.NewEntity(
		&tc.Tile{X: 33, Y: 22, State: 123},
	)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Starss")
	if err := ebiten.RunGame(&Game{w, st.Init(&ts.GameplayState{}, w)}); err != nil {
		log.Fatal(err)
	}
}
