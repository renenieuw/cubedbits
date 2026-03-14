package loader

import (
	"log"
	"strings"
	"github.com/renenieuw/cubedbits/assets"

	"github.com/renenieuw/cubedbits/resources"
	// "github.com/renenieuw/cubedbits/utils"

	"github.com/BurntSushi/toml"
)

type FontMetadata struct {
	Fonts map[string]resources.Font `toml:"font"`
}

// LoadFonts loads fonts from a TOML file
func LoadFonts(fontPath string) FontMetadata {
	lib, asset, _ := strings.Cut(fontPath, "/")
	log.Printf("loading: %s %s %s" , fontPath, lib, asset)
	fontFile := assets.GetAssetByLib(lib,string(asset));
	log.Printf("loading: %d" , len(fontFile))



	var fontMetadata FontMetadata
	toml.Unmarshal(fontFile, &fontMetadata)
	return fontMetadata
}
