package spritesystem

import (
	ecs "github.com/mlange-42/ark/ecs"
	c "remapit.visualstudio.com/cubedbits/cubedbitsengine/components"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/resources"
)

// TransformSystem updates geometry matrix.
// Geometry matrix is first recentered, then scaled and rotated, and finally translated.
func TransformSystem(world *ecs.World) {

	filter := ecs.NewFilter2[c.SpriteRender, c.Transform](world)
	query := filter.Query()
	for query.Next() {
		sprite, transform := query.Get()

		// world.Manager.Join(world.Components.Engine.SpriteRender, world.Components.Engine.Transform).Visit(ecs.Visit(func(entity ecs.Entity) {
		// 	sprite := world.Components.Engine.SpriteRender.Get(entity).(*c.SpriteRender)
		// 	transform := world.Components.Engine.Transform.Get(entity).(*c.Transform)

		spriteWidth := float64(sprite.SpriteSheet.Sprites[sprite.SpriteNumber].Width)
		spriteHeight := float64(sprite.SpriteSheet.Sprites[sprite.SpriteNumber].Height)

		// Reset geometry matrix
		sprite.Options.GeoM.Reset()

		// Center sprite on top left pixel
		sprite.Options.GeoM.Translate(-spriteWidth/2, -spriteHeight/2)

		// Perform scale
		sprite.Options.GeoM.Scale(transform.Scale1.X+1, transform.Scale1.Y+1)

		// Perform rotation
		sprite.Options.GeoM.Rotate(-transform.Rotation)

		sd := ecs.NewResource[resources.ScreenDimensions](world)
		screenDimensions := sd.Get()

		// Perform translation
		screenWidth := float64(screenDimensions.Width)
		screenHeight := float64(screenDimensions.Height)

		offsetX, offsetY := transform.ComputeOriginOffset(screenWidth, screenHeight)
		sprite.Options.GeoM.Translate(transform.Translation.X+offsetX, screenHeight-transform.Translation.Y-offsetY)
	}
}
