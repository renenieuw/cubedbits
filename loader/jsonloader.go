package loader

import (
	_ "image/png"
	"log"

	c "github.com/renenieuw/cubedbits/components"
	"github.com/renenieuw/cubedbits/libraries/texturepacker"
	"image"
	"bytes"
	"github.com/renenieuw/cubedbits/assets"
	"github.com/hajimehoshi/ebiten/v2"

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

func LoadSpriteSheetsFromJson(data []byte, lib string, spritesheetName string) map[string]c.SpriteSheet {

	assets.Assets = map[string]func(string)[]byte{
        "cubedBits": assets.GetAsset,
    }

	sheet, err := texturepacker.SheetFromData(data, texturepacker.FormatJSONHash{})
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(assets.GetAssetByLib(lib,string(spritesheetName)) ))
	if err != nil {
		log.Fatal(err)
	}
	textureImage := ebiten.NewImageFromImage(img)

	var retVal = make(map[string]c.SpriteSheet)
	var sprites []c.Sprite

	for _, sprite := range sheet.Sprites {
		spr := c.Sprite{X: sprite.Frame.Min.X, Y: sprite.Frame.Min.Y, Width: sprite.Frame.Dx(), Height: sprite.Frame.Dy()}
		sprites = append(sprites, spr)
	}

	var tex c.Texture
	tex.Image =  textureImage

	retVal[spritesheetName] = c.SpriteSheet{ Texture: tex, Sprites: sprites}
	return retVal
}
