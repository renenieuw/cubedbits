package uisystem

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/labstack/gommon/log"
	ecs "github.com/mlange-42/ark/ecs"
	c "github.com/renenieuw/cubedbits/components"
	"github.com/renenieuw/cubedbits/resources"
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

		var tochJustPressed bool = false
		var touchIDs []ebiten.TouchID

		touchIDs = inpututil.AppendJustPressedTouchIDs(touchIDs[:0])
		if len(touchIDs) != 0 {
			tochJustPressed = true
			x,y = ebiten.TouchPosition(touchIDs[0])
			//x,y = inpututil.touch (touchIDs[0])
			log.Printf("just toched:%s %d %d %d", tochJustPressed, x, y, len(touchIDs))
		}


		mouseReactive.Hovered = minX <= float64(x) && float64(x) <= maxX && minY <= float64(y) && float64(y) <= maxY
		mouseReactive.JustClicked = mouseReactive.Hovered && (inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || tochJustPressed)

	}

}
