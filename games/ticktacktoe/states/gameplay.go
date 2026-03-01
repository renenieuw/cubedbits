package states

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/labstack/gommon/log"
	"github.com/mlange-42/ark/ecs"
	gc "remapit.visualstudio.com/cubedbits/cubedbitsengine/components"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/loader"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/math"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/resources"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/states"
)

// GameplayState is the main game state
type GameplayState struct{}

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

	spriteSheet, ok := (*spriteSheets)["game"]
	if !ok {
		log.Error("SpriteSheet 'game' not found")
		return
	}

	spriteSheet2, ok := (*spriteSheets)["GameEngineBackground"]
	if !ok {
		log.Error("SpriteSheet 'game' not found")
		return
	}

	mapper2 := ecs.NewMap2[gc.SpriteRender, gc.Transform](world)

	mapper2.NewEntity(
		&gc.SpriteRender{
			SpriteSheet:  &spriteSheet,
			SpriteNumber: 3,
			Options:      ebiten.DrawImageOptions{},
		},
		&gc.Transform{Translation: math.Vector2{X: 133, Y: 220}},
	)

	mapper2.NewEntity(
		&gc.SpriteRender{
			SpriteSheet:  &spriteSheet2,
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

	mapperText := ecs.NewMap2[gc.Text, gc.UITransform](world)

	mapperText.NewEntity(
		tt,
		&gc.UITransform{Translation: math.VectorInt2{X: 220, Y: 220}},
	)

	mapper := ecs.NewMap1[gc.SpriteRender](world)
	// mapperText := ecs.NewMap1[gc.Text](w)

	for range 30 {
		_ = mapper.NewEntity(

			&gc.SpriteRender{
				SpriteSheet:  &spriteSheet,
				SpriteNumber: 3,
				Options:      ebiten.DrawImageOptions{},
			},
		)
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

	// world.Resources.Game = nil
	// world.Manager.DeleteAllEntities()
}

// Update method
func (st *GameplayState) Update(world *ecs.World) states.Transition {
	// DemoSystem(world)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return states.Transition{Type: states.TransQuit}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		return states.Transition{Type: states.TransSwitch, NewStates: []states.State{&MenuState{}}}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return states.Transition{Type: states.TransSwitch, NewStates: []states.State{&GameOverMenuState{}}}
	}

	return states.Transition{}
}
