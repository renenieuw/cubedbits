package resources

import (
	c "remapit.visualstudio.com/cubedbits/cubedbitsengine/components"
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
