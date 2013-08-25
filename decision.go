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
	var scoredGrid GameGrid
	var shipSpacesDiscovered uint8
	var result [2]uint8

	for y := uint8(0); y < uint8(10); y++ {
		for x := uint8(0); x < uint8(10); x++ {
			if d.hitGrid[y][x] == AvailableSpace {
				scoredGrid[y][x]++ // Increment the score for this cell
			} else if d.hitGrid[y][x] == HitSpace {
				shipSpacesDiscovered++
				scoredGrid[y][x] = 0 // Ensure this cell is ignored
				if y < 9 {
					if d.hitGrid[y+1][x] == AvailableSpace {
						scoredGrid[y+1][x]++
					}
					if y > 0 && d.hitGrid[y-1][x] == HitSpace {
						scoredGrid[y+1][x] += 2
					}
				}
				if y > 0 {
					if d.hitGrid[y-1][x] == AvailableSpace {
						scoredGrid[y-1][x]++
					}
					if y < 9 && d.hitGrid[y+1][x] == HitSpace {
						scoredGrid[y-1][x] += 2
					}
				}
				if x < 9 {
					if d.hitGrid[y][x+1] == AvailableSpace {
						scoredGrid[y][x+1]++
					}
					if x > 0 && d.hitGrid[y][x-1] == HitSpace {
						scoredGrid[y][x+1] += 2
					}
				}
				if x > 0 {
					if d.hitGrid[y][x-1] == AvailableSpace {
						scoredGrid[y][x-1]++
					}
					if x < 9 && d.hitGrid[y][x+1] == HitSpace {
						scoredGrid[y][x-1] += 2
					}
				}
			} else {
				scoredGrid[y][x] = 0 // Ensure this cell is ignored
			}
		}
	}

	ratedMoves := make([][2]uint8, 0)
	var highestScore uint8 = 0
	for y := uint8(0); y < uint8(10); y++ {
		for x := uint8(0); x < uint8(10); x++ {
			if scoredGrid[y][x] > highestScore {
				highestScore = scoredGrid[y][x]
				ratedMoves = make([][2]uint8, 0)                // Reset rated moves
				ratedMoves = append(ratedMoves, [2]uint8{y, x}) // Add to rated moves
			} else if scoredGrid[y][x] == highestScore {
				ratedMoves = append(ratedMoves, [2]uint8{y, x}) // Add to rated moves
			}
		}
	}

	if len(ratedMoves) == 0 || highestScore == 0 {
		return result, errors.New("No available moves")
	} else {
		rand.Seed(time.Now().UnixNano())
		result = ratedMoves[rand.Intn(len(ratedMoves))]
	}

	return result, nil
}
