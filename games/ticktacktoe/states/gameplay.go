package states

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/labstack/gommon/log"
	"github.com/mlange-42/ark/ecs"
	gc "github.com/renenieuw/cubedbits/components"
	tc "github.com/renenieuw/cubedbits/games/ticktacktoe/components"
	ts "github.com/renenieuw/cubedbits/games/ticktacktoe/systems"
	"github.com/renenieuw/cubedbits/loader"
	"github.com/renenieuw/cubedbits/math"
	"github.com/renenieuw/cubedbits/resources"
	"github.com/renenieuw/cubedbits/states"
)

// GameplayState is the main game state
type GameplayState struct {
	index      int
	turn       int
	TileSystem ts.TileSystem
}

// OnPause method
func (st *GameplayState) OnPause(world *ecs.World) {
	log.Info("Gameplay.OnPause")
}

func (st *GameplayState) GetTurn() int {
	return st.turn
}

func (st *GameplayState) SetTurn(turn int) {
	st.turn = turn
}

// OnResume method
func (st *GameplayState) OnResume(world *ecs.World) {
	log.Info("Gameplay.Resume")
}

// OnStart method
func (st *GameplayState) OnStart(world *ecs.World) {
	log.Info("Gameplay.Start")

	st.TileSystem = ts.TileSystem{}

	// var resources = ecs.GetResource[resources.Resources](world)
	// spriteSheets := resources.SpriteSheetsGame
	resources := ecs.GetResource[resources.Resources](world)
	spriteSheets := resources.SpriteSheets

	spriteSheetBigBackground, ok := (*spriteSheets)["background"]
	if !ok {
		log.Error("SpriteSheet 'game' not found")
		return
	}

	mapper2 := ecs.NewMap2[gc.SpriteRender, gc.Transform](world)
	//mapper3 := ecs.NewMap3[gc.SpriteRender, gc.Transform, gc.MouseReactive](world)

	mapper2.NewEntity(
		&gc.SpriteRender{
			SpriteSheet:  &spriteSheetBigBackground,
			SpriteNumber: 0,
			Options:      ebiten.DrawImageOptions{},
		},
		&gc.Transform{Translation: math.Vector2{X: 0, Y: 0}, Origin: "Middle"},
	)

	td := loader.TextData{
		ID:       "score",
		Text:     "Score",
		FontFace: loader.FontFaceData{Font: "joystix", Options: loader.FontFaceOptions{Size: 25.0}},
		Color:    [4]uint8{255, 0, 0, 255},
	}
	tt := loader.ProcessTextData(world, &td)

	mapperText := ecs.NewMap2[gc.Text, gc.UITransform](world)

	mapperText.NewEntity(
		tt,
		&gc.UITransform{Translation: math.VectorInt2{X: 200, Y: 460}},
	)

	InitTiles(world)

}

func InitTiles(world *ecs.World) {

	resources := ecs.GetResource[resources.Resources](world)
	spriteSheets := resources.SpriteSheets

	spriteSheetTiles, ok := (*spriteSheets)["Tiles"]

	if !ok {
		log.Error("SpriteSheet 'Tiles' not found")
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			mapper := ecs.NewMap4[gc.SpriteRender, gc.Transform, gc.MouseReactive, tc.Tile](world)

			mapper.NewEntity(
				&gc.SpriteRender{SpriteSheet: &spriteSheetTiles, SpriteNumber: 0, Options: ebiten.DrawImageOptions{}},
				&gc.Transform{Translation: math.Vector2{X: float64(j*140) + 180, Y: float64(i*140) + 100}},
				&gc.MouseReactive{ID: fmt.Sprint("test", "", j, " - ", i)},
				&tc.Tile{X: j, Y: i, State: 0},
			)
		}
	}
}

// OnStop method
func (st *GameplayState) OnStop(world *ecs.World) {

	filter := ecs.NewFilter1[gc.SpriteRender](world)
	world.RemoveEntities(filter.Batch(), func(entity ecs.Entity) {
		log.Info("Removing", entity)
	})

	filter2 := ecs.NewFilter1[gc.Text](world)
	world.RemoveEntities(filter2.Batch(), func(entity ecs.Entity) {
		log.Info("Removing", entity)
	})

	log.Info("Gameplay.Stop")

}

// Update method
func (st *GameplayState) Update(world *ecs.World) states.Transition {
	st.TileSystem.Update(world)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return states.Transition{Type: states.TransQuit}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		return states.Transition{Type: states.TransSwitch, NewStates: []states.State{&MainMenuState{}}}
	}

	mapper3 := ecs.NewMap3[gc.SpriteRender, gc.Transform, tc.Tile](world)

	if inpututil.IsKeyJustPressed(ebiten.Key1) {

		resources := ecs.GetResource[resources.Resources](world)
		spriteSheets := resources.SpriteSheets

		spriteSheetTiles, ok := (*spriteSheets)["Tiles"]
		if !ok {
			log.Error("SpriteSheet 'Tiles' not found")
		}

		mapper3.NewEntity(
			&gc.SpriteRender{
				SpriteSheet:  &spriteSheetTiles,
				SpriteNumber: 0,
				Options:      ebiten.DrawImageOptions{},
			},
			&gc.Transform{Translation: math.Vector2{X: -140, Y: -140}, Origin: "Middle", Scale1: math.Vector2{X: -0.75, Y: -0.75}},
			&tc.Tile{X: st.index % 3, Y: st.index / 3, State: (st.index % 2)},
		)
		st.index++
	}

	return states.Transition{}
}
