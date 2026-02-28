package states

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/labstack/gommon/log"
	"github.com/mlange-42/ark/ecs"
	gc "remapit.visualstudio.com/cubedbits/cubedbitsengine/components"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/loader"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/math"
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

	spriteSheets := ecs.GetResource[loader.SpriteSheetMetadata](world)
	spriteSheet := spriteSheets.SpriteSheets["game"]

	mapper2 := ecs.NewMap2[gc.SpriteRender, gc.Transform](world)

	mapper2.NewEntity(
		&gc.SpriteRender{
			SpriteSheet:  &spriteSheet,
			SpriteNumber: 4,
			Options:      ebiten.DrawImageOptions{},
		},
		&gc.Transform{Translation: math.Vector2{X: 133, Y: 220}},
	)
	// // Init rand seed
	// rand.Seed(time.Now().UnixNano())

	// // Load game and text entities
	// LoadEntities("metadata/start.toml", world)
	// LoadEntities("metadata/text.toml", world)

	// world.Resources.Game = NewGame()
}

// OnStop method
func (st *GameplayState) OnStop(world *ecs.World) {

	filter := ecs.NewFilter1[gc.SpriteRender](world)
	world.RemoveEntities(filter.Batch(), func(entity ecs.Entity) {
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

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		return states.Transition{Type: states.TransSwitch, NewStates: []states.State{&MenuState{}}}
	}

	return states.Transition{}
}
