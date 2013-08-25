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

func (d Decision) Make() ([2]uint8, error) {
	availableSpaces := make([][2]uint8, 0)
	var result [2]uint8

	for y := uint8(0); y < uint8(10); y++ {
		for x := uint8(0); x < uint8(10); x++ {
			if d.hitGrid[y][x] == AvailableSpace {
				availableSpaces = append(availableSpaces, [2]uint8{y, x})
			}
		}
	}

	if len(availableSpaces) == 0 {
		return result, errors.New("No available moves")
	}

	rand.Seed(time.Now().UnixNano())

	return availableSpaces[rand.Intn(len(availableSpaces))], nil

}
