package battleship

import "testing"

func TestNewDecision(t *testing.T) {
	var shipGrid GameGrid
	var hitGrid GameGrid

	NewDecision(&shipGrid, &hitGrid)
}

func TestMakeDecision(t *testing.T) {
	var shipGrid GameGrid
	var hitGrid GameGrid

	// Tell the hit grid that all spaces have missed
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			hitGrid[y][x] = MissSpace
		}
	}

	decision := NewDecision(&shipGrid, &hitGrid)
	_, err := decision.Make()
	if err == nil {
		t.Error("Unavailable decision not discovered")
	}

	// Make a single space available
	hitGrid[9][9] = AvailableSpace
	point, _ := decision.Make()
	if point[0] != 9 && point[1] != 9 {
		t.Error("Decision didn't select final available space")
	}
}
