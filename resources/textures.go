package resources

import (
	"fmt"
	"log/slog"
	"maps"
	"strings"
	"sync/atomic"

	"github.com/renenieuw/cubedbits/components"
)

var defaultTextures atomic.Pointer[Textures]


type Textures struct {
	Spritesheets map[string]components.SpriteSheet
}

func SetDefault(t *Textures){
	defaultTextures.Store(t)
}

func Default() *Textures {
	return defaultTextures.Load()
}

func (t *Textures) Dump() {
	logger := slog.Default().With("Context","GetTexture")
	logger.Debug("DumpTextures","Dump", t.Spritesheets)
}

func (t *Textures) GetSpriteRender(path string) *components.SpriteRender  {

	s1 := SplitPath(path)

	rootPath := s1[0]
	spriteName := strings.Join(s1[1:], "/")


	logger := slog.Default().With("Context","GetTexture")
	logger.Debug(fmt.Sprintf("path %s", s1[0]))
	logger.Debug(fmt.Sprintf("filenamer %s", spriteName))

	for key, s := range t.Spritesheets {

		logger.Debug(fmt.Sprintf("Searching %s for %s %s", key, rootPath, spriteName), "spritesheet", s)
		found := s.Sprites[s1[0]] != nil
		logger.Debug(fmt.Sprintf("Found %t", found))
		if(found){
			for index, sp := range s.Sprites[s1[0]] {
				if(sp.Name == spriteName) {
					logger.Debug(fmt.Sprintf("Found complete %s", sp.Name))
					sr := components.SpriteRender {
						SpriteSheet: &s,
						SpriteNumber: index,
						SpriteGroup: sp.SpriteGroup,
					}
					return &sr;
				}
			}
		}
	}
	return nil
}

func (t *Textures) AddSpritesheet(spriteSheet map[string]components.SpriteSheet) {
	if(t.Spritesheets == nil) {
		t.Spritesheets = spriteSheet
	} else {
		maps.Copy(t.Spritesheets, spriteSheet)
	}

}



func SplitPath(path string) ([]string) {

	// Split the directory part into individual directory names
	parts := strings.Split(path, "/")


	return parts
}
