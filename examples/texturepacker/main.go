package main

import (
	"fmt"
	_ "image/png"
	"log/slog"

	"github.com/renenieuw/cubedbits/loader"
)

func main() {
	logger := slog.Default().With("Context", "TexturePacker")

	//	sse := loader.LoadSpriteSheetsFromJson("../../assets/metadata/spritesheets/spritesheets.toml")
	sse, _ := loader.LoadSpriteSheetsFromJson("../../games/roam/Assets/mainsprites-0.json")
	logger.Debug(fmt.Sprintf("count %d", sse["ice.png"].Sprites[0].X))
	logger.Debug(fmt.Sprintf("count %d", sse["ice.png"].Sprites[0].Y))
	logger.Debug(fmt.Sprintf("count %d", sse["ice.png"].Sprites[0].Width))
	logger.Debug(fmt.Sprintf("count %d", sse["ice.png"].Sprites[0].Height))

	//	log.Printf("count %d", len(sse[0].Sprites))

}
