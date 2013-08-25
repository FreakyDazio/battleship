package battleship

type GameGrid [10][10]byte

func (grid GameGrid) LocateSpace(maxLength int) [][][2]uint8 {
	result := make([][][2]uint8, 0)
	for y := 0; y <= 9; y++ {
		for x := 0; x <= 9; x++ {
			verticalComb := make([][2]uint8, 0, maxLength)
			for vi := 0; vi < maxLength; vi++ {
				if !grid.SpaceAvailable(y+vi, x) {
					break
				}
				verticalComb = append(verticalComb, [2]uint8{uint8(y + vi), uint8(x)})
			}
			if len(verticalComb) == maxLength {
				result = append(result, verticalComb)
			}

			horizontalComb := make([][2]uint8, 0, maxLength)
			for hi := 0; hi < maxLength; hi++ {
				if !grid.SpaceAvailable(y, x+hi) {
					break
				}
				horizontalComb = append(horizontalComb, [2]uint8{uint8(y), uint8(x + hi)})
			}
			if len(horizontalComb) == maxLength {
				result = append(result, horizontalComb)
			}
		}
	}
	return result
}

func (grid GameGrid) SpaceAvailable(y, x int) bool {
	if y > 9 || x > 9 {
		return false
	}
	return grid[y][x] == BoardWaterSpace
}
