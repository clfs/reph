package chess

// A Position weakly describes the state of a chess game.
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
	p.Board.ClearPiece(fromPiece, m.From)

	// Handle castling logic.
	var castleRightsUsed CastleRights

	switch fromPiece.Type {
	case Rook:
		// Moving a rook from its starting square relinquishes one castling right.
		switch {
		case fromPiece.Color == White && m.From == A1:
			castleRightsUsed = WhiteQueenSide
		case fromPiece.Color == White && m.From == H1:
			castleRightsUsed = WhiteKingSide
		case fromPiece.Color == Black && m.From == A8:
			castleRightsUsed = BlackQueenSide
		case fromPiece.Color == Black && m.From == H8:
			castleRightsUsed = BlackKingSide
		}
	case King:
		// King moves always relinquish both castling rights.
		if fromPiece.Color == White {
			castleRightsUsed = WhiteKingSide | WhiteQueenSide
		} else {
			castleRightsUsed = BlackKingSide | BlackQueenSide
		}
		// If this is a castle move, adjust the castling rook.
		switch {
		case m.From == E1 && m.To == G1:
			p.Board.MovePieceToEmptySquare(Piece{White, Rook}, H1, F1)
		case m.From == E1 && m.To == C1:
			p.Board.MovePieceToEmptySquare(Piece{White, Rook}, A1, D1)
		case m.From == E8 && m.To == G8:
			p.Board.MovePieceToEmptySquare(Piece{Black, Rook}, H8, F8)
		case m.From == E8 && m.To == C8:
			p.Board.MovePieceToEmptySquare(Piece{Black, Rook}, A8, D8)
		}
	}

	p.CastleRights.Clear(castleRightsUsed)

	// Handle en passant capture.
	if fromPiece.Type == Pawn && p.EnPassantRight.Valid && p.EnPassantRight.Square == m.To {
		if fromPiece.Color == White {
			p.Board.ClearPiece(Piece{Black, Pawn}, m.To.Below())
		} else {
			p.Board.ClearPiece(Piece{White, Pawn}, m.To.Above())
		}
	}

	// Place the from piece on the to square.
	if m.IsPromotion {
		fromPiece.Type = m.Promotion
	}
	p.Board.SetPiece(fromPiece, m.To)

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

// SetPieceOnEmptySquare sets a piece on the given empty square.
// When available, this is more performant than SetPiece.
func (b *Board) SetPieceOnEmptySquare(p Piece, s Square) {
	b.Types[p.Type].Set(s)
	b.Colors[p.Color.Int()].Set(s)
}

// SetPiece sets a piece on the given square.
// Any piece on the destination square is removed.
func (b *Board) SetPiece(p Piece, s Square) {
	toPiece, ok := b.Get(s)
	if ok {
		b.ClearPiece(toPiece, s)
	}
	b.SetPieceOnEmptySquare(p, s)
}

// MovePieceToEmptySquare moves a piece to an empty square.
// When available, this is more performant than MovePiece.
func (b *Board) MovePieceToEmptySquare(p Piece, from, to Square) {
	b.Types[p.Type].Clear(from).Set(to)
	b.Colors[p.Color.Int()].Clear(from).Set(to)
}

// MovePiece moves a piece between squares.
// Any piece on the destination square is removed.
func (b *Board) MovePiece(p Piece, from, to Square) {
	toPiece, ok := b.Get(to)
	if ok {
		b.ClearPiece(toPiece, to)
	}
	b.MovePieceToEmptySquare(p, from, to)
}

// ClearPiece clears a piece from the given square, if any.
func (b *Board) ClearPiece(p Piece, s Square) {
	b.Types[p.Type].Clear(s)
	b.Colors[p.Color.Int()].Clear(s)
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
