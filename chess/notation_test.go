package chess

import "testing"

func TestGame_FEN(t *testing.T) {
	g := NewGame()
	got := g.FEN()
	if got != StartingFEN {
		t.Errorf("got %q, want %q", got, StartingFEN)
	}
}
