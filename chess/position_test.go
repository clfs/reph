package chess

import (
	"testing"
)

func TestPosition_Move(t *testing.T) {
	p := NewPosition()

	reset := p.Move(Move{From: E2, To: E4})
	if !reset {
		t.Error("Position.Move(): no counter reset after 1. e2e4")
	}

	e4piece, ok := p.Board.Get(E4)
	if !ok {
		t.Error("Position.Move(): no piece at e4")
	}
	if e4piece.Color != White || e4piece.Type != Pawn {
		t.Errorf(
			"Position.Move(): wrong piece at e4: got %v, want %v",
			e4piece, Piece{White, Pawn},
		)
	}

	_, ok = p.Board.Get(E2)
	if ok {
		t.Error("Position.Move(): unexpected piece at e2")
	}

	if !p.EnPassantRight.Valid {
		t.Error("Position.Move(): en passant right not set")
	}
	if p.EnPassantRight.Square != E3 {
		t.Errorf(
			"Position.Move(): wrong en passant square: got %v, want %v",
			p.EnPassantRight.Square, E3,
		)
	}

	reset = p.Move(Move{From: E7, To: E5})
	if !reset {
		t.Error("Position.Move(): no counter reset after 1... e7e5")
	}
}
