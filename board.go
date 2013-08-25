package battleship

import (
	"math/rand"
	"time"
)

const (
	BoardWaterSpace byte = 0
	BoardShipSpace  byte = 1
)

type GameGrid [10][10]byte

type Board struct {
	player   *Player
	elements GameGrid
}

func NewBoard(p *Player) *Board {
	b := Board{player: p}
	ships := [5]*Ship{
		NewShip(2),
		NewShip(3),
		NewShip(3),
		NewShip(4),
		NewShip(5),
	}
	rand.Seed(time.Now().UnixNano())

	for _, ship := range ships {
		availableSpaces := b.elements.locateSpace(int(ship.Length()))
		options := len(availableSpaces)
		if options > 0 {
			selected := availableSpaces[rand.Intn(options)]
			for _, coord := range selected {
				b.elements[coord[0]][coord[1]] = BoardShipSpace
			}
		}
	}

	return &b
}

func (b Board) Player() *Player {
	return b.player
}

func (grid GameGrid) locateSpace(maxLength int) [][][2]uint8 {
	result := make([][][2]uint8, 0)
	for y := 0; y <= (9 - (maxLength - 1)); y++ {
		for x := 0; x <= (9 - (maxLength - 1)); x++ {
			verticalComb := make([][2]uint8, 0, maxLength)
			for vi := 0; vi < maxLength; vi++ {
				if !grid.spaceAvailable(y+vi, x) {
					break
				}
				verticalComb = append(verticalComb, [2]uint8{uint8(y + vi), uint8(x)})
			}
			if len(verticalComb) == maxLength {
				result = append(result, verticalComb)
			}

			horizontalComb := make([][2]uint8, 0, maxLength)
			for hi := 0; hi < maxLength; hi++ {
				if !grid.spaceAvailable(y, x+hi) {
					break
				}
				horizontalComb = append(horizontalComb, [2]uint8{uint8(y), uint8(x + hi)})
			}
			if len(horizontalComb) == maxLength {
				result = append(result, horizontalComb)
			}
		}
	}
	return result
}

func (grid GameGrid) spaceAvailable(y, x int) bool {
	if y > 9 || x > 9 {
		return false
	}
	return grid[y][x] == BoardWaterSpace
}
