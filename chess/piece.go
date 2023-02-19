package chess

// Piece is a chess piece.
type Piece struct {
	Color Color
	Role  Role
}

// Color is either White or Black.
type Color bool

const (
	White Color = false
	Black Color = true
)

// Role represents a piece's role.
type Role uint8

const (
	Pawn Role = iota
	Knight
	Bishop
	Rook
	Queen
	King
)
