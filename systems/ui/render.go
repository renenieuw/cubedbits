package uisystem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	ecs "github.com/mlange-42/ark/ecs"
	c "remapit.visualstudio.com/cubedbits/cubedbitsengine/components"
)

// RenderUISystem draws text entities
func RenderUISystem(world *ecs.World, screen *ebiten.Image) {
	filter := ecs.NewFilter1[c.Text](world)
	query := filter.Query()
	for query.Next() {
		sr := query.Get()
		text.Draw(screen, sr.Text, sr.FontFace, 222, 222, sr.Color)

	}
	// col := color.RGBA{0x80, 0x80, 0x80, 0xff}
	// fonts := ecs.GetResource[loader.FontMetadata](world)
	// font := fonts.Fonts["game"]
	// world.Manager.Join(world.Components.Engine.Text, world.Components.Engine.UITransform).Visit(ecs.Visit(func(entity ecs.Entity) {
	//textData := world.Components.Engine.Text.Get(entity).(*c.Text)
	// 	uiTransform := world.Components.Engine.UITransform.Get(entity).(*c.UITransform)

	// 	// Compute dot offset
	// 	x, y := utils.Try2(c.ComputeDotOffset(textData.Text, textData.FontFace, uiTransform.Pivot))

	// 	// Draw text
	// 	screenWidth := world.Resources.ScreenDimensions.Width
	// 	screenHeight := world.Resources.ScreenDimensions.Height

	// 	offsetX, offsetY := uiTransform.ComputeOriginOffset(screenWidth, screenHeight)
	// 	text.Draw(screen, textData.Text, textData.FontFace, uiTransform.Translation.X+offsetX-x, screenHeight-uiTransform.Translation.Y-offsetY-y, textData.Color)
	// }))
	//
	//	text.Draw(screen, "text.text", font.Font, 222, 222, col)
}
