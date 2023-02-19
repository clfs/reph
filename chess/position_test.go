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
