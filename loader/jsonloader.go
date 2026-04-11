package loader

import (
	"bytes"
	"image"
	_ "image/png"
	"log/slog"
	"sort"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/renenieuw/cubedbits/assets"
	c "github.com/renenieuw/cubedbits/components"
	"github.com/renenieuw/cubedbits/libraries/texturepacker"
)


func LoadSpriteSheetsFromJson(data []byte, lib string, imgName string) (map[string]c.SpriteSheet, error) {
	logger := slog.Default().With("Context","Loader.LoadSpritesheets")
	logger.Debug("LoadSpriteSheetsFromJson: ", "lib", lib, "imgName", imgName)

	sheet, err := texturepacker.SheetFromData(data, texturepacker.FormatJSONHash{})
	if err != nil {
		slog.Error("LoadSpriteSheetsFromJson error","error", err, "Object", )
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(assets.GetAssetByLib(lib,string(imgName)) ))
	if err != nil {
		slog.Error("Failed to decode sprite sheet: %s %s", "lib",lib, "imagename", imgName, "Error", err)
		return nil, err
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
		logger.Info("sprite", "s", spr)
		sprites[group] = append(sprites[group], spr)

		sort.Slice(sprites[group], func(i, j int) bool {
			return sprites[group][i].Name < sprites[group][i].Name
		})

		logger.Debug("Loaded sprite: ", "lib", lib, "imgName", imgName, "name", filename, "spritegroup", group, "found", found)
	}

	for _, sprites := range sprites {
		sort.Slice(sprites, func(i, j int) bool {
			return sprites[i].Name < sprites[j].Name
		})
	}

	var tex c.Texture
	tex.Image =  textureImage

	retVal[lib] = c.SpriteSheet{ Texture: tex, Sprites: sprites}

	logger.Debug("Loaded sprite sheet: ", "lib", lib, "imgName", imgName, "data", retVal)
	return retVal, nil
}
