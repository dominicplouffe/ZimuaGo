package main

import (
	"fmt"
	"testing"

	"github.com/dominicplouffe/chess"
)

func TestGetKnightMobilitySquares(t *testing.T) {

	fen, _ := chess.FEN("7k/8/8/8/3N4/8/8/K7 w - - 0 1")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	var bitboards []uint64 = game.Position().Board().Bitboards()
	var bbWhiteKing uint64 = bitboards[0]
	var bbWhiteQueen uint64 = bitboards[1]
	var bbWhiteRook uint64 = bitboards[2]
	var bbWhiteBishop uint64 = bitboards[3]
	var bbWhiteKnight uint64 = bitboards[4]
	var bbWhitePawn uint64 = bitboards[5]
	var allWhiteBBs uint64 = bbWhiteKing | bbWhiteQueen | bbWhiteRook | bbWhiteBishop | bbWhiteKnight | bbWhitePawn

	mobValue, _ := getKnightMobilitySquares(bbWhiteKnight, allWhiteBBs)

	if mobValue != 8 {
		t.Error(fmt.Sprintf("Knight Mobility is incorrect: %v should be 8", mobValue))
	}

	// *******************

	fen, _ = chess.FEN("7k/8/2P1P3/5P2/3N4/5P2/8/K7 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	bitboards = game.Position().Board().Bitboards()
	bbWhiteKing = bitboards[0]
	bbWhiteQueen = bitboards[1]
	bbWhiteRook = bitboards[2]
	bbWhiteBishop = bitboards[3]
	bbWhiteKnight = bitboards[4]
	bbWhitePawn = bitboards[5]
	allWhiteBBs = bbWhiteKing | bbWhiteQueen | bbWhiteRook | bbWhiteBishop | bbWhiteKnight | bbWhitePawn

	mobValue, _ = getKnightMobilitySquares(bbWhiteKnight, allWhiteBBs)
	if mobValue != 4 {
		t.Error(fmt.Sprintf("Knight Mobility is incorrect: %v should be 4", mobValue))
	}

	// *******************

	fen, _ = chess.FEN("7k/8/2P1P3/1P3P2/3N4/1P3P2/2P1P3/K7 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	bitboards = game.Position().Board().Bitboards()
	bbWhiteKing = bitboards[0]
	bbWhiteQueen = bitboards[1]
	bbWhiteRook = bitboards[2]
	bbWhiteBishop = bitboards[3]
	bbWhiteKnight = bitboards[4]
	bbWhitePawn = bitboards[5]
	allWhiteBBs = bbWhiteKing | bbWhiteQueen | bbWhiteRook | bbWhiteBishop | bbWhiteKnight | bbWhitePawn

	mobValue, _ = getKnightMobilitySquares(bbWhiteKnight, allWhiteBBs)
	if mobValue != 0 {
		t.Error(fmt.Sprintf("Knight Mobility is incorrect: %v should be 0", mobValue))
	}

	// *******************

	fen, _ = chess.FEN("7k/8/8/8/8/8/6N1/K7 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	bitboards = game.Position().Board().Bitboards()
	bbWhiteKing = bitboards[0]
	bbWhiteQueen = bitboards[1]
	bbWhiteRook = bitboards[2]
	bbWhiteBishop = bitboards[3]
	bbWhiteKnight = bitboards[4]
	bbWhitePawn = bitboards[5]
	allWhiteBBs = bbWhiteKing | bbWhiteQueen | bbWhiteRook | bbWhiteBishop | bbWhiteKnight | bbWhitePawn

	mobValue, _ = getKnightMobilitySquares(bbWhiteKnight, allWhiteBBs)
	if mobValue != 4 {
		t.Error(fmt.Sprintf("Knight Mobility is incorrect: %v should be 4", mobValue))
	}

	// *******************

	fen, _ = chess.FEN("7k/8/8/8/8/8/8/K5N1 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	bitboards = game.Position().Board().Bitboards()
	bbWhiteKing = bitboards[0]
	bbWhiteQueen = bitboards[1]
	bbWhiteRook = bitboards[2]
	bbWhiteBishop = bitboards[3]
	bbWhiteKnight = bitboards[4]
	bbWhitePawn = bitboards[5]
	allWhiteBBs = bbWhiteKing | bbWhiteQueen | bbWhiteRook | bbWhiteBishop | bbWhiteKnight | bbWhitePawn

	mobValue, _ = getKnightMobilitySquares(bbWhiteKnight, allWhiteBBs)
	if mobValue != 3 {
		t.Error(fmt.Sprintf("Knight Mobility is incorrect: %v should be 3", mobValue))
	}

	// *******************

	fen, _ = chess.FEN("N6k/8/8/8/8/8/8/K7 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	bitboards = game.Position().Board().Bitboards()
	bbWhiteKing = bitboards[0]
	bbWhiteQueen = bitboards[1]
	bbWhiteRook = bitboards[2]
	bbWhiteBishop = bitboards[3]
	bbWhiteKnight = bitboards[4]
	bbWhitePawn = bitboards[5]
	allWhiteBBs = bbWhiteKing | bbWhiteQueen | bbWhiteRook | bbWhiteBishop | bbWhiteKnight | bbWhitePawn

	mobValue, _ = getKnightMobilitySquares(bbWhiteKnight, allWhiteBBs)
	if mobValue != 2 {
		t.Error(fmt.Sprintf("Knight Mobility is incorrect: %v should be 2", mobValue))
	}

	// *******************

	fen, _ = chess.FEN("7k/8/2p1p3/1p3p2/3N4/1p3p2/2p1p3/K7 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	bitboards = game.Position().Board().Bitboards()
	bbWhiteKing = bitboards[0]
	bbWhiteQueen = bitboards[1]
	bbWhiteRook = bitboards[2]
	bbWhiteBishop = bitboards[3]
	bbWhiteKnight = bitboards[4]
	bbWhitePawn = bitboards[5]
	allWhiteBBs = bbWhiteKing | bbWhiteQueen | bbWhiteRook | bbWhiteBishop | bbWhiteKnight | bbWhitePawn

	mobValue, _ = getKnightMobilitySquares(bbWhiteKnight, allWhiteBBs)
	if mobValue != 8 {
		t.Error(fmt.Sprintf("Knight Mobility is incorrect: %v should be 8", mobValue))
	}
}

func TestGetRookMobilitySquares(t *testing.T) {

	fen, _ := chess.FEN("k7/8/8/8/8/K7/8/7R w - - 0 1")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	var bitboards []uint64 = game.Position().Board().Bitboards()
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

	mobValue, _ := getRookMobilitySquares(bbWhiteRook, allWhiteBBs, allBlackBBs)

	if mobValue != 14 {
		t.Error(fmt.Sprintf("Rook Mobility is incorrect: %v should be 14", mobValue))
	}

	// *******************

	fen, _ = chess.FEN("k7/8/8/4R3/8/K7/8/8 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	bitboards = game.Position().Board().Bitboards()
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

	mobValue, _ = getRookMobilitySquares(bbWhiteRook, allWhiteBBs, allBlackBBs)

	if mobValue != 14 {
		t.Error(fmt.Sprintf("Rook Mobility is incorrect: %v should be 14", mobValue))
	}

	// *******************

	fen, _ = chess.FEN("k7/8/4p3/4R3/4p3/K7/8/8 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	bitboards = game.Position().Board().Bitboards()
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

	mobValue, _ = getRookMobilitySquares(bbWhiteRook, allWhiteBBs, allBlackBBs)

	if mobValue != 9 {
		t.Error(fmt.Sprintf("Rook Mobility is incorrect: %v should be 9", mobValue))
	}

	// *******************

	fen, _ = chess.FEN("k7/8/4p3/3pRp2/4p3/K7/8/8 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	bitboards = game.Position().Board().Bitboards()
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

	mobValue, _ = getRookMobilitySquares(bbWhiteRook, allWhiteBBs, allBlackBBs)

	if mobValue != 4 {
		t.Error(fmt.Sprintf("Rook Mobility is incorrect: %v should be 4", mobValue))
	}

	// *******************

	fen, _ = chess.FEN("k7/4P3/8/2P1R1P1/8/K3P3/8/8 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	bitboards = game.Position().Board().Bitboards()
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

	mobValue, _ = getRookMobilitySquares(bbWhiteRook, allWhiteBBs, allBlackBBs)

	if mobValue != 4 {
		t.Error(fmt.Sprintf("Rook Mobility is incorrect: %v should be 4", mobValue))
	}
}

func TestGetBishopMobilitySquares(t *testing.T) {
	fen, _ := chess.FEN("k7/8/8/4B3/8/K7/8/8 w - - 0 1")
	game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	var bitboards []uint64 = game.Position().Board().Bitboards()
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

	mobValue, _ := getBishopMobilitySquares(bbWhiteBishop, allWhiteBBs, allBlackBBs)

	if mobValue != 13 {
		t.Error(fmt.Sprintf("Bishop Mobility is incorrect: %v should be 13", mobValue))
	}

	// *******************

	fen, _ = chess.FEN("k7/8/8/8/8/K7/7B/8 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	bitboards = game.Position().Board().Bitboards()
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

	mobValue, _ = getBishopMobilitySquares(bbWhiteBishop, allWhiteBBs, allBlackBBs)

	if mobValue != 7 {
		t.Error(fmt.Sprintf("Bishop Mobility is incorrect: %v should be 7", mobValue))
	}

	// *******************

	fen, _ = chess.FEN("k7/8/8/8/8/K5P1/7B/6P1 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	bitboards = game.Position().Board().Bitboards()
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

	mobValue, _ = getBishopMobilitySquares(bbWhiteBishop, allWhiteBBs, allBlackBBs)

	if mobValue != 0 {
		t.Error(fmt.Sprintf("Bishop Mobility is incorrect: %v should be 0", mobValue))
	}

	// *******************

	fen, _ = chess.FEN("k7/8/8/8/8/K5p1/7B/6p1 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	bitboards = game.Position().Board().Bitboards()
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

	mobValue, _ = getBishopMobilitySquares(bbWhiteBishop, allWhiteBBs, allBlackBBs)

	if mobValue != 2 {
		t.Error(fmt.Sprintf("Bishop Mobility is incorrect: %v should be 2", mobValue))
	}

	// *******************

	fen, _ = chess.FEN("k7/8/2p3p1/8/4B3/K7/2p3p1/8 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	bitboards = game.Position().Board().Bitboards()
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

	mobValue, _ = getBishopMobilitySquares(bbWhiteBishop, allWhiteBBs, allBlackBBs)

	if mobValue != 8 {
		t.Error(fmt.Sprintf("Bishop Mobility is incorrect: %v should be 8", mobValue))
	}

	// *******************

	fen, _ = chess.FEN("k7/8/2P3P1/8/4B3/K7/2P3P1/8 w - - 0 1")
	game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	bitboards = game.Position().Board().Bitboards()
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

	mobValue, _ = getBishopMobilitySquares(bbWhiteBishop, allWhiteBBs, allBlackBBs)

	if mobValue != 4 {
		t.Error(fmt.Sprintf("Bishop Mobility is incorrect: %v should be 4", mobValue))
	}
}
