package main

import (
	_ "image/png"
	"log"

	"remapit.visualstudio.com/cubedbits/cubedbitsengine/loader"
)

func main() {

	//	sse := loader.LoadSpriteSheetsFromJson("../../assets/metadata/spritesheets/spritesheets.toml")
	sse := loader.LoadSpriteSheetsFromJson("../../games/roam/Assets/mainsprites-0.json")
	log.Printf("count %d", sse["ice.png"].Sprites[0].X)
	log.Printf("count %d", sse["ice.png"].Sprites[0].Y)
	log.Printf("count %d", sse["ice.png"].Sprites[0].Width)
	log.Printf("count %d", sse["ice.png"].Sprites[0].Height)

	//	log.Printf("count %d", len(sse[0].Sprites))

}
