package chess

// A Position weakly describes the state of a chess game.
type Position struct {
	Board          Board
	CastleRights   CastleRights
	EnPassantRight EnPassantRight
	ActiveColor    Color

	// The number of plies since the last capture or pawn movement.
	// See FIDE [Laws of Chess], Article 9.6.2.
	//
	// [Laws of Chess]: https://handbook.fide.com/chapter/E012023
	Rule75 uint8
}

// Move applies a move to the position. The move must be legal.
func (p *Position) Move(m Move) {
	// Invert the active color.
	p.ActiveColor = !p.ActiveColor

	// Update the 75-move-rule counter.
	fromPiece, _ := p.Board.Get(m.From)
	toPiece, isCapture := p.Board.Get(m.To)
	if fromPiece.Type == Pawn || isCapture {
		p.Rule75 = 0
	} else {
		p.Rule75++
	}

	// Update the en passant right.
	if fromPiece.Type == Pawn {
		fromRank, toRank := m.From.Rank(), m.To.Rank()
		switch {
		case fromRank == Rank2 && toRank == Rank4:
			p.EnPassantRight = EnPassantRight{Square: m.From.Above(), Valid: true}
		case fromRank == Rank7 && toRank == Rank5:
			p.EnPassantRight = EnPassantRight{Square: m.From.Below(), Valid: true}
		default:
			p.EnPassantRight.Valid = false
		}
	} else {
		p.EnPassantRight.Valid = false
	}

	// Update the "from" square.
	p.Board.ClearPiece(fromPiece, m.From)

	// Update the "to" square.
	toPiece = fromPiece
	if m.IsPromotion {
		toPiece.Type = m.Promotion
	}
	p.Board.SetPiece(toPiece, m.To)

	// Handle castling.
	castleRightsUsed := m.CastleRightsUsed()
	p.CastleRights.Clear(castleRightsUsed)
	switch castleRightsUsed {
	case WhiteKingSide:
		p.Board.MovePieceToEmptySquare(Piece{White, Rook}, H1, F1)
	case WhiteQueenSide:
		p.Board.MovePieceToEmptySquare(Piece{White, Rook}, A1, D1)
	case BlackKingSide:
		p.Board.MovePieceToEmptySquare(Piece{Black, Rook}, H8, F8)
	case BlackQueenSide:
		p.Board.MovePieceToEmptySquare(Piece{Black, Rook}, A8, D8)
	}
}

// A Board describes the placement of pieces.
type Board struct {
	Types  [6]Bitboard // Piece occupancies indexed by piece type.
	Colors [2]Bitboard // Piece occupancies indexed by piece color.
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
