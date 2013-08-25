package battleship

import "testing"

var player *Player = NewPlayer("example", "example@example.com")

func TestNewBoard(t *testing.T) {
	var b *Board

	// Because it is a randomly generated board I want to try
	// and catch any funky problems with board generation. This
	// stemmed from it occasionaly failing to place the last ship.
	for i := 0; i < 10000; i++ {
		b = NewBoard(player)

		if b.Player() != player {
			t.Error("Player is incorrect")
		}

		if b.elements[0][0] > BoardWaterSpace|BoardShipSpace {
			t.Errorf("Board space x:0 y:0 is not available. Got: %d", b.elements[0][0])
		}

		if b.elements[9][9] > BoardWaterSpace|BoardShipSpace {
			t.Errorf("Board space x:9 y:9 is not available. Got: %d", b.elements[9][9])
		}

		shipTiles := 0
		for _, col := range b.elements {
			for _, val := range col {
				if val == BoardShipSpace {
					shipTiles += 1
				}
			}
		}

		if shipTiles != (2 + 3 + 3 + 4 + 5) {
			t.Log(b.elements)
			t.Errorf("Not enough ships spaces on board. Only found %d", shipTiles)
		}
	}
}

func BenchmarkNewBoard(b *testing.B) {
	p := NewPlayer("test", "test@example.com")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewBoard(p)
	}
}

// The locateSpace method is a bit beefy. It has some very minor
// optimisations aroud discovery. It's good to see how it varies
// between sizes
func BenchmarkLocationDiscoverySmall(b *testing.B) {
	var grid GameGrid
	for i := 0; i < b.N; i++ {
		grid.locateSpace(2)
	}
}

func BenchmarkLocationDiscoveryMedium(b *testing.B) {
	var grid GameGrid
	for i := 0; i < b.N; i++ {
		grid.locateSpace(5)
	}
}

func BenchmarkLocationDiscoveryLarge(b *testing.B) {
	var grid GameGrid
	for i := 0; i < b.N; i++ {
		grid.locateSpace(8)
	}
}

func BenchmarkLocationDiscoveryTooLarge(b *testing.B) {
	var grid GameGrid
	for i := 0; i < b.N; i++ {
		grid.locateSpace(11)
	}
}
