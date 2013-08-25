package battleship

import "testing"

func TestLocateSpace(t *testing.T) {
	var grid GameGrid
	options := grid.LocateSpace(10)
	if length := len(options); length != 20 {
		t.Errorf("Failed to find all available spaces. Found %d", length)
	}

	options = grid.LocateSpace(5)
	if length := len(options); length != 120 {
		t.Errorf("Failed to find all available spaces. Found %d", length)
	}

	for i, _ := range grid[0] {
		grid[0][i] = BoardShipSpace
	}
	options = grid.LocateSpace(10)
	if length := len(options); length != 9 {
		t.Errorf("Failed to find all available spaces. Found %d", length)
	}
}

func TestSpaceAvailable(t *testing.T) {
	var grid GameGrid
	if !grid.SpaceAvailable(0, 0) {
		t.Error("Grid reported free space as occupied")
	}

	grid[0][0] = BoardShipSpace
	if grid.SpaceAvailable(0, 0) {
		t.Error("Grid reported occupied space as free")
	}
}

func BenchmarkLocateSpaceSmall(b *testing.B) {
	var grid GameGrid
	for i := 0; i < b.N; i++ {
		grid.LocateSpace(2)
	}
}

func BenchmarkLocateSpaceMedium(b *testing.B) {
	var grid GameGrid
	for i := 0; i < b.N; i++ {
		grid.LocateSpace(5)
	}
}

func BenchmarkLocateSpaceLarge(b *testing.B) {
	var grid GameGrid
	for i := 0; i < b.N; i++ {
		grid.LocateSpace(8)
	}
}

func BenchmarkLocateSpaceTooLarge(b *testing.B) {
	var grid GameGrid
	for i := 0; i < b.N; i++ {
		grid.LocateSpace(11)
	}
}
