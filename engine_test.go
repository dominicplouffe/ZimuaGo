package main

import (
	"testing"

	"github.com/dominicplouffe/chess"
)

// Test Example
func TestPieceScoring(t *testing.T) {

	zg := Zimua("White", 5.0)

	fen, _ := chess.FEN("r3kb1r/1p3ppp/pn6/2p5/2bP1Q2/1B6/PP1PqPPP/R1B3KR b - - 1 2")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	score := zg.pieceScoring(game.Position().Board())

	if score != -228 {
		t.Error("Piece scoring should be -228")
	}

}

// Testing the MinMax Scoring to 1, 2 and 3 depth
func TestMinMax(t *testing.T) {
	zg := Zimua("White", 5.0)

	fen, _ := chess.FEN("r1bqkbnr/ppp1pppp/2n5/3p4/3P4/5N2/PPP1PPPP/RNBQKB1R w KQkq - 0 3")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	res := zg.alphaBeta(game.Position(), 1, -9999999999, 9999999999, true, 1, false, false)
	if res.move.String() != "e2e4" || res.score != 55 {
		t.Error("Depth 1 is incorrect")
	}

	res = zg.alphaBeta(game.Position(), 2, -9999999999, 9999999999, true, 2, false, false)
	if res.move.String() != "e2e4" || res.score != -77 {
		t.Error("Depth 2 is incorrect")
	}

	res = zg.alphaBeta(game.Position(), 3, -9999999999, 9999999999, true, 3, false, false)
	if res.move.String() != "b1c3" || res.score != 123 {
		t.Error("Depth 3 is incorrect")
	}

}

// Testing the NegaMax Scoring to 1, 2 and 3 depth
func TestNegaMax(t *testing.T) {
	zg := Zimua("White", 5.0)

	fen, _ := chess.FEN("r1bqkbnr/ppp1pppp/2n5/3p4/3P4/5N2/PPP1PPPP/RNBQKB1R w KQkq - 0 3")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	res := zg.alphaBetaNM(game.Position(), 1, -9999999999, 9999999999, true, 1, false, false)
	if res.move.String() != "e2e4" || res.score != 55 {
		t.Error("Depth 1 NM is incorrect")
	}

	res = zg.alphaBetaNM(game.Position(), 2, -9999999999, 9999999999, true, 2, false, false)
	if res.move.String() != "e2e4" || res.score != -77 {
		t.Error("Depth 2 is incorrect")
	}

	res = zg.alphaBetaNM(game.Position(), 3, -9999999999, 9999999999, true, 3, false, false)
	if res.move.String() != "b1c3" || res.score != 123 {
		t.Error("Depth 3 is incorrect")
	}
}

// Some testing for null moves
func TestNullMove(t *testing.T) {
	game := chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))

	newPos := game.Position().NullMove()

	if newPos.Turn() != chess.Black {
		t.Error("Null Move did not work")
	}

}

func TestMoveScoring(t *testing.T) {
	zg := Zimua("White", 5.0)

	fen, _ := chess.FEN("r1bqkbnr/ppp1pppp/2n5/3pN3/3P4/8/PPP1PPPP/RNBQKB1R w KQkq - 0 1")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	legalMoves := zg.getMoves(game.Position(), 1)

	prev := legalMoves[0].score
	for i, m := range legalMoves {
		if i == 0 {
			continue
		}
		if m.score > prev {
			t.Error("Invalid legal moves")
		}
		prev = m.score
	}
}

func TestOpenningMoves(t *testing.T) {
	zg := Zimua("White", 5.0)

	game := chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))
	zg.openingMove(game)
}
