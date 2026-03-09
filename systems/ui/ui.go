package uisystem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	ecs "github.com/mlange-42/ark/ecs"
	c "remapit.visualstudio.com/cubedbits/cubedbitsengine/components"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/resources"
)

// UISystem sets mouse reactive components
func UISystem(world *ecs.World) {
	filter := ecs.NewFilter3[c.SpriteRender, c.Transform, c.MouseReactive](world)
	query := filter.Query()
	for query.Next() {
		sprite, transform, mouseReactive := query.Get()

		resources := ecs.GetResource[resources.Resources](world)

		screenWidth := float64(resources.ScreenDimensions.Width)
		screenHeight := float64(resources.ScreenDimensions.Height)

		spriteWidth := float64(sprite.SpriteSheet.Sprites[sprite.SpriteNumber].Width)
		spriteHeight := float64(sprite.SpriteSheet.Sprites[sprite.SpriteNumber].Height)

		offsetX, offsetY := transform.ComputeOriginOffset(screenWidth, screenHeight)

		minX := (offsetX + transform.Translation.X) - spriteWidth/2
		maxX := (offsetX + transform.Translation.X) + spriteWidth/2
		minY := screenHeight - (offsetY + transform.Translation.Y) - spriteHeight/2
		maxY := screenHeight - (offsetY + transform.Translation.Y) + spriteHeight/2

		x, y := ebiten.CursorPosition()

		mouseReactive.Hovered = minX <= float64(x) && float64(x) <= maxX && minY <= float64(y) && float64(y) <= maxY
		mouseReactive.JustClicked = mouseReactive.Hovered && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)

		//	log.Printf("%s %t", mouseReactive.ID, mouseReactive.Hovered)
	}

}
