package states

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mlange-42/ark/ecs"
	gc "remapit.visualstudio.com/cubedbits/cubedbitsengine/components"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/games/ticktacktoe/math"
	m "remapit.visualstudio.com/cubedbits/cubedbitsengine/math"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/states"
)

type menu interface {
	getSelection() int
	setSelection(selection int)
	confirmSelection() states.Transition
	getMenuIDs() []string
	getCursorMenuIDs() []string
}

var menuLastCursorPosition = m.VectorInt2{}

func updateMenu(menu menu, world *ecs.World) states.Transition {
	var transition states.Transition
	selection := menu.getSelection()
	numItems := len(menu.getCursorMenuIDs())

	// Handle keyboard events
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyDown):
		menu.setSelection(math.Mod(selection+1, numItems))
	case inpututil.IsKeyJustPressed(ebiten.KeyUp):
		menu.setSelection(math.Mod(selection-1, numItems))
	case inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeySpace):
		return menu.confirmSelection()
	}

	filter := ecs.NewFilter3[gc.SpriteRender, gc.Transform, gc.MouseReactive](world)

	// Handle mouse events only if mouse is moved or clicked
	x, y := ebiten.CursorPosition()
	if x != menuLastCursorPosition.X || y != menuLastCursorPosition.Y || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		for iElem, id := range menu.getMenuIDs() {
			menuLastCursorPosition = m.VectorInt2{X: x, Y: y}
			query := filter.Query()
			for query.Next() {
				_, _, mouseReactive := query.Get()
				if mouseReactive.ID == id && mouseReactive.Hovered {
					menu.setSelection(iElem)
					if mouseReactive.JustClicked {
						transition = menu.confirmSelection()
					}
				}
			}
		}
	}

	filterText := ecs.NewFilter1[gc.Text](world)

	// Set cursor color
	newSelection := menu.getSelection()
	for iCursor, id := range menu.getCursorMenuIDs() {
		query := filterText.Query()
		for query.Next() {
			text := query.Get()

			if text.ID == id {
				text.Color.A = 0
				text.Color.R = 0
				text.Color.G = 0
				text.Color.B = 0
				if iCursor == newSelection {
					text.Color.R = 255
					text.Color.G = 255
					text.Color.B = 255
				}
			}
		}
	}
	return transition
}
