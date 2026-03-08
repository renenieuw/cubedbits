package states

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/labstack/gommon/log"
	"github.com/mlange-42/ark/ecs"
	gc "remapit.visualstudio.com/cubedbits/cubedbitsengine/components"
	tc "remapit.visualstudio.com/cubedbits/cubedbitsengine/games/ticktacktoe/components"
	ts "remapit.visualstudio.com/cubedbits/cubedbitsengine/games/ticktacktoe/systems"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/loader"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/math"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/resources"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/states"
)

// GameplayState is the main game state
type GameplayState struct {
	index int
}

// OnPause method
func (st *GameplayState) OnPause(world *ecs.World) {
	log.Info("Gameplay.OnPause")

}

// OnResume method
func (st *GameplayState) OnResume(world *ecs.World) {
	log.Info("Gameplay.Resume")
}

// OnStart method
func (st *GameplayState) OnStart(world *ecs.World) {
	log.Info("Gameplay.Start")

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
	mapper3 := ecs.NewMap3[gc.SpriteRender, gc.Transform, gc.MouseReactive](world)

	mapper2.NewEntity(
		&gc.SpriteRender{
			SpriteSheet:  &spriteSheetBigBackground,
			SpriteNumber: 0,
			Options:      ebiten.DrawImageOptions{},
		},
		&gc.Transform{Translation: math.Vector2{X: 0, Y: 0}, Origin: "Middle"},
	)

	td := loader.TextData{
		ID:       "gameplay",
		Text:     "gameplay",
		FontFace: loader.FontFaceData{Font: "joystix", Options: loader.FontFaceOptions{Size: 25.0}},
		Color:    [4]uint8{255, 0, 0, 255},
	}
	tt := loader.ProcessTextData(world, &td)

	srd := loader.SpriteRenderData{Fill: &loader.FillData{Width: 40, Height: 40, Color: [4]uint8{255, 0, 0, 255}}}

	srg := loader.ProcessSpriteRenderData(world, &srd)

	mapper3.NewEntity(
		srg,
		&gc.Transform{Translation: math.Vector2{X: 0, Y: 0}, Origin: "Middle"},
		&gc.MouseReactive{ID: "test1"},
	)

	mapperText := ecs.NewMap2[gc.Text, gc.UITransform](world)

	mapperText.NewEntity(
		tt,
		&gc.UITransform{Translation: math.VectorInt2{X: 220, Y: 220}},
	)

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
	ts.TileSystem(world)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return states.Transition{Type: states.TransQuit}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		return states.Transition{Type: states.TransSwitch, NewStates: []states.State{&MainMenuState{}}}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return states.Transition{Type: states.TransSwitch, NewStates: []states.State{&GameOverMenuState{}}}
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
