package states

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/labstack/gommon/log"
	"github.com/mlange-42/ark/ecs"
	gc "remapit.visualstudio.com/cubedbits/cubedbitsengine/components"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/states"
)

type GameOverMenuState struct {
	selection int
}

// OnPause method
func (st *GameOverMenuState) OnPause(world *ecs.World) {
	log.Info("Menu.OnPause")

}

// OnResume method
func (st *GameOverMenuState) OnResume(world *ecs.World) {
	log.Info("Menu.Resume")
}

// OnStart method
func (st *GameOverMenuState) OnStart(world *ecs.World) {
	log.Info("Menu.Start")
	// // Init rand seed
	// rand.Seed(time.Now().UnixNano())

	// // Load game and text entities
	// LoadEntities("metadata/start.toml", world)
	// LoadEntities("metadata/text.toml", world)

	// world.Resources.Game = NewGame()
}

// OnStop method
func (st *GameOverMenuState) OnStop(world *ecs.World) {
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
func (st *GameOverMenuState) Update(world *ecs.World) states.Transition {
	// DemoSystem(world)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return states.Transition{Type: states.TransQuit}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		return states.Transition{Type: states.TransSwitch, NewStates: []states.State{&GameplayState{}}}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		return states.Transition{Type: states.TransSwitch, NewStates: []states.State{&MainMenuState{}}}
	}

	return states.Transition{}
}
