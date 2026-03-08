package uisystem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	ecs "github.com/mlange-42/ark/ecs"
	c "remapit.visualstudio.com/cubedbits/cubedbitsengine/components"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/resources"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/utils"
)

// RenderUISystem draws text entities
func RenderUISystem(world *ecs.World, screen *ebiten.Image) {
	resources := ecs.GetResource[resources.Resources](world)
	filter := ecs.NewFilter2[c.Text, c.UITransform](world)
	query := filter.Query()
	for query.Next() {
		textData, uiTransform := query.Get()

		// Compute dot offset
		x, y := utils.Try2(c.ComputeDotOffset(textData.Text, textData.FontFace, uiTransform.Pivot))
		// Draw text
		screenWidth := resources.ScreenDimensions.Width
		screenHeight := resources.ScreenDimensions.Height

		offsetX, offsetY := uiTransform.ComputeOriginOffset(screenWidth, screenHeight)
		text.Draw(screen, textData.Text, textData.FontFace, uiTransform.Translation.X+offsetX-x, screenHeight-uiTransform.Translation.Y-offsetY-y, textData.Color)
	}
}
