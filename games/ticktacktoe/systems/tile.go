package systems

import (
	"log"
	"time"

	"github.com/mlange-42/ark/ecs"
	c "remapit.visualstudio.com/cubedbits/cubedbitsengine/components"
	tc "remapit.visualstudio.com/cubedbits/cubedbitsengine/games/ticktacktoe/components"
)

type TileSystem struct {
	Turn        int
	WonAt       time.Time
	WinningLine [3][2]int
}

func (ts *TileSystem) Update(world *ecs.World) {
	filter := ecs.NewFilter4[c.SpriteRender, tc.Tile, c.Transform, c.MouseReactive](world)
	query := filter.Query()
	for query.Next() {
		spriteRender, tile, transform, mouseReactive := query.Get()
		// spriteRender.SpriteNumber = tile.State
		transform.Translation.X = float64(180 + (tile.X * 140))
		transform.Translation.Y = float64((100 + (tile.Y * 140)))
		transform.Rotation = 0
		if mouseReactive.Hovered && mouseReactive.JustClicked && tile.State == 0 && ts.WonAt.IsZero() {
			tile.State = (ts.Turn % 2) + 1
			ts.Turn = ts.Turn + 1
			spriteRender.SpriteNumber = tile.State
		}
	}

	winner, line := ts.CheckWin(world)
	if winner > 0 {
		if ts.WonAt.IsZero() {
			ts.WonAt = time.Now()
			ts.WinningLine = line
		}
	}

	if !ts.WonAt.IsZero() {
		ts.Blink(world)
		log.Printf("%s", ts.WonAt.String())
		if time.Since(ts.WonAt).Seconds() > 2 {
			ts.WonAt = time.Time{}
			ts.Turn = 0 // Reset turn too when resetting board

			query := filter.Query()
			for query.Next() {
				spriteRender, tile, _, _ := query.Get()
				spriteRender.SpriteNumber = 0
				tile.State = 0
			}
		}
	}
}

func CheckWon(world *ecs.World) bool {
	var len int
	filter := ecs.NewFilter4[c.SpriteRender, tc.Tile, c.Transform, c.MouseReactive](world)
	query := filter.Query()
	for query.Next() {
		_, tile, _, _ := query.Get()
		if tile.State != 0 {
			len = len + 1
		}
	}
	return len > 2
}

// CheckWin returns the winning player (1 or 2), and the coordinates of the winning line, or 0 if no one has won yet.
func CheckWin(board [3][3]int) (int, [3][2]int) {
	var empty [3][2]int

	// Check rows
	for y := 0; y < 3; y++ {
		if board[y][0] != 0 && board[y][0] == board[y][1] && board[y][1] == board[y][2] {
			return board[y][0], [3][2]int{{0, y}, {1, y}, {2, y}}
		}
	}

	// Check columns
	for x := 0; x < 3; x++ {
		if board[0][x] != 0 && board[0][x] == board[1][x] && board[1][x] == board[2][x] {
			return board[0][x], [3][2]int{{x, 0}, {x, 1}, {x, 2}}
		}
	}

	// Check main diagonal (top-left to bottom-right)
	if board[0][0] != 0 && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0], [3][2]int{{0, 0}, {1, 1}, {2, 2}}
	}

	// Check anti-diagonal (top-right to bottom-left)
	if board[0][2] != 0 && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2], [3][2]int{{2, 0}, {1, 1}, {0, 2}}
	}

	// No winner yet
	return 0, empty
}

func (ts *TileSystem) CheckWin(world *ecs.World) (int, [3][2]int) {
	var board [3][3]int

	// Set up the filter to grab all tiles
	filter := ecs.NewFilter1[tc.Tile](world)
	query := filter.Query()

	// Map the individual ECS tiles onto a standard 3x3 array
	for query.Next() {
		tile := query.Get()

		// Assuming X and Y are 0, 1, or 2
		if tile.X >= 0 && tile.X < 3 && tile.Y >= 0 && tile.Y < 3 {
			board[tile.Y][tile.X] = tile.State
		}
	}

	// Run the win check
	return CheckWin(board)
}

func (ts *TileSystem) Blink(world *ecs.World) {
	var t = time.Now()
	// Blink by checking the current time in 200ms intervals
	blinkOff := (t.Nanosecond()/200000000)%2 == 0

	filter := ecs.NewFilter4[c.SpriteRender, tc.Tile, c.Transform, c.MouseReactive](world)
	query := filter.Query()
	for query.Next() {
		spriteRender, tile, _, _ := query.Get()

		isWinningTile := false
		for _, coords := range ts.WinningLine {
			if tile.X == coords[0] && tile.Y == coords[1] {
				isWinningTile = true
				break
			}
		}

		if isWinningTile && blinkOff {
			spriteRender.SpriteNumber = 0
		} else {
			spriteRender.SpriteNumber = tile.State
		}
	}
}
