package battleship

type GameGrid [10][10]byte

func (grid GameGrid) Iterate(exec func(y, x uint8, val byte)) {
	for y := uint8(0); y < uint8(10); y++ {
		for x := uint8(0); x < uint8(10); x++ {
			exec(y, x, grid[y][x])
		}
	}
}

func (grid GameGrid) LocateSpace(maxLength uint8) [][][2]uint8 {
	result := make([][][2]uint8, 0)

	grid.Iterate(func(y, x uint8, _ byte) {
		verticalComb := make([][2]uint8, 0, maxLength)
		for vi := uint8(0); vi < maxLength; vi++ {
			if !grid.SpaceAvailable(y+vi, x) {
				break
			}
			verticalComb = append(verticalComb, [2]uint8{y + vi, x})
		}
		if uint8(len(verticalComb)) == maxLength {
			result = append(result, verticalComb)
		}

		horizontalComb := make([][2]uint8, 0, maxLength)
		for hi := uint8(0); hi < maxLength; hi++ {
			if !grid.SpaceAvailable(y, x+hi) {
				break
			}
			horizontalComb = append(horizontalComb, [2]uint8{y, x + hi})
		}
		if uint8(len(horizontalComb)) == maxLength {
			result = append(result, horizontalComb)
		}
	})

	return result
}

func (grid GameGrid) SpaceAvailable(y, x uint8) bool {
	if y > uint8(9) || x > uint8(9) {
		return false
	}
	return grid[y][x] == BoardWaterSpace
}
