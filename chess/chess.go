// Package chess implements basic chess functionality.
package chess

// A Square represents a square on the chess board.
// Squares are ordered left to right, then bottom to top.
type Square uint8

const (
	A1 Square = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

// NewSquare returns a square at the given file and rank.
func NewSquare(f File, r Rank) Square {
	return Square(r)*8 + Square(f)
}

// File returns the file the square is on.
func (s Square) File() File {
	return File(s % 8)
}

// Rank returns the rank the square is on.
func (s Square) Rank() Rank {
	return Rank(s / 8)
}

// Bitboard returns a bitboard representing the square.
func (s Square) Bitboard() Bitboard {
	return Bitboard(1 << s)
}

// Above returns the square above s, wrapping around.
func (s Square) Above() Square {
	return (s + 8) % 64
}

// Below returns the square below s, wrapping around.
func (s Square) Below() Square {
	return (s - 8) % 64
}

// Next returns the square after s, wrapping around.
func (s Square) Next() Square {
	return (s + 1) % 64
}

// NextN returns the square n squares after s, wrapping around.
func (s Square) NextN(n int) Square {
	return (s + Square(n)) % 64
}

// Prev returns the square before s, wrapping around.
func (s Square) Prev() Square {
	return (s - 1) % 64
}

// PrevN returns the square n squares before s, wrapping around.
func (s Square) PrevN(n int) Square {
	return (s - Square(n)) % 64
}

// A File is a column on the chess board.
type File uint8

const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

// A Rank is a row on the chess board.
type Rank uint8

const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

// Bitboard is an integer where each bit represents one square. From LSB to MSB,
// the bits represent squares from left to right, then bottom to top.
type Bitboard uint64

// Get returns true if the square is set.
func (b *Bitboard) Get(s Square) bool {
	return *b&s.Bitboard() != 0
}

// Set sets s and returns b.
func (b *Bitboard) Set(s Square) *Bitboard {
	*b |= s.Bitboard()
	return b
}

// Clear clears s and returns b.
func (b *Bitboard) Clear(s Square) *Bitboard {
	*b &^= s.Bitboard()
	return b
}

// Toggle toggles s and returns b.
func (b *Bitboard) Toggle(s Square) *Bitboard {
	*b ^= s.Bitboard()
	return b
}

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

// Move is a chess move.
type Move struct {
	// The square the piece starts on. For castling moves, this is the king's
	// starting square.
	From Square

	// The square the piece ends on. For castling moves, this is the king's
	// ending square.
	To Square

	// True if the move is a promotion.
	IsPromotion bool

	// The type of the promotion piece. Only valid if IsPromotion is true.
	Promotion Type
}

// NewMove returns a new move.
func NewMove(from, to Square) Move {
	return Move{From: from, To: to}
}

// NewPromotionMove returns a new promotion move.
func NewPromotionMove(from, to Square, promotion Type) Move {
	return Move{From: from, To: to, IsPromotion: true, Promotion: promotion}
}

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
	if fromPiece.Type() == Pawn || toPieceExists {
		reset = true
	}

	// Invert the active color.
	p.ActiveColor = p.ActiveColor.Other()

	// Update the en passant right.
	p.EnPassantRight.Valid = false
	if fromPiece.Type() == Pawn {
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

	switch fromPiece.Type() {
	case Rook:
		switch {
		case fromPiece.Color() == White && m.From == A1:
			castleRightsToClear = WhiteQueenSide
		case fromPiece.Color() == White && m.From == H1:
			castleRightsToClear = WhiteKingSide
		case fromPiece.Color() == Black && m.From == A8:
			castleRightsToClear = BlackQueenSide
		case fromPiece.Color() == Black && m.From == H8:
			castleRightsToClear = BlackKingSide
		}
	case King:
		if fromPiece.Color() == White {
			castleRightsToClear = WhiteKingSide | WhiteQueenSide
		} else {
			castleRightsToClear = BlackKingSide | BlackQueenSide
		}
		// If this is a castle move, adjust the castling rook.
		switch {
		case m.From == E1 && m.To == G1:
			p.Board.Move(WhiteRook, H1, F1)
		case m.From == E1 && m.To == C1:
			p.Board.Move(WhiteRook, A1, D1)
		case m.From == E8 && m.To == G8:
			p.Board.Move(BlackRook, H8, F8)
		case m.From == E8 && m.To == C8:
			p.Board.Move(BlackRook, A8, D8)
		}
	}

	p.CastleRights.Clear(castleRightsToClear)

	// Handle en passant capture.
	if fromPiece.Type() == Pawn && p.EnPassantRight.Valid && p.EnPassantRight.Square == m.To {
		if fromPiece.Color() == White {
			p.Board.Clear(m.To.Below())
		} else {
			p.Board.Clear(m.To.Above())
		}
	}

	// Place the from piece on the to square.
	if m.IsPromotion {
		fromPiece = NewPiece(fromPiece.Color(), m.Promotion)
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
			if b.Colors[White].Get(s) {
				return NewPiece(White, t), true
			} else {
				return NewPiece(Black, t), true
			}
		}
	}
	return 0, false
}

// Set sets a piece on a square.
// Any piece previously occupying the square is removed.
func (b *Board) Set(p Piece, s Square) {
	b.Clear(s)
	b.Types[p.Type()].Set(s)
	b.Colors[p.Color()].Set(s)
}

// Move moves a piece between squares.
// Any piece previously occupying the destination square is removed.
func (b *Board) Move(p Piece, from, to Square) {
	b.Clear(to)
	b.Types[p.Type()].Clear(from).Set(to)
	b.Colors[p.Color()].Clear(from).Set(to)
}

// Clear removes a piece from a square.
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

// A Game contains the history of a chess game.
type Game struct {
	Positions      []Position
	Moves          []Move
	HalfMoveClock  int // Half moves since the last capture or pawn advance.
	FullMoveNumber int // Full move number for the next move.
}

// NewGame returns a new game.
func NewGame() *Game {
	return &Game{
		Positions:      []Position{NewPosition()},
		Moves:          []Move{},
		FullMoveNumber: 1,
	}
}

// CurrentPosition returns the current game position.
func (g *Game) CurrentPosition() Position {
	return g.Positions[len(g.Positions)-1]
}

// Move applies a move to the game. The move must be legal.
func (g *Game) Move(m Move) {
	p := g.CurrentPosition()

	if p.ActiveColor == Black {
		g.FullMoveNumber++
	}

	if reset := p.Move(m); reset {
		g.HalfMoveClock = 0
	} else {
		g.HalfMoveClock++
	}

	g.Positions = append(g.Positions, p)
	g.Moves = append(g.Moves, m)
}
