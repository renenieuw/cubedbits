package assets

import (
	_ "embed"
)


// var (
// 	//go:embed metadata/spritesheets/spritesheets.toml
// 	Spritesheets []byte
// )

// var (
// 	//go:embed textures/Background.png
// 	Background []byte
// )

// var (
// 	//go:embed textures/Tiles.png
// 	Tiles []byte
// )
//
//
//
var (
	//go:embed metadata/spritesheets/tictactoe.json
	TictactoeJson []byte
)

var (
	//go:embed textures/tictactoe.png
	TictactoeImg []byte
)

func GetAsset(name string) []byte {
	switch name {
	case "tictactoe.png":
			return TictactoeImg
	case "tictactoe.json":
		return TictactoeJson
	default:
		return nil
	}
}
