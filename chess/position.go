package chess

// A Position weakly describes the state of a chess game.
type Position struct {
	Board          Board
	CastleRights   CastleRights
	EnPassantRight EnPassantRight
	ActiveColor    Color

	// The number of half moves since the last capture or pawn move.
	// See FIDE [Laws of Chess], Article 9.6.2.
	//
	// [Laws of Chess]: https://handbook.fide.com/chapter/E012023
	Rule75 uint8
}

// A Board describes the placement of pieces.
type Board struct {
	Types [6]Bitboard // Piece occupancies indexed by piece type.
	White Bitboard    // Piece occupancy for the white pieces.
	Black Bitboard    // Piece occupancy for the black pieces.
}

// At returns the piece at the given square, if any.
func (b Board) At(s Square) (Piece, bool) {
	for t := Pawn; t <= King; t++ {
		if b.Types[t].Get(s) {
			c := Color(b.White.Get(s))
			return Piece{c, t}, true
		}
	}
	return Piece{}, false
}

// CastleRights represents the castling rights of both players.
type CastleRights uint8

const (
	WhiteKingSide CastleRights = 1 << iota
	WhiteQueenSide
	BlackKingSide
	BlackQueenSide
)

// Get returns true if any of the given castle rights are set.
func (c *CastleRights) Get(r CastleRights) bool {
	return *c&r != 0
}

// Set sets the given castle rights.
func (c *CastleRights) Set(r CastleRights) {
	*c |= r
}

// Clear clears the given castle rights.
func (c *CastleRights) Clear(r CastleRights) {
	*c &^= r
}

// EnPassantRight represents the right to en passant.
type EnPassantRight struct {
	Square Square
	Valid  bool // Valid is true if an en passant target square exists.
}
