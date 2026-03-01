package states

import (
	"fmt"

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

type MenuState struct {
	selection int
}

// OnPause method
func (st *MenuState) OnPause(world *ecs.World) {
	log.Info("Menu.OnPause")

}

// OnResume method
func (st *MenuState) OnResume(world *ecs.World) {
	log.Info("Menu.Resume")
}

// OnStart method
func (st *MenuState) OnStart(world *ecs.World) {
	log.Info("Menu.Start")
	log.Info("Menu.Start")

	resources := ecs.GetResource[resources.Resources](world)
	spriteSheets := resources.SpriteSheets
	spriteSheet, ok := (*spriteSheets)["game"]
	if !ok {
		log.Error("SpriteSheet 'game' not found")
		return
	}

	log.Info(fmt.Sprintf("%d", len(spriteSheet.Sprites)))

	mapper2 := ecs.NewMap2[gc.SpriteRender, gc.Transform](world)

	mapper2.NewEntity(
		&gc.SpriteRender{
			SpriteSheet:  &spriteSheet,
			SpriteNumber: 2,
			Options:      ebiten.DrawImageOptions{},
		},
		&gc.Transform{Translation: math.Vector2{X: 133, Y: 220}},
	)

	td := loader.TextData{
		ID:       "menu",
		Text:     "Menu",
		FontFace: loader.FontFaceData{Font: "joystix", Options: loader.FontFaceOptions{Size: 25.0}},
		Color:    [4]uint8{255, 0, 0, 255},
	}
	tt := loader.ProcessTextData(world, &td)
	mapperText := ecs.NewMap2[gc.Text, gc.UITransform](world)

	mapperText.NewEntity(
		tt,
		&gc.UITransform{Translation: math.VectorInt2{X: 22, Y: 22}},
	)

	// // Init rand seed
	// rand.Seed(time.Now().UnixNano())

	// // Load game and text entities
	// LoadEntities("metadata/start.toml", world)
	// LoadEntities("metadata/text.toml", world)

	// world.Resources.Game = NewGame()
}

// OnStop method
func (st *MenuState) OnStop(world *ecs.World) {
	filter := ecs.NewFilter1[gc.SpriteRender](world)
	world.RemoveEntities(filter.Batch(), func(entity ecs.Entity) {
		log.Info("Removing", entity)
	})

	filter2 := ecs.NewFilter1[gc.Text](world)
	world.RemoveEntities(filter2.Batch(), func(entity ecs.Entity) {
		log.Info("Removing", entity)
	})

	log.Info("Menu.Stop")
}

// Update method
func (st *MenuState) Update(world *ecs.World) states.Transition {
	// DemoSystem(world)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return states.Transition{Type: states.TransQuit}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		return states.Transition{Type: states.TransSwitch, NewStates: []states.State{&GameplayState{}}}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return states.Transition{Type: states.TransSwitch, NewStates: []states.State{&GameOverMenuState{}}}
	}

	return states.Transition{}
}
