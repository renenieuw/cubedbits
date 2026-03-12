package loader

import (
	_ "image/png"
	"log"
	"os"

	c "github.com/renenieuw/cubedbits/components"
	"github.com/renenieuw/cubedbits/libraries/texturepacker"

	// "github.com/renenieuw/cubedbits/loader"
	"github.com/renenieuw/cubedbits/math"
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

func LoadSpriteSheetsFromJson(spriteSheetMetadataPath string) map[string]c.SpriteSheet {

	math.Max(1, 2)
	sheet, err := texturepacker.SheetFromFile(spriteSheetMetadataPath, texturepacker.FormatJSONHash{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", "../../games/roam/Assets/"+sheet.Meta.Image)
	imageFile, err := os.Open("../../games/roam/Assets/" + sheet.Meta.Image)
	if err != nil {
		log.Fatal(err)
	}
	defer imageFile.Close()
	//img, _, err := image.Decode(imageFile)
	//	sheetImage, ok := img. (image.RGBA)
	//if !ok {
	//log.Fatal("expected RGBA image")
	//	}
	var retVal = make(map[string]c.SpriteSheet)

	for name, sprite := range sheet.Sprites {
		//		spriteImage := sheetImage.SubImage(sprite.Frame)
		log.Printf("%s %t", name, sprite.Rotated)
		retVal[name] = c.SpriteSheet{Sprites: []c.Sprite{c.Sprite{X: sprite.Frame.Min.X, Y: sprite.Frame.Min.Y, Width: sprite.Frame.Dx(), Height: sprite.Frame.Dy()}}}
	}

	//	retVal["twee"] = c.SpriteSheet{Sprites: []c.Sprite{}}
	//	retVal["drie"] = c.SpriteSheet{Sprites: []c.Sprite{}}
	return retVal
}
