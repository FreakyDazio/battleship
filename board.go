package battleship

import (
	"math/rand"
	"time"
)

const (
	BoardWaterSpace byte = 0
	BoardShipSpace  byte = 1
)

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
		availableSpaces := b.elements.LocateSpace(int(ship.Length()))
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
