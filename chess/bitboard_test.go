package chess

import (
	"testing"
	"testing/quick"
)

func TestBitboard(t *testing.T) {
	f := func(b Bitboard, s Square) bool {
		if !b.Set(s).Get(s) {
			return false
		}
		if b.Clear(s).Get(s) {
			return false
		}
		if b.Get(s) == b.Toggle(s).Get(s) {
			return false
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
