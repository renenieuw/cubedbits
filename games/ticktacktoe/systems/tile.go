package systems

import (
	"github.com/mlange-42/ark/ecs"
	c "remapit.visualstudio.com/cubedbits/cubedbitsengine/components"
	tc "remapit.visualstudio.com/cubedbits/cubedbitsengine/games/ticktacktoe/components"
)

func TileSystem(world *ecs.World) {
	filter := ecs.NewFilter3[c.SpriteRender, tc.Tile, c.Transform](world)
	query := filter.Query()
	for query.Next() {
		spriteRender, tile, transform := query.Get()
		spriteRender.SpriteNumber = tile.State
		transform.Translation.X = float64(-150 + (tile.X * 150))
		transform.Translation.Y = float64((150 - (tile.Y * 150)))
	}
}
