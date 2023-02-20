package chess

// A Game contains the history of a chess game.
//
// FEN and PGN parsing and generation are mostly compliant with PGN Standard
// Version 1.0 (Revised 2020-06-03).
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

// NewGameFromFEN returns a new game from a FEN string.
func NewGameFromFEN(fen string) (*Game, error) {
	panic("not implemented")
}

// NewGameFromPGN returns a new game from a PGN string.
func NewGameFromPGN(pgn string) (*Game, error) {
	panic("not implemented")
}

// StartingPosition returns the starting position of the game.
func (g *Game) StartingPosition() Position {
	return g.Positions[0]
}

// CurrentPosition returns the latest position of the game.
func (g *Game) CurrentPosition() Position {
	return g.Positions[len(g.Positions)-1]
}

// FEN returns a FEN string representing the current position.
func (g *Game) FEN() (string, error) {
	panic("not implemented")
}

// PGN returns a PGN string representing the game.
func (g *Game) PGN() (string, error) {
	panic("not implemented")
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
