package chess

// Bitboard is an integer where each bit represents one square.
// From LSB to MSB, the bits represent A1, B1, ..., H1, A2, ..., H8.
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
