package battleship

import (
	"errors"
	"math/rand"
	"time"
)

const (
	AvailableSpace byte = 0
	HitSpace       byte = 1
	MissSpace      byte = 2
)

type Decision struct {
	shipGrid *GameGrid
	hitGrid  *GameGrid
}

func NewDecision(shipGrid, hitGrid *GameGrid) *Decision {
	return &Decision{shipGrid: shipGrid, hitGrid: hitGrid}
}

func (grid GameGrid) scoreYAxis(y, x uint8, inc int, scoreGrid *GameGrid, accessible bool) {
	ny := uint8(int(y) + inc)
	negY := uint8(int(y) + (-1 * inc))
	if grid[ny][x] == AvailableSpace {
		scoreGrid[ny][x]++
	}
	if accessible && grid[negY][x] == HitSpace {
		scoreGrid[ny][x] += 2
	}
}

func (grid GameGrid) scoreXAxis(y, x uint8, inc int, scoreGrid *GameGrid, accessible bool) {
	nx := uint8(int(x) + inc)
	negX := uint8(int(x) + (-1 * inc))
	if grid[y][nx] == AvailableSpace {
		scoreGrid[y][nx]++
	}
	if accessible && grid[y][negX] == HitSpace {
		scoreGrid[y][nx] += 2
	}
}

func (d Decision) Make() ([2]uint8, error) {
	var scoredGrid GameGrid
	var shipSpacesDiscovered uint8
	var result [2]uint8

	d.hitGrid.Iterate(func(y, x uint8, val *byte) {
		switch *val {
		case AvailableSpace:
			scoredGrid[y][x]++
		case HitSpace:
			shipSpacesDiscovered++
			scoredGrid[y][x] = 0
			if y < 9 {
				d.hitGrid.scoreYAxis(y, x, 1, &scoredGrid, y > 0)
			}
			if y > 0 {
				d.hitGrid.scoreYAxis(y, x, -1, &scoredGrid, y < 9)
			}
			if x < 9 {
				d.hitGrid.scoreXAxis(y, x, 1, &scoredGrid, x > 0)
			}
			if x > 0 {
				d.hitGrid.scoreXAxis(y, x, -1, &scoredGrid, x < 9)
			}
		default:
			scoredGrid[y][x] = 0
		}
	})

	ratedMoves := make([][2]uint8, 0)
	var highestScore uint8 = 0
	scoredGrid.Iterate(func(y, x uint8, val *byte) {
		if *val > highestScore {
			highestScore = scoredGrid[y][x]
			ratedMoves = make([][2]uint8, 0)                // Reset rated moves
			ratedMoves = append(ratedMoves, [2]uint8{y, x}) // Add to rated moves
		} else if *val == highestScore {
			ratedMoves = append(ratedMoves, [2]uint8{y, x}) // Add to rated moves
		}
	})

	if highestScore == 0 {
		return result, errors.New("No available moves")
	} else {
		rand.Seed(time.Now().UnixNano())
		result = ratedMoves[rand.Intn(len(ratedMoves))]
	}

	return result, nil
}
