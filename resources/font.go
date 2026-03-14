package resources

import (
	"strings"
	"github.com/renenieuw/cubedbits/assets"

	"github.com/golang/freetype/truetype"
	"github.com/renenieuw/cubedbits/utils"
)

// Font structure
type Font struct {
	Font *truetype.Font
}

// UnmarshalTOML fills structure fields from TOML data
func (f *Font) UnmarshalTOML(i interface{}) error {
	path := i.(map[string]interface{})["font"].(string)
	lib, asset, _ := strings.Cut(path, "/")

	fontFile := assets.GetAssetByLib(lib,string(asset));

	f.Font = utils.Try(truetype.Parse(fontFile))
	return nil
}
