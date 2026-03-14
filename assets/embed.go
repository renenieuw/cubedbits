package assets

import (
	_ "embed"
)


var (
	//go:embed metadata/spritesheets/spritesheets.toml
	Spritesheets []byte
)

var (
	//go:embed metadata/fonts/fonts.toml
	Fonts []byte
)

var (
	//go:embed textures/background.png
	Background []byte
)

var (
	//go:embed fonts/hack.ttf
	Hack []byte
)

var (
	//go:embed fonts/joystix.ttf
	Joystix []byte
)

var (
	Assets map[string]func(string)[]byte
)

type Menu interface {
	getItems() int
}

func GetAssetByLib(lib string, name string) []byte {
	return Assets[lib](name)
}


func GetAsset(name string) []byte {
	switch name {
	case "Background":
			return Background
	case "hack.ttf":
		return Hack
	case "joystix.ttf":
		return Joystix
	case "fonts.toml":
		return Fonts
	default:
		return Background
	}
}
