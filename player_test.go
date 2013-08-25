package battleship

import "testing"

func TestNewPlayer(t *testing.T) {
	var player *Player = NewPlayer("Example", "test@example.com")

	if player.Name() != "Example" {
		t.Error("Player `Name` is incorrect")
	}

	if player.Email() != "test@example.com" {
		t.Error("Player `Email` is incorrect")
	}
}
