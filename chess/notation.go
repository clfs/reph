package chess

import (
	"fmt"
	"strings"
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

func fenToBoard(s string) (Board, bool) {
	var b Board

	sq := A8
	for _, r := range s {
		switch r {
		case '1', '2', '3', '4', '5', '6', '7', '8':
			sq = sq.NextN(int(r - '0'))
		case '/':
			sq = sq.PrevN(16)
		default:
			piece, ok := fenToPiece[string(r)]
			if !ok {
				return Board{}, false
			}
			b.SetPiece(piece, sq)
			sq = sq.Next()
		}
	}

	return b, true
}

func boardToFEN(b Board) string {
	var sb strings.Builder

	for r := Rank8; r <= Rank8; r-- {
		skip := 0
		for f := FileA; f <= FileH; f++ {
			piece, ok := b.Get(NewSquare(f, r))
			if !ok {
				skip++
				continue
			}
			if skip > 0 {
				fmt.Fprintf(&sb, "%d", skip)
				skip = 0
			}
			fmt.Fprintf(&sb, "%s", pieceToFEN[piece])
		}
		if skip > 0 {
			fmt.Fprintf(&sb, "%d", skip)
		}
		if r != Rank1 {
			fmt.Fprintf(&sb, "/")
		}
	}

	return sb.String()
}

// NewGameFromFEN returns a new game from a FEN string.
func NewGameFromFEN(fen string) (*Game, error) {
	var game Game

	var (
		boardFEN     string
		colorFEN     string
		castleFEN    string
		enPassantFEN string
	)

	n, err := fmt.Sscanf(
		fen,
		"%s %s %s %s %d %d",
		&boardFEN, &colorFEN, &castleFEN, &enPassantFEN, &game.HalfMoveClock, &game.FullMoveNumber,
	)
	if err != nil {
		return nil, fmt.Errorf("bad fmt.Sscanf: %w", err)
	}
	if n != 6 {
		return nil, fmt.Errorf("wrong number of fields: %d", n)
	}

	var (
		p  Position
		ok bool
	)

	p.Board, ok = fenToBoard(boardFEN)
	if !ok {
		return nil, fmt.Errorf("invalid board: %q", boardFEN)
	}
	p.ActiveColor, ok = fenToColor[colorFEN]
	if !ok {
		return nil, fmt.Errorf("invalid color: %q", colorFEN)
	}
	p.CastleRights, ok = fenToCastleRights[castleFEN]
	if !ok {
		return nil, fmt.Errorf("invalid castling rights: %q", castleFEN)
	}
	p.EnPassantRight, ok = fenToEnPassantRight[enPassantFEN]
	if !ok {
		return nil, fmt.Errorf("invalid en passant right: %q", enPassantFEN)
	}

	game.Positions = []Position{p}

	return &game, nil
}

// FEN returns a FEN string representing the current position.
func (g *Game) FEN() string {
	p := g.CurrentPosition()
	return fmt.Sprintf(
		"%s %s %s %s %d %d",
		boardToFEN(p.Board),
		colorToFEN[p.ActiveColor],
		castleRightsToFEN[p.CastleRights],
		enPassantRightToFEN[p.EnPassantRight],
		g.HalfMoveClock,
		g.FullMoveNumber,
	)
}