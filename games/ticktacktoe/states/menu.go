package states

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/labstack/gommon/log"
	"github.com/mlange-42/ark/ecs"
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
	// // Init rand seed
	// rand.Seed(time.Now().UnixNano())

	// // Load game and text entities
	// LoadEntities("metadata/start.toml", world)
	// LoadEntities("metadata/text.toml", world)

	// world.Resources.Game = NewGame()
}

// OnStop method
func (st *MenuState) OnStop(world *ecs.World) {
	log.Info("Menu.Stop")

	// world.Resources.Game = nil
	// world.Manager.DeleteAllEntities()
}

// Update method
func (st *MenuState) Update(world *ecs.World) states.Transition {
	// DemoSystem(world)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return states.Transition{Type: states.TransQuit}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		return states.Transition{Type: states.TransSwitch, NewStates: []states.State{&GameplayState{}}}
	}

	return states.Transition{}
}
