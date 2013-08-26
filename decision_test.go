package battleship

import "testing"

func TestNewDecision(t *testing.T) {
	hitGrid := GameGrid{}

	NewDecision(&hitGrid)
}

func TestMakeDecision(t *testing.T) {
	hitGrid := GameGrid{}

	// Tell the hit grid that all spaces have missed
	hitGrid.Iterate(func(y, x uint8, _ byte) {
		hitGrid[y][x] = MissSpace
	})

	decision := NewDecision(&hitGrid)
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
	hitGrid.Iterate(func(y, x uint8, _ byte) {
		hitGrid[y][x] = AvailableSpace
	})
	hitGrid[5][5] = HitSpace
	point, _ = decision.Make()
	if !(point[0] == 5 && point[1] == 4) && !(point[0] == 5 && point[1] == 6) &&
		!(point[0] == 4 && point[1] == 5) && !(point[0] == 6 && point[1] == 5) {
		t.Error("Decision didn't select the best available space")
	}

	// Show a preference to spaces that follow same axis as previous
	// hits
	hitGrid[5][6] = HitSpace
	point, _ = decision.Make()
	if !(point[0] == 5 && point[1] == 7) && !(point[0] == 5 && point[1] == 4) {
		t.Error("Decision didn't favour hit axis")
	}
}

func BenchmarkMakeDecision(b *testing.B) {
	hitGrid := GameGrid{}

	hitGrid.Iterate(func(y, x uint8, _ byte) {
		hitGrid[y][x] = AvailableSpace
	})

	hitGrid[4][4] = HitSpace
	hitGrid[5][4] = HitSpace
	hitGrid[4][5] = HitSpace
	hitGrid[5][5] = HitSpace

	d := NewDecision(&hitGrid)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		d.Make()
	}
}
