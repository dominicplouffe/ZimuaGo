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

//TestCheckmat1 should produce a checkmate
func TestCheckmate1(t *testing.T) {
	zg := Zimua("White", 5.0)
	zg.doOpen = false

	fen, _ := chess.FEN("rn1k2nr/p1qpp1Q1/1p3pBP/2pP4/P4B2/2P5/1P3PP1/R3K2R w KQ - 5 22")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	_, mv := zg.evaluate(game, false)

	if mv.String() != "g7f8" {
		t.Error("Move should be g7f8")
	}
	if game.Outcome() != chess.WhiteWon {
		t.Error("Move should cause a checkmate")
	}
	// fmt.Println("Outcome: ", game.Outcome())
}

//TestCheckmat2 should produce a checkmate
func TestCheckmate2(t *testing.T) {
	zg := Zimua("White", 5.0)
	zg.doOpen = false

	fen, _ := chess.FEN("8/3R1Q2/8/8/4k1B1/7P/P4PP1/4K2R w K - 15 46")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	_, mv := zg.evaluate(game, false)

	if mv.String() != "f7f5" {
		t.Error("Move should be g7f8")
	}
	if game.Outcome() != chess.WhiteWon {
		t.Error("Move should cause a checkmate")
	}
	// fmt.Println("Outcome: ", game.Outcome())
}
