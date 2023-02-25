package chess

// A Position describes an arbitrary chess position.
type Position struct {
	Board          Board
	CastleRights   CastleRights
	EnPassantRight EnPassantRight
	ActiveColor    Color
}

// NewPosition returns a new starting position.
func NewPosition() Position {
	return Position{
		Board:        NewBoard(),
		CastleRights: WhiteKingSide | WhiteQueenSide | BlackKingSide | BlackQueenSide,
	}
}

// Move applies a move to the position. The move must be legal.
//
// Move returns true if it resets the FIDE 50-move and 75-move rule
// counters. See FIDE Laws of Chess, Articles 9.3 and 9.6.2.
func (p *Position) Move(m Move) (reset bool) {
	// Determine whether move counters must be reset.
	fromPiece, _ := p.Board.Get(m.From)
	_, toPieceExists := p.Board.Get(m.To)
	if fromPiece.Type == Pawn || toPieceExists {
		reset = true
	}

	// Invert the active color.
	p.ActiveColor = !p.ActiveColor

	// Update the en passant right.
	p.EnPassantRight.Valid = false
	if fromPiece.Type == Pawn {
		fromRank, toRank := m.From.Rank(), m.To.Rank()
		switch {
		case fromRank == Rank2 && toRank == Rank4:
			p.EnPassantRight = EnPassantRight{Square: m.From.Above(), Valid: true}
		case fromRank == Rank7 && toRank == Rank5:
			p.EnPassantRight = EnPassantRight{Square: m.From.Below(), Valid: true}
		}
	}

	// Clear the "from" square.
	p.Board.Clear(m.From)

	// Handle castling logic.
	var castleRightsToClear CastleRights

	switch fromPiece.Type {
	case Rook:
		switch {
		case fromPiece.Color == White && m.From == A1:
			castleRightsToClear = WhiteQueenSide
		case fromPiece.Color == White && m.From == H1:
			castleRightsToClear = WhiteKingSide
		case fromPiece.Color == Black && m.From == A8:
			castleRightsToClear = BlackQueenSide
		case fromPiece.Color == Black && m.From == H8:
			castleRightsToClear = BlackKingSide
		}
	case King:
		if fromPiece.Color == White {
			castleRightsToClear = WhiteKingSide | WhiteQueenSide
		} else {
			castleRightsToClear = BlackKingSide | BlackQueenSide
		}
		// If this is a castle move, adjust the castling rook.
		switch {
		case m.From == E1 && m.To == G1:
			p.Board.Move(Piece{White, Rook}, H1, F1)
		case m.From == E1 && m.To == C1:
			p.Board.Move(Piece{White, Rook}, A1, D1)
		case m.From == E8 && m.To == G8:
			p.Board.Move(Piece{Black, Rook}, H8, F8)
		case m.From == E8 && m.To == C8:
			p.Board.Move(Piece{Black, Rook}, A8, D8)
		}
	}

	p.CastleRights.Clear(castleRightsToClear)

	// Handle en passant capture.
	if fromPiece.Type == Pawn && p.EnPassantRight.Valid && p.EnPassantRight.Square == m.To {
		if fromPiece.Color == White {
			p.Board.Clear(m.To.Below())
		} else {
			p.Board.Clear(m.To.Above())
		}
	}

	// Place the from piece on the to square.
	if m.IsPromotion {
		fromPiece.Type = m.Promotion
	}
	p.Board.Set(fromPiece, m.To)

	return
}

// A Board describes the placement of pieces.
type Board struct {
	Types  [6]Bitboard // Piece occupancies indexed by piece type.
	Colors [2]Bitboard // Piece occupancies indexed by piece color.
}

// NewBoard returns a new board with all pieces in their starting positions.
func NewBoard() Board {
	return Board{
		Types: [6]Bitboard{
			0x00FF00000000FF00, // Pawns
			0x4200000000000042, // Knights
			0x2400000000000024, // Bishops
			0x8100000000000081, // Rooks
			0x0800000000000008, // Queens
			0x1000000000000010, // Kings
		},
		Colors: [2]Bitboard{
			0x000000000000FFFF, // White
			0xFFFF000000000000, // Black
		},
	}
}

// Get returns the piece at the given square, if any.
func (b *Board) Get(s Square) (Piece, bool) {
	for t := Pawn; t <= King; t++ {
		if b.Types[t].Get(s) {
			isBlack := b.Colors[Black.Int()].Get(s)
			return Piece{Color(isBlack), t}, true
		}
	}
	return Piece{}, false
}

// Set sets a piece on a square.
// Any piece previously occupying the square is removed.
func (b *Board) Set(p Piece, s Square) {
	b.Clear(s)
	b.Types[p.Type].Set(s)
	b.Colors[p.Color.Int()].Set(s)
}

// Move moves a piece between squares.
// Any piece previously occupying the destination square is removed.
func (b *Board) Move(p Piece, from, to Square) {
	b.Clear(to)
	b.Types[p.Type].Clear(from).Set(to)
	b.Colors[p.Color.Int()].Clear(from).Set(to)
}

// Clear clears the piece on a square, if any.
func (b *Board) Clear(s Square) {
	for i := range b.Types {
		b.Types[i].Clear(s)
	}
	for i := range b.Colors {
		b.Colors[i].Clear(s)
	}
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

// An Outcome represents the outcome of a game.
type Outcome int

const (
	Undecided Outcome = iota
	BlackWon
	Draw
	WhiteWon
)
