package loader

import (
	"log/slog"
	"strings"
	"fmt"
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
	logger := slog.Default().With("Context", "Loader.Fonts")

	lib, asset, _ := strings.Cut(fontPath, "/")
	logger.Debug(fmt.Sprintf("loading: %s %s %s" , fontPath, lib, asset))
	fontFile := assets.GetAssetByLib(lib,string(asset));
	logger.Debug(fmt.Sprintf("loading: %d" , len(fontFile)))



	var fontMetadata FontMetadata
	toml.Unmarshal(fontFile, &fontMetadata)
	return fontMetadata
}
