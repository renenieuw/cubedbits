package loader

import (
	"testing"

	"github.com/renenieuw/cubedbits/assets"
)

// // TestHelloName calls greetings.Hello with a name, checking
// // for a valid return value.
// func TestHelloName(t *testing.T) {
//     name := "Gladys"
//     want := "Gladyss"
//     if want != name  {
//         t.Errorf(`fouts %s`, want)
//     }
// }

func TestLoadSpritesheet(t *testing.T) {
	var a = LoadSpriteSheetsFromJson(assets.CubedBitsSpritesheets, "cubedBits", "CubedBits" );
	log.Printf("%d", len(a["CubedBits"].Sprites))
}
