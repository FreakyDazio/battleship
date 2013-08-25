package battleship

import "testing"

var player *Player = NewPlayer("example", "example@example.com")

func TestNewBoard(t *testing.T) {
	b := NewBoard(player)

	if b.Player() != player {
		t.Error("Player is incorrect")
	}

	shipTiles := 0
	b.elements.Iterate(func(y, x uint8, val *byte) {
		if *val == BoardShipSpace {
			shipTiles++
		}
	})

	if shipTiles != (2 + 3 + 3 + 4 + 5) {
		t.Log(b.elements)
		t.Errorf("Not enough ships spaces on board. Only found %d", shipTiles)
	}
}

func BenchmarkNewBoard(b *testing.B) {
	p := NewPlayer("test", "test@example.com")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewBoard(p)
	}
}
