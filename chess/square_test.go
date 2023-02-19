package chess

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

// Generate implements testing/quick.Generator for Square.
func (Square) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Square(r.Intn(64)))
}

// Generate implements testing/quick.Generator for File.
func (File) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(File(r.Intn(8)))
}

// Generate implements testing/quick.Generator for Rank.
func (Rank) Generate(r *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Rank(r.Intn(8)))
}

func TestSquare_File(t *testing.T) {
	f := func(f File, r Rank) bool {
		return NewSquare(f, r).File() == f
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSquare_Rank(t *testing.T) {
	f := func(f File, r Rank) bool {
		return NewSquare(f, r).Rank() == r
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
