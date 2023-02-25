package chess

import "testing"

func TestGame_FEN(t *testing.T) {
	got := NewGame().FEN()
	want := StartingFEN

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestNewGameFromFEN(t *testing.T) {
	game, err := NewGameFromFEN(StartingFEN)
	if err != nil {
		t.Fatalf("NewGameFromFEN: %v", err)
	}

	got := game.FEN()
	want := StartingFEN

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
