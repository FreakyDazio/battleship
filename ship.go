package battleship

type Ship struct {
	length uint8
}

func NewShip(length uint8) *Ship {
	return &Ship{length: length}
}

func (s Ship) Length() uint8 {
	return s.length
}
