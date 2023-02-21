package chess

// A Square represents a square on the chess board.
// Squares are ordered left-to-right, then bottom-to-top.
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
