package chess

// A Game contains the history of a chess game.
type Game struct {
	Positions []Position
	Moves     []Move

	// The number of half moves since the last capture or pawn advance.
	HalfMoveClock int

	// The full move number for the next move.
	FullMoveNumber int
}

// NewGame returns a new game.
func NewGame() *Game {
	return &Game{
		Positions:      []Position{NewPosition()},
		Moves:          []Move{},
		FullMoveNumber: 1,
	}
}

// FirstPosition returns the first position of the game.
func (g *Game) FirstPosition() Position {
	return g.Positions[0]
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

	reset := p.Move(m)

	if reset {
		g.HalfMoveClock = 0
	} else {
		g.HalfMoveClock++
	}

	g.Positions = append(g.Positions, p)
	g.Moves = append(g.Moves, m)
}
