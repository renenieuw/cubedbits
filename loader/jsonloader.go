package loader

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"log/slog"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/renenieuw/cubedbits/assets"
	c "github.com/renenieuw/cubedbits/components"
	"github.com/renenieuw/cubedbits/libraries/texturepacker"
	// "github.com/renenieuw/cubedbits/loader"
	//"github.com/renenieuw/cubedbits/math"
)

// type spriteSheetMetadata struct {
// 	SpriteSheets map[string]c.SpriteSheet
// }

// LoadSpriteSheets loads sprite sheets from a TOML file
// func LoadSpriteSheetsFromJson(spriteSheetMetadataPath string) map[string]c.SpriteSheet {
// 	var spriteSheetMetadata spriteSheetMetadata
// 	utils.Try(toml.DecodeFile(spriteSheetMetadataPath, &spriteSheetMetadata))
// 	return spriteSheetMetadata.SpriteSheets
// }

func LoadSpriteSheetsFromJson(data []byte, lib string, imgName string) map[string]c.SpriteSheet {
	logger := slog.Default().With("Context","JsonLoader")
	logger.Info("LoadSpriteSheetsFromJsonInfo: ", "lib", lib, "imgName", imgName)
	logger.Debug("LoadSpriteSheetsFromJsonDebug: ", "lib", lib, "imgName", imgName)

	sheet, err := texturepacker.SheetFromData(data, texturepacker.FormatJSONHash{})
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(assets.GetAssetByLib(lib,string(imgName)) ))
	if err != nil {
		slog.Debug("Failed to decode sprite sheet: %s %s", lib, imgName)
		log.Fatal(err)
	}
	textureImage := ebiten.NewImageFromImage(img)

	var retVal = make(map[string]c.SpriteSheet)
	var sprites = make(map[string][]c.Sprite)

	for name, sprite := range sheet.Sprites {
		group, filename, found := strings.Cut(name, "/")
		var spr c.Sprite
		if(!found) {
			filename = name;
			group = filename
		}
		spr = c.Sprite{X: sprite.Frame.Min.X, Y: sprite.Frame.Min.Y, Width: sprite.Frame.Dx(), Height: sprite.Frame.Dy(), Name: filename, SpriteGroup: group}
		sprites[group] = append(sprites[group], spr)
		logger.Debug("Loaded sprite: ", "lib", lib, "imgName", imgName, "name", filename, "spritegroup", group, "found", found)
	}

	var tex c.Texture
	tex.Image =  textureImage

	retVal[lib] = c.SpriteSheet{ Texture: tex, Sprites: sprites}

	logger.Debug("Loaded sprite sheet: ", "lib", lib, "imgName", imgName, "data", retVal)
	return retVal
}
