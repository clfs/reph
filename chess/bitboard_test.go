package chess

import (
	"testing"
	"testing/quick"
)

func TestBitboard(t *testing.T) {
	setThenGet := func(b Bitboard, s Square) bool {
		return b.Set(s).Get(s)
	}

	if err := quick.Check(setThenGet, nil); err != nil {
		t.Error(err)
	}

	clearThenGet := func(b Bitboard, s Square) bool {
		return !b.Clear(s).Get(s)
	}

	if err := quick.Check(clearThenGet, nil); err != nil {
		t.Error(err)
	}

	toggleNotEq := func(b Bitboard, s Square) bool {
		return b.Get(s) != b.Toggle(s).Get(s)
	}

	if err := quick.Check(toggleNotEq, nil); err != nil {
		t.Error(err)
	}
}
