package chess

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
