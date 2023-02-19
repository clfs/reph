package chess

// Piece is a chess piece.
type Piece struct {
	Color Color
	Type  Type
}

// Color is either White or Black.
type Color bool

const (
	White Color = false
	Black Color = true
)

// Int returns 0 for White and 1 for Black.
func (c Color) Int() int {
	if c == White {
		return 0
	}
	return 1
}

// Type represents a piece type.
type Type uint8

const (
	Pawn Type = iota
	Knight
	Bishop
	Rook
	Queen
	King
)
