package states

import (
	"github.com/labstack/gommon/log"
	"github.com/mlange-42/ark/ecs"
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
	// // Init rand seed
	// rand.Seed(time.Now().UnixNano())

	// // Load game and text entities
	// LoadEntities("metadata/start.toml", world)
	// LoadEntities("metadata/text.toml", world)

	// world.Resources.Game = NewGame()
}

// OnStop method
func (st *GameplayState) OnStop(world *ecs.World) {
	log.Info("Gameplay.Stop")
	// world.Resources.Game = nil
	// world.Manager.DeleteAllEntities()
}

// Update method
func (st *GameplayState) Update(world *ecs.World) states.Transition {
	// DemoSystem(world)

	// if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
	// 	return states.Transition{Type: states.TransQuit}
	// }
	return states.Transition{}
}
