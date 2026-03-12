package resources

import (
	c "github.com/renenieuw/cubedbits/components"
)

// Resources contains references to data not related to any entity
type Resources struct {
	ScreenDimensions *ScreenDimensions
	SpriteSheets     *map[string]c.SpriteSheet
	Fonts            *map[string]Font
	Prefabs          interface{}
	Game             interface{}
}

// InitResources initializes resources
func InitResources() *Resources {
	return &Resources{}
}
