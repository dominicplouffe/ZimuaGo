package main

import (
	"fmt"
	"testing"

	"github.com/dominicplouffe/chess"
)

func TestKingMobility(t *testing.T) {

	initSquareIndexes()

	// zg := Zimua("White", 5.0)

	fen, _ := chess.FEN("k7/8/8/2PPP3/2NKN3/2RQR3/8/8 w - - 0 1")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	b := game.Position().Board()

	var bitboards []uint64 = b.Bitboards()
	var bbWhiteKing uint64 = bitboards[0]
	var bbWhiteQueen uint64 = bitboards[1]
	var bbWhiteRook uint64 = bitboards[2]
	var bbWhiteBishop uint64 = bitboards[3]
	var bbWhiteKnight uint64 = bitboards[4]
	var bbWhitePawn uint64 = bitboards[5]
	var bbBlackKing uint64 = bitboards[6]
	var bbBlackQueen uint64 = bitboards[7]
	var bbBlackRook uint64 = bitboards[8]
	var bbBlackBishop uint64 = bitboards[9]
	var bbBlackKnight uint64 = bitboards[10]
	var bbBlackPawn uint64 = bitboards[11]
	var allWhiteBBs uint64 = bbWhiteKing | bbWhiteQueen | bbWhiteRook | bbWhiteBishop | bbWhiteKnight | bbWhitePawn
	var allBlackBBs uint64 = bbBlackKing | bbBlackQueen | bbBlackRook | bbBlackBishop | bbBlackKnight | bbBlackPawn

	_ = allWhiteBBs
	_ = allBlackBBs

	mob := getKingMobility(bbWhiteKing, allWhiteBBs)
	if mob != 8 {
		t.Error("King mobility should be 8", mob)
	}

	fen, _ = chess.FEN("k7/8/8/2PPP3/3K4/8/8/8 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	b = game.Position().Board()

	bitboards = b.Bitboards()
	bbWhiteKing = bitboards[0]
	bbWhiteQueen = bitboards[1]
	bbWhiteRook = bitboards[2]
	bbWhiteBishop = bitboards[3]
	bbWhiteKnight = bitboards[4]
	bbWhitePawn = bitboards[5]
	bbBlackKing = bitboards[6]
	bbBlackQueen = bitboards[7]
	bbBlackRook = bitboards[8]
	bbBlackBishop = bitboards[9]
	bbBlackKnight = bitboards[10]
	bbBlackPawn = bitboards[11]
	allWhiteBBs = bbWhiteKing | bbWhiteQueen | bbWhiteRook | bbWhiteBishop | bbWhiteKnight | bbWhitePawn
	allBlackBBs = bbBlackKing | bbBlackQueen | bbBlackRook | bbBlackBishop | bbBlackKnight | bbBlackPawn

	_ = allWhiteBBs
	_ = allBlackBBs

	mob = getKingMobility(bbWhiteKing, allWhiteBBs)
	if mob != 3 {
		t.Error("King mobility should be 3", mob)
	}

	// Test 3
	fen, _ = chess.FEN("k7/8/8/8/8/8/8/7K w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	b = game.Position().Board()

	bitboards = b.Bitboards()
	bbWhiteKing = bitboards[0]
	bbWhiteQueen = bitboards[1]
	bbWhiteRook = bitboards[2]
	bbWhiteBishop = bitboards[3]
	bbWhiteKnight = bitboards[4]
	bbWhitePawn = bitboards[5]
	bbBlackKing = bitboards[6]
	bbBlackQueen = bitboards[7]
	bbBlackRook = bitboards[8]
	bbBlackBishop = bitboards[9]
	bbBlackKnight = bitboards[10]
	bbBlackPawn = bitboards[11]
	allWhiteBBs = bbWhiteKing | bbWhiteQueen | bbWhiteRook | bbWhiteBishop | bbWhiteKnight | bbWhitePawn
	allBlackBBs = bbBlackKing | bbBlackQueen | bbBlackRook | bbBlackBishop | bbBlackKnight | bbBlackPawn

	_ = allWhiteBBs
	_ = allBlackBBs

	mob = getKingMobility(bbWhiteKing, allWhiteBBs)
	if mob != 5 {
		t.Error("King mobility should be 5", mob)
	}
}

func TestPieceInCheckmate(t *testing.T) {

	zg := Zimua("White", 5.0)

	fen, _ := chess.FEN("r3kb1r/1p3ppp/pn6/2p5/2bP1Q2/1B6/PP1P1PPP/R1B1q1KR w - - 2 3")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	score := zg.pieceScoring(game.Position())

	if score != 99999999 {
		t.Error("Piece scoring should be 99999999", score)
	}
}

func TestPieceInFull(t *testing.T) {

	zg := Zimua("White", 5.0)

	fen, _ := chess.FEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	score := zg.pieceScoring(game.Position())

	if score != 0 {
		t.Error("Piece scoring should be 0", score)
	}
}

func TestPieceScoring(t *testing.T) {

	zg := Zimua("White", 5.0)

	fen, _ := chess.FEN("rq2k2r/pp1nbpp1/2p1pn2/2Pp4/3P2pP/2N1PP2/PPQ3P1/1K1RBB1R w Kkq - 0 1")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	score := zg.pieceScoring(game.Position())

	if score != 190 {
		t.Error("Piece scoring should be 190", score)
	}
}

func TestPieceInCheckmateQsearch(t *testing.T) {

	zg := Zimua("White", 5.0)

	fen, _ := chess.FEN("r3kb1r/1p3ppp/pn6/2p5/2bP1Q2/1B6/PP1P1PPP/R1B1q1KR w - - 2 3")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	score := zg.qsearch(game.Position())

	if score != 99999999 {
		t.Error("Piece scoring should be 99999999", score)
	}
}

func TestPieceInCheckmateAlphaBeta(t *testing.T) {

	zg := Zimua("White", 5.0)

	fen, _ := chess.FEN("r3kb1r/1p3ppp/pn6/2p5/2bP1Q2/1B6/PP1P1PPP/R1B1q1KR w - - 2 3")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))
	ply := 6

	siblings := make([]MoveScore, ply)
	res := zg.alphaBetaNM(
		game.Position(),
		ply,
		zg.minValue,
		zg.maxValue,
		ply,
		true,
		false,
		siblings,
	)

	if res.score != 99999999 {
		t.Error("Piece scoring should be 99999999", res.score)
	}
}

func TestPieceInCheckmateAnotherQsearch(t *testing.T) {

	zg := Zimua("White", 5.0)

	fen, _ := chess.FEN("8/2R5/1k6/1pNP4/1P6/P4P2/8/1K6 w - - 3 73")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	_, move := zg.search(game, false)

	if move.S1() != chess.C7 || move.S2() != chess.B7 {
		t.Error("Move Square should be C7 to B7")
	}
}

func TestPieceInCheckmateAlphaBetaNextMoveCMBlack(t *testing.T) {

	zg := Zimua("White", 5.0)

	fen, _ := chess.FEN("r3kb1r/1p3ppp/pn6/2p5/2bP1Q2/1B6/PP1PqPPP/R1B3KR b - - 1 2")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))
	ply := 1

	siblings := make([]MoveScore, ply)
	res := zg.alphaBetaNM(
		game.Position(),
		ply,
		zg.minValue,
		zg.maxValue,
		ply,
		false,
		false,
		siblings,
	)

	// smoves := ""
	// for i := len(siblings) - 1; i >= 0; i-- {
	// 	smoves += siblings[i].move.String() + " "
	// }
	// fmt.Println(smoves)

	if res.score != 99999999 {
		t.Error("Piece scoring should be 99999999", res.score)
	}
}

func TestPieceInCheckmateAlphaBetaNextMoveCMWhite(t *testing.T) {

	zg := Zimua("White", 5.0)

	fen, _ := chess.FEN("k7/8/8/8/5Q2/1R6/8/7K w - - 0 1")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))
	ply := 2

	siblings := make([]MoveScore, ply)
	res := zg.alphaBetaNM(
		game.Position(),
		ply,
		zg.minValue,
		zg.maxValue,
		ply,
		false,
		false,
		siblings,
	)

	if res.score != 99999999 {
		t.Error("Piece scoring should be 99999999", res.score)
	}
}

func TestPieceInCheckmateAlphaBetaNextMoveCMWhite2(t *testing.T) {

	zg := Zimua("White", 5.0)

	fen, _ := chess.FEN("8/1k6/8/8/5Q2/2R5/8/7K w - - 2 2")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	// moves := zg.getMoves(game.Position(), 0)

	// for _, mv := range moves {
	// 	fmt.Println(mv.move.S1(), mv.move.S2())
	// }

	ply := 5

	siblings := make([]MoveScore, ply)
	res := zg.alphaBetaNM(
		game.Position(),
		ply,
		zg.minValue,
		zg.maxValue,
		ply,
		false,
		false,
		siblings,
	)

	smoves := ""
	for i := len(siblings) - 1; i >= 0; i-- {
		smoves += siblings[i].move.String() + " "
	}
	fmt.Println(smoves)

	if res.score != 99999999 {
		t.Error("Piece scoring should be 99999999", res.score)
	}
}
