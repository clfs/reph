package chess

import (
	"fmt"
)

var fenToPiece = map[string]Piece{
	"P": {White, Pawn},
	"N": {White, Knight},
	"B": {White, Bishop},
	"R": {White, Rook},
	"Q": {White, Queen},
	"K": {White, King},
	"p": {Black, Pawn},
	"n": {Black, Knight},
	"b": {Black, Bishop},
	"r": {Black, Rook},
	"q": {Black, Queen},
	"k": {Black, King},
}

var fenToColor = map[string]Color{
	"w": White,
	"b": Black,
}

var fenToCastleRights = map[string]CastleRights{
	"-":    0,
	"K":    WhiteKingSide,
	"Q":    WhiteQueenSide,
	"k":    BlackKingSide,
	"q":    BlackQueenSide,
	"KQ":   WhiteKingSide | WhiteQueenSide,
	"Kk":   WhiteKingSide | BlackKingSide,
	"Kq":   WhiteKingSide | BlackQueenSide,
	"Qk":   WhiteQueenSide | BlackKingSide,
	"Qq":   WhiteQueenSide | BlackQueenSide,
	"kq":   BlackKingSide | BlackQueenSide,
	"KQk":  WhiteKingSide | WhiteQueenSide | BlackKingSide,
	"KQq":  WhiteKingSide | WhiteQueenSide | BlackQueenSide,
	"Kkq":  WhiteKingSide | BlackKingSide | BlackQueenSide,
	"Qkq":  WhiteQueenSide | BlackKingSide | BlackQueenSide,
	"KQkq": WhiteKingSide | WhiteQueenSide | BlackKingSide | BlackQueenSide,
}

var fenToEnPassantRight = map[string]EnPassantRight{
	"-":  {Valid: false},
	"a3": {Square: A3, Valid: true},
	"b3": {Square: B3, Valid: true},
	"c3": {Square: C3, Valid: true},
	"d3": {Square: D3, Valid: true},
	"e3": {Square: E3, Valid: true},
	"f3": {Square: F3, Valid: true},
	"g3": {Square: G3, Valid: true},
	"h3": {Square: H3, Valid: true},
	"a6": {Square: A6, Valid: true},
	"b6": {Square: B6, Valid: true},
	"c6": {Square: C6, Valid: true},
	"d6": {Square: D6, Valid: true},
	"e6": {Square: E6, Valid: true},
	"f6": {Square: F6, Valid: true},
	"g6": {Square: G6, Valid: true},
	"h6": {Square: H6, Valid: true},
}

func invertMap[K, V comparable](m map[K]V) map[V]K {
	out := make(map[V]K)
	for k, v := range m {
		out[v] = k
	}
	return out
}

var (
	pieceToFEN          = invertMap(fenToPiece)
	colorToFEN          = invertMap(fenToColor)
	castleRightsToFEN   = invertMap(fenToCastleRights)
	enPassantRightToFEN = invertMap(fenToEnPassantRight)
)

// NewGameFromFEN returns a new game from a FEN string.
// It returns InvalidFENError on failure.
func NewGameFromFEN(fen string) (*Game, error) {
	var game Game

	var (
		boardFEN     string
		colorFEN     string
		castlingFEN  string
		enPassantFEN string
	)

	n, err := fmt.Sscanf(
		fen,
		"%s %s %s %s %d %d",
		&boardFEN, &colorFEN, &castlingFEN, &enPassantFEN, &game.HalfMoveClock, &game.FullMoveNumber,
	)
	if err != nil {
		return nil, fmt.Errorf("bad fmt.Sscanf: %w", err)
	}
	if n != 6 {
		return nil, fmt.Errorf("wrong number of fields: %d", n)
	}

	board, err := newBoardFromFEN(boardFEN)
	if err != nil {
		return nil, err
	}
	color, err := newColorFromFEN(colorFEN)
	if err != nil {
		return nil, err
	}
	castleRights, err := newCastleRightsFromFEN(castlingFEN)
	if err != nil {
		return nil, err
	}
	enPassantRight, err := newEnPassantRightFromFEN(enPassantFEN)
	if err != nil {
		return nil, err
	}

	game.Positions = []Position{{
		Board:          board,
		ActiveColor:    color,
		CastleRights:   castleRights,
		EnPassantRight: enPassantRight,
	}}

	return &game, nil
}

func newBoardFromFEN(s string) (Board, error) {
	var b Board

	sq := A8
	for _, r := range s {
		switch r {
		case '1', '2', '3', '4', '5', '6', '7', '8':
			sq = sq.NextN(int(r - '0'))
		case '/':
			sq = sq.PrevN(16)
		default:
			piece, err := newPieceFromFEN(string(r))
			if err != nil {
				return Board{}, err
			}
			b.SetPiece(piece, sq)
			sq = sq.Next()
		}
	}

	return b, nil
}

func newPieceFromFEN(s string) (Piece, error) {
	piece, ok := fenToPiece[s]
	if !ok {
		return Piece{}, fmt.Errorf("invalid piece: %q", s)
	}
	return piece, nil
}

func newColorFromFEN(fen string) (Color, error) {
	c, ok := fenToColor[fen]
	if !ok {
		return White, fmt.Errorf("invalid color: %q", fen)
	}
	return c, nil
}

func newCastleRightsFromFEN(s string) (CastleRights, error) {
	r, ok := fenToCastleRights[s]
	if !ok {
		return 0, fmt.Errorf("invalid castle rights: %q", s)
	}
	return r, nil
}

func newEnPassantRightFromFEN(fen string) (EnPassantRight, error) {
	r, ok := fenToEnPassantRight[fen]
	if !ok {
		return EnPassantRight{}, fmt.Errorf("invalid en passant right: %q", fen)
	}
	return r, nil
}

// FEN returns a FEN string representing the current position.
// It returns InvalidGameError on failure.
func (g *Game) FEN() (string, error) {
	panic("not implemented")
}
