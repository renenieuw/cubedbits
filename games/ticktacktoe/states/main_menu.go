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
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/states"
)

type MainMenuState struct {
	selection int
	Ids       []string
	CursorIds []string
}

// OnPause method
func (st *MainMenuState) OnPause(world *ecs.World) {
	log.Info("Menu.OnPause")

}

// OnResume method
func (st *MainMenuState) OnResume(world *ecs.World) {
	log.Info("Menu.Resume")
}

// OnStart method
func (st *MainMenuState) OnStart(world *ecs.World) {
	createMenuItem("start", "Start", math.VectorInt2{X: 320, Y: 300}, world)
	createMenuItem("quit", "Quit", math.VectorInt2{X: 320, Y: 200}, world)

	filter := ecs.NewFilter1[gc.MenuItem](world)
	query := filter.Query()
	//	st.Ids = make([]string, query.Count())
	for query.Next() {
		mi := query.Get()
		st.Ids = append(st.Ids, mi.ID)
		st.CursorIds = append(st.CursorIds, "cursor_"+mi.ID)
	}

}

func createMenuItem(id string, text string, location math.VectorInt2, world *ecs.World) {

	mapperText := ecs.NewMap3[gc.Text, gc.UITransform, gc.MenuItem](world)
	mapperText2 := ecs.NewMap2[gc.Text, gc.UITransform](world)
	mapper3 := ecs.NewMap3[gc.SpriteRender, gc.Transform, gc.MouseReactive](world)

	srd := loader.SpriteRenderData{Fill: &loader.FillData{Width: 100, Height: 40}}
	srg := loader.ProcessSpriteRenderData(world, &srd)

	mapper3.NewEntity(
		srg,
		&gc.Transform{Translation: math.Vector2{X: float64(location.X), Y: float64(location.Y)}, Origin: "BottomLeft"},
		&gc.MouseReactive{ID: id},
	)

	td := loader.TextData{
		ID:       id,
		Text:     text,
		FontFace: loader.FontFaceData{Font: "joystix", Options: loader.FontFaceOptions{Size: 40.0}},
		Color:    [4]uint8{255, 0, 0, 255},
	}
	tt := loader.ProcessTextData(world, &td)

	tdc := loader.TextData{
		ID:       id,
		Text:     text,
		FontFace: loader.FontFaceData{Font: "hack", Options: loader.FontFaceOptions{Size: 40.0}},
		Color:    [4]uint8{255, 0, 0, 255},
	}

	mapperText.NewEntity(
		tt,
		&gc.UITransform{Translation: math.VectorInt2{X: location.X, Y: location.Y}},
		&gc.MenuItem{ID: id},
	)

	tdc.Text = "\u25ba"
	tdc.ID = "cursor_" + id

	tt = loader.ProcessTextData(world, &tdc)

	mapperText2.NewEntity(
		tt,
		&gc.UITransform{Translation: math.VectorInt2{X: location.X - 100, Y: location.Y}},
	)

}

// OnStop method
func (st *MainMenuState) OnStop(world *ecs.World) {
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

func (st *MainMenuState) confirmSelection() states.Transition {
	log.Printf("confirmSelection:%d", st.selection)

	switch st.selection {
	case 0:
		// Resume
		return states.Transition{Type: states.TransSwitch, NewStates: []states.State{&GameplayState{}}}
	case 1:
		// Game
		return states.Transition{Type: states.TransQuit}
	}
	panic(fmt.Errorf("unknown selection: %d", st.selection))
}

// Update method
func (st *MainMenuState) Update(world *ecs.World) states.Transition {
	// DemoSystem(world)

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return states.Transition{Type: states.TransQuit}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		return states.Transition{Type: states.TransSwitch, NewStates: []states.State{&GameplayState{}}}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return states.Transition{Type: states.TransQuit}
	}

	return updateMenu(st, world)

	// return states.Transition{}
}

func (st *MainMenuState) getMenuIDs() []string {
	return st.Ids
}

func (st *MainMenuState) getCursorMenuIDs() []string {
	return st.CursorIds
}

func (st *MainMenuState) getSelection() int {
	return st.selection
}

func (st *MainMenuState) setSelection(selection int) {
	st.selection = selection
}
