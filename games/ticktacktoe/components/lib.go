package components

type Board struct {
	// Texture image
	Color int
}

type Tile struct {
	// Texture image
	X     int
	Y     int
	State int
}

type Line struct {
	// Texture image
	X1 int
	Y1 int
	X2 int
	Y2 int
}
