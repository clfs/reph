// Package chess implements basic chess functionality.
package chess

// An Outcome represents the outcome of a game.
type Outcome int

const (
	Undecided Outcome = iota
	BlackWon
	Draw
	WhiteWon
)
