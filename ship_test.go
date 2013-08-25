package battleship

import "testing"

func TestNewShip(t *testing.T) {
	var ship *Ship = NewShip(4)

	if ship.Length() != 4 {
		t.Error("Ship length is incorrect")
	}
}
