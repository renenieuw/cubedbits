package spritesystem

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	ecs "github.com/mlange-42/ark/ecs"
	c "remapit.visualstudio.com/cubedbits/cubedbitsengine/components"
	m "remapit.visualstudio.com/cubedbits/cubedbitsengine/math"
)

func RenderSpriteSystem(world *ecs.World, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, starz4!")

	// col := color.RGBA{0x80, 0x80, 0x80, 0xff}

	// // col := color.RGBA{0x80, 0x80, 0x80, 0xff}
	filter := ecs.NewFilter1[c.SpriteRender](world)

	// // for range 2 {
	// // 	// Get a fresh query and iterate it
	query := filter.Query()
	for query.Next() {
		// 	// Component access through the Query
		sr := query.Get()
		// 	// Update component fields
		// 	// pos.X += vel.DX
		// 	// pos.Y += vel.DY
		// 	//
		//

		drawImageWithWrap(screen, sr)
		// 	vector.FillRect(screen, pos.Position.X, pos.Position.Y, 2, 2, col, false)

	}

	// sprites := world.Manager.Join(world.Components.Engine.SpriteRender, world.Components.Engine.Transform)

	// // Copy query slice into a struct slice for sorting
	// iSprite := 0
	// spritesDepths := make([]spriteDepth, sprites.Size())
	// sprites.Visit(ecs.Visit(func(entity ecs.Entity) {
	// 	spritesDepths[iSprite] = spriteDepth{
	// 		sprite: world.Components.Engine.SpriteRender.Get(entity).(*c.SpriteRender),
	// 		depth:  world.Components.Engine.Transform.Get(entity).(*c.Transform).Depth,
	// 	}
	// 	iSprite++
	// }))

	// // Sort by increasing values of depth
	// sort.Slice(spritesDepths, func(i, j int) bool {
	// 	return spritesDepths[i].depth < spritesDepths[j].depth
	// })

	// // Sprites with higher values of depth are drawn later so they are on top
	// for _, st := range spritesDepths {
	// 	drawImageWithWrap(screen, st.sprite)
	// }
}

// Draw sprite with texture wrapping.
// Image is tiled when texture coordinates are greater than image size.
func drawImageWithWrap(screen *ebiten.Image, spriteRender *c.SpriteRender) {
	sprite := spriteRender.SpriteSheet.Sprites[spriteRender.SpriteNumber]
	texture := spriteRender.SpriteSheet.Texture
	textureWidth, textureHeight := texture.Image.Size()

	startX := int(math.Floor(float64(sprite.X) / float64(textureWidth)))
	startY := int(math.Floor(float64(sprite.Y) / float64(textureHeight)))

	stopX := int(math.Ceil(float64(sprite.X+sprite.Width) / float64(textureWidth)))
	stopY := int(math.Ceil(float64(sprite.Y+sprite.Height) / float64(textureHeight)))

	currentX := 0
	for indX := startX; indX < stopX; indX++ {
		left := m.Max(0, sprite.X-indX*textureWidth)
		right := m.Min(textureWidth, sprite.X+sprite.Width-indX*textureWidth)

		currentY := 0
		for indY := startY; indY < stopY; indY++ {
			top := m.Max(0, sprite.Y-indY*textureHeight)
			bottom := m.Min(textureHeight, sprite.Y+sprite.Height-indY*textureHeight)

			op := spriteRender.Options
			op.GeoM.Translate(float64(currentX), float64(currentY))
			screen.DrawImage(texture.Image.SubImage(image.Rect(left, top, right, bottom)).(*ebiten.Image), &op)

			currentY += bottom - top
		}
		currentX += right - left
	}
}
