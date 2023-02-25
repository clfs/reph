package chess

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
