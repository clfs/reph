package chess

import "testing"

func TestPosition_Move(t *testing.T) {
	p := NewPosition()

	reset := p.Move(Move{From: E2, To: E4})
	if !reset {
		t.Error("no counter reset after 1. e2e4")
	}

	e4piece, ok := p.Board.Get(E4)
	if !ok {
		t.Error("no piece at e4")
	}
	if e4piece != WhitePawn {
		t.Errorf("wrong piece at e4: got %v, want %v", e4piece, WhitePawn)
	}

	_, ok = p.Board.Get(E2)
	if ok {
		t.Error("unexpected piece at e2")
	}

	if !p.EnPassantRight.Valid {
		t.Error("en passant right not set")
	}
	if p.EnPassantRight.Square != E3 {
		t.Errorf("wrong en passant square: got %v, want %v", p.EnPassantRight.Square, E3)
	}

	reset = p.Move(Move{From: E7, To: E5})
	if !reset {
		t.Error("no counter reset after 1... e7e5")
	}
}
