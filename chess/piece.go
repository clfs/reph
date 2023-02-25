package chess

// Piece is a chess piece.
type Piece uint8

// NewPiece returns a new piece of the given color and type.
func NewPiece(c Color, t Type) Piece {
	return Piece(c)<<3 | Piece(t)
}

const (
	WhitePawn Piece = iota
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
	_
	_
	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

// Color returns the color of the piece.
func (p Piece) Color() Color {
	return Color(p >> 3)
}

// Type returns the type of the piece.
func (p Piece) Type() Type {
	return Type(p & 0b0111)
}

// Color is either White or Black.
type Color uint8

const (
	White Color = iota
	Black
)

// Other returns the other color.
func (c Color) Other() Color {
	return c ^ 1
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
