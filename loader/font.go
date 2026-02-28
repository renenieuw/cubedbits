package loader

import (
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/resources"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/utils"

	"github.com/BurntSushi/toml"
)

type FontMetadata struct {
	Fonts map[string]resources.Font `toml:"font"`
}

// LoadFonts loads fonts from a TOML file
func LoadFonts(fontPath string) FontMetadata {
	var fontMetadata FontMetadata
	utils.Try(toml.DecodeFile(fontPath, &fontMetadata))
	return fontMetadata
}
