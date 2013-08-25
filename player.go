package battleship

type Player struct {
	name  string
	email string
}

func NewPlayer(name, email string) *Player {
	return &Player{name: name, email: email}
}

func (p Player) Name() string {
	return p.name
}

func (p Player) Email() string {
	return p.email
}
