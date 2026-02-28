package resources

import (
	"os"

	"github.com/golang/freetype/truetype"
	"remapit.visualstudio.com/cubedbits/cubedbitsengine/utils"
)

// Font structure
type Font struct {
	Font *truetype.Font
}

// UnmarshalTOML fills structure fields from TOML data
func (f *Font) UnmarshalTOML(i interface{}) error {
	fontFile := utils.Try(os.ReadFile(i.(map[string]interface{})["font"].(string)))
	f.Font = utils.Try(truetype.Parse(fontFile))
	return nil
}
