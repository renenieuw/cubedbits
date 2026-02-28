package main

import (
	"image/color"
	"log"

	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/mlange-42/ark/ecs"
)

type Game struct{}

// Position component
type Position struct {
	X, Y float32
}

// Velocity component
type Velocity struct {
	DX, DY float32
}

func (g *Game) Update() error {
	return nil
}

var col color.RGBA
var world *ecs.World

func (g *Game) Draw(screen *ebiten.Image) {
	filter := ecs.NewFilter2[Position, Velocity](world)

	for range 2 {
		// Get a fresh query and iterate it
		query := filter.Query()
		for query.Next() {
			// Component access through the Query
			pos, vel := query.Get()
			// Update component fields
			// pos.X += vel.DX
			// pos.Y += vel.DY
			//
			vel.DX = 0
			vel.DY = 0

			pos.X = pos.X + vel.DX
			pos.Y = pos.Y + vel.DY

			vector.FillRect(screen, pos.X, pos.Y, 2, 2, col, false)

		}
	}

	//	ebitenutil.DrawRect(screen, 11, 12, settings.Scale, settings.Scale, particleData.Color)
	ebitenutil.DebugPrint(screen, "Hello, starss2!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	col = color.RGBA{0x80, 0x80, 0x80, 0xff}

	world = ecs.NewWorld()

	mapper := ecs.NewMap1[Sprite](world)

	for range 30 {
		_ = mapper.NewEntity(
			&Sprite
			},
		)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Starss")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
