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

// CastleRightsUsed returns the castle rights used by the move, if any.
func (m Move) CastleRightsUsed() CastleRights {
	if m.From == E1 {
		if m.To == G1 {
			return WhiteKingSide
		} else if m.To == C1 {
			return WhiteQueenSide
		}
	} else if m.From == E8 {
		if m.To == G8 {
			return BlackKingSide
		} else if m.To == C8 {
			return BlackQueenSide
		}
	}

	return 0
}

// NewMove returns a new move.
func NewMove(from, to Square) Move {
	return Move{From: from, To: to}
}

// NewPromotionMove returns a new promotion move.
func NewPromotionMove(from, to Square, promotion Type) Move {
	return Move{From: from, To: to, IsPromotion: true, Promotion: promotion}
}
