package loader

import (
	"github.com/renenieuw/cubedbits/resources"
	"github.com/renenieuw/cubedbits/utils"

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
