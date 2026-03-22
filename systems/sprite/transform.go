package spritesystem

import (
	"log/slog"

	ecs "github.com/mlange-42/ark/ecs"
	c "github.com/renenieuw/cubedbits/components"
	"github.com/renenieuw/cubedbits/resources"
)

// TransformSystem updates geometry matrix.
// Geometry matrix is first recentered, then scaled and rotated, and finally translated.
func TransformSystem(world *ecs.World) {

	logger := slog.Default().With("Context","Sprite.TransformSystem")
	logger.Debug("TransformSystem Start")

	filter := ecs.NewFilter2[c.SpriteRender, c.Transform](world)
	query := filter.Query()
	for query.Next() {
		sprite, transform := query.Get()

		if(sprite.SpriteSheet.Sprites[sprite.SpriteGroup] == nil) {
			logger.Debug("sprite.SpriteGroup not found", "SpriteGroup",sprite.SpriteGroup, "SpriteNumber",sprite.SpriteNumber)
		} else if (len(sprite.SpriteSheet.Sprites[sprite.SpriteGroup]) < (sprite.SpriteNumber + 1)) {
			logger.Debug("sprite not found", "SpriteGroup",sprite.SpriteGroup, "SpriteNumber",sprite.SpriteNumber)
		}

		spriteWidth := float64(sprite.SpriteSheet.Sprites[sprite.SpriteGroup][sprite.SpriteNumber].Width)
		spriteHeight := float64(sprite.SpriteSheet.Sprites[sprite.SpriteGroup][sprite.SpriteNumber].Height)

		// Reset geometry matrix
		sprite.Options.GeoM.Reset()

		// Center sprite on top left pixel
		sprite.Options.GeoM.Translate(-spriteWidth/2, -spriteHeight/2)

		// Perform scale
		sprite.Options.GeoM.Scale(transform.Scale1.X+1, transform.Scale1.Y+1)

		// Perform rotation
		sprite.Options.GeoM.Rotate(-transform.Rotation)

		resources := ecs.GetResource[resources.Resources](world)
		//		sd := ecs.NewResource[resources.ScreenDimensions](world)
		screenDimensions := resources.ScreenDimensions

		// Perform translation
		screenWidth := float64(screenDimensions.Width)
		screenHeight := float64(screenDimensions.Height)

		offsetX, offsetY := transform.ComputeOriginOffset(screenWidth, screenHeight)
		sprite.Options.GeoM.Translate(transform.Translation.X+offsetX, screenHeight-transform.Translation.Y-offsetY)
	}
}
