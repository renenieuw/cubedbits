package assets

import (
	_ "embed"
)


var (
	//go:embed metadata/spritesheets/spritesheets.toml
	Spritesheets []byte
)

var (
	//go:embed textures/Background.png
	Background []byte
)

var (
	//go:embed textures/Tiles.png
	Tiles []byte
)

func GetAsset(name string) []byte {
	switch name {
	case "background.png":
			return Background
	case "Tiles.png":
			return Tiles
	default:
		return Tiles
	}
}
