package loader

import (
	"log/slog"

	c "github.com/renenieuw/cubedbits/components"
	"github.com/renenieuw/cubedbits/utils"

	"github.com/BurntSushi/toml"
)

type spriteSheetMetadata struct {
	SpriteSheets map[string]c.SpriteSheet `toml:"sprite_sheet"`
}

// LoadSpriteSheets loads sprite sheets from a TOML file
func LoadSpriteSheets(spriteSheetMetadataPath string) map[string]c.SpriteSheet {
	var spriteSheetMetadata spriteSheetMetadata
	utils.Try(toml.DecodeFile(spriteSheetMetadataPath, &spriteSheetMetadata))
	return spriteSheetMetadata.SpriteSheets
}

func LoadSpriteSheetsFromString(spriteSheetMetadataString string) map[string]c.SpriteSheet {
	logger := slog.Default().With("Context","Loader")
	logger.Debug("LoadSpriteSheetsFromString", "file", spriteSheetMetadataString)
	var spriteSheetMetadata spriteSheetMetadata
	utils.Try(toml.Decode(spriteSheetMetadataString, &spriteSheetMetadata))
	return spriteSheetMetadata.SpriteSheets
}
