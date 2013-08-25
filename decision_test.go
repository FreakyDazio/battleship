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
	val, err := decision.Make()
	if err == nil {
		t.Errorf("Unavailable decision not discovered. Got: %i", val)
	}

	// Make a single space available
	hitGrid[9][9] = AvailableSpace
	point, _ := decision.Make()
	if point[0] != 9 && point[1] != 9 {
		t.Error("Decision didn't select final available space")
	}

	// Make 2 available spaces but 1 next to a hit
	hitGrid[9][9] = AvailableSpace
	hitGrid[9][8] = AvailableSpace
	hitGrid[9][7] = HitSpace
	point, _ = decision.Make()
	if point[0] != 9 && point[1] != 8 {
		t.Error("Decision didn't select the best available space")
	}

	// Make all spaces available but 1 hit
	// AI should only select spaces around the hit to begin with
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			hitGrid[y][x] = AvailableSpace
		}
	}
	hitGrid[5][5] = HitSpace
	point, _ = decision.Make()
	if !(point[0] == 5 && point[1] == 4) && !(point[0] == 5 && point[1] == 6) &&
		!(point[0] == 4 && point[1] == 5) && !(point[0] == 6 && point[1] == 5) {
		t.Error("Decision didn't select the best available space")
	}
}

func BenchmarkMakeDecision(b *testing.B) {
	var hitGrid GameGrid

	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			hitGrid[y][x] = AvailableSpace
		}
	}

	hitGrid[4][4] = HitSpace
	hitGrid[5][4] = HitSpace
	hitGrid[4][5] = HitSpace
	hitGrid[5][5] = HitSpace

	d := NewDecision(&GameGrid{}, &hitGrid)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		d.Make()
	}
}
