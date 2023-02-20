package chess

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

// Generate implements testing/quick.Generator for CastleRights.
func (CastleRights) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(CastleRights(r.Intn(16)))
}

func TestCastleRights(t *testing.T) {
	setThenGet := func(a, b CastleRights) bool {
		a.Set(b)
		return a.Get(b) || b == 0
	}

	if err := quick.Check(setThenGet, nil); err != nil {
		t.Error(err)
	}

	clearThenGet := func(a, b CastleRights) bool {
		a.Clear(b)
		return !a.Get(b)
	}

	if err := quick.Check(clearThenGet, nil); err != nil {
		t.Error(err)
	}
}

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
