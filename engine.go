package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/dominicplouffe/chess"
)

var checkmate = 99999998

// MoveScore is used to store the moves importance when generating the list of moves
type MoveScore struct {
	move      chess.Move
	score     int
	killer    bool
	inCheck   bool
	capture   bool
	promotion bool
}

// TimeControl is used to track the amount of time the engine can use to make a move
type TimeControl struct {
	expectedMoves float64
	totalTime     float64
	remainingTime float64
	timePerMove   float64
	moveCount     float64
	totalElapsed  float64
	totalNodes    int
}

//ZimuaGame defines a chess engine
type ZimuaGame struct {
	posPointsBlack map[int][]int
	posPointsWhite map[int][]int
	piecePoints    map[int]int
	moveCount      int
	moveSearched   int
	nilMove        chess.Move
	timeControl    TimeControl
	inCheck        bool
	tableAge       int
	minValue       int
	maxValue       int
	name           string
	doOpen         bool
}

//Zimua creates an instance of the Zimua chess engine
func Zimua(name string, maxMinutes float64) ZimuaGame {
	zg := ZimuaGame{
		posPointsBlack: make(map[int][]int),
		posPointsWhite: make(map[int][]int),
		piecePoints:    make(map[int]int),
		moveSearched:   0,
		nilMove:        chess.Move{},
		minValue:       -9999999999,
		maxValue:       9999999999,
		timeControl:    getTimeControl(maxMinutes),
		name:           name,
		moveCount:      0,
		doOpen:         true,
	}
	zg.initGame()

	return zg
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func getTimeControl(totalTime float64) TimeControl {
	totalTime = totalTime * 60.0
	expectedMoves := 60.0
	remainingTime := totalTime
	timePerMove := remainingTime / expectedMoves

	tc := TimeControl{
		expectedMoves: expectedMoves,
		totalTime:     totalTime,
		remainingTime: remainingTime,
		timePerMove:   timePerMove,
		moveCount:     0,
		totalElapsed:  0,
		totalNodes:    0,
	}

	log.Println(tc)
	return tc
}

func (zg ZimuaGame) initGame() {
	zg.posPointsBlack[0] = append(zg.posPointsBlack[0], 0, 0, 0, 0, 0, 0, 0, 0, -10, 3, -14, -36, -37, -7, 8, -10, 10, 3, -2, 10, 10, 5, 9, 10, 10, 0, 1, 20, 20, 10, 3, 10, 20, 15, 0, 14, 15, -2, 16, 20, 20, 44, 31, 40, 44, 21, 29, 20, 90, 85, 82, 102, 73, 86, 83, 78, 900, 900, 900, 900, 900, 900, 900, 900)
	zg.posPointsBlack[1] = append(zg.posPointsBlack[1], -69, -22, -35, -19, -24, -26, -23, -74, -20, -23, 0, 2, 0, 2, -15, -23, -14, 11, 25, 18, 22, 25, 10, -18, 0, 2, 35, 22, 21, 31, 5, -1, 17, 25, 41, 33, 37, 45, 24, 24, -2, 62, 100, 73, 74, 100, 67, 10, -14, -4, 62, 4, -36, 100, -6, -3, -70, -58, -55, -10, -75, -75, -53, -66)
	zg.posPointsBlack[2] = append(zg.posPointsBlack[2], -10, -10, -15, -14, -12, -15, 2, -7, 16, 20, 6, 7, 6, 11, 20, 19, 15, 20, 25, 8, 15, 24, 25, 14, 7, 0, 16, 17, 23, 17, 10, 13, 10, 15, 25, 26, 34, 20, 17, 25, -14, 28, -10, 52, 41, -32, 39, -9, -22, 2, 31, -39, -42, 35, 20, -11, -50, -37, -107, -23, -76, -82, -78, -59)
	zg.posPointsBlack[3] = append(zg.posPointsBlack[3], -32, -31, -18, -2, 5, -18, -24, -30, -53, -44, -43, -29, -26, -31, -38, -53, -46, -26, -35, -25, -25, -42, -28, -42, -30, -46, -29, -13, -21, -16, -35, -28, -6, -9, -4, 18, 13, 16, 5, 0, 15, 25, 27, 45, 33, 28, 35, 19, 60, 34, 62, 55, 67, 56, 29, 55, 50, 56, 33, 37, 4, 33, 29, 35)
	zg.posPointsBlack[4] = append(zg.posPointsBlack[4], -42, -34, -36, -31, -13, -31, -30, -39, -38, -21, -15, -15, -19, 0, -18, -36, -27, -16, -11, -16, -11, -13, -6, -30, -22, -20, -10, -1, -5, -2, -15, -14, -6, -13, 20, 25, 17, 22, -16, 1, 2, 43, 63, 72, 60, 32, 43, -2, 24, 57, 76, 20, -10, 60, 32, 14, 26, 88, 24, 69, -104, -8, 1, 6)
	zg.posPointsBlack[5] = append(zg.posPointsBlack[5], 18, 40, -1, 6, -14, -3, 30, 17, 4, 13, -18, -57, -50, -14, 3, -4, -32, -29, -32, -64, -79, -43, -42, -47, -50, -8, -47, -51, -28, -52, -43, -55, -49, 0, 13, -19, -4, 11, 50, -55, -31, 37, 28, -67, 44, -57, 12, -62, 3, 10, 55, 56, 56, 55, 10, -32, -62, 83, 60, -99, -99, 47, 54, 4)

	zg.posPointsWhite[0] = append(zg.posPointsWhite[0], 900, 900, 900, 900, 900, 900, 900, 900, 90, 85, 82, 102, 73, 86, 83, 78, 20, 44, 31, 40, 44, 21, 29, 20, 20, 15, 0, 14, 15, -2, 16, 20, 10, 0, 1, 20, 20, 10, 3, 10, 10, 3, -2, 10, 10, 5, 9, 10, -10, 3, -14, -36, -37, -7, 8, -10, 0, 0, 0, 0, 0, 0, 0, 0)
	zg.posPointsWhite[1] = append(zg.posPointsWhite[1], -70, -58, -55, -10, -75, -75, -53, -66, -14, -4, 62, 4, -36, 100, -6, -3, -2, 62, 100, 73, 74, 100, 67, 10, 17, 25, 41, 33, 37, 45, 24, 24, 0, 2, 35, 22, 21, 31, 5, -1, -14, 11, 25, 18, 22, 25, 10, -18, -20, -23, 0, 2, 0, 2, -15, -23, -69, -22, -35, -19, -24, -26, -23, -74)
	zg.posPointsWhite[2] = append(zg.posPointsWhite[2], -50, -37, -107, -23, -76, -82, -78, -59, -22, 2, 31, -39, -42, 35, 20, -11, -14, 28, -10, 52, 41, -32, 39, -9, 10, 15, 25, 26, 34, 20, 17, 25, 7, 0, 16, 17, 23, 17, 10, 13, 15, 20, 25, 8, 15, 24, 25, 14, 16, 20, 6, 7, 6, 11, 20, 19, -10, -10, -15, -14, -12, -15, 2, -7)
	zg.posPointsWhite[3] = append(zg.posPointsWhite[3], 50, 56, 33, 37, 4, 33, 29, 35, 60, 34, 62, 55, 67, 56, 29, 55, 15, 25, 27, 45, 33, 28, 35, 19, -6, -9, -4, 18, 13, 16, 5, 0, -30, -46, -29, -13, -21, -16, -35, -28, -46, -26, -35, -25, -25, -42, -28, -42, -53, -44, -43, -29, -26, -31, -38, -53, -32, -31, -18, -2, 5, -18, -24, -30)
	zg.posPointsWhite[4] = append(zg.posPointsWhite[4], 26, 88, 24, 69, -104, -8, 1, 6, 24, 57, 76, 20, -10, 60, 32, 14, 2, 43, 63, 72, 60, 32, 43, -2, -6, -13, 20, 25, 17, 22, -16, 1, -22, -20, -10, -1, -5, -2, -15, -14, -27, -16, -11, -16, -11, -13, -6, -30, -38, -21, -15, -15, -19, 0, -18, -36, -42, -34, -36, -31, -13, -31, -30, -39)
	zg.posPointsWhite[5] = append(zg.posPointsWhite[5], -62, 83, 60, -99, -99, 47, 54, 4, 3, 10, 55, 56, 56, 55, 10, -32, -31, 37, 28, -67, 44, -57, 12, -62, -49, 0, 13, -19, -4, 11, 50, -55, -50, -8, -47, -51, -28, -52, -43, -55, -32, -29, -32, -64, -79, -43, -42, -47, 4, 13, -18, -57, -50, -14, 3, -4, 18, 40, -1, 6, -14, -3, 30, 17)

	// zg.posPointsBlack[0] = append(zg.posPointsBlack[0], 900, 900, 900, 900, 900, 900, 900, 900, 78, 83, 86, 73, 102, 82, 85, 90, 20, 29, 21, 44, 40, 31, 44, 20, 20, 16, -2, 15, 14, 0, 15, 20, 10, 3, 10, 20, 20, 1, 0, 10, 10, 9, 5, 10, 10, -2, 3, 10, -10, 8, -7, -37, -36, -14, 3, -10, 0, 0, 0, 0, 0, 0, 0, 0)
	// zg.posPointsBlack[1] = append(zg.posPointsBlack[1], -66, -53, -75, -75, -10, -55, -58, -70, -3, -6, 100, -36, 4, 62, -4, -14, 10, 67, 100, 74, 73, 100, 62, -2, 24, 24, 45, 37, 33, 41, 25, 17, -1, 5, 31, 21, 22, 35, 2, 0, -18, 10, 25, 22, 18, 25, 11, -14, -23, -15, 2, 0, 2, 0, -23, -20, -74, -23, -26, -24, -19, -35, -22, -69)
	// zg.posPointsBlack[2] = append(zg.posPointsBlack[2], -59, -78, -82, -76, -23, -107, -37, -50, -11, 20, 35, -42, -39, 31, 2, -22, -9, 39, -32, 41, 52, -10, 28, -14, 25, 17, 20, 34, 26, 25, 15, 10, 13, 10, 17, 23, 17, 16, 0, 7, 14, 25, 24, 15, 8, 25, 20, 15, 19, 20, 11, 6, 7, 6, 20, 16, -7, 2, -15, -12, -14, -15, -10, -10)
	// zg.posPointsBlack[3] = append(zg.posPointsBlack[3], 35, 29, 33, 4, 37, 33, 56, 50, 55, 29, 56, 67, 55, 62, 34, 60, 19, 35, 28, 33, 45, 27, 25, 15, 0, 5, 16, 13, 18, -4, -9, -6, -28, -35, -16, -21, -13, -29, -46, -30, -42, -28, -42, -25, -25, -35, -26, -46, -53, -38, -31, -26, -29, -43, -44, -53, -30, -24, -18, 5, -2, -18, -31, -32)
	// zg.posPointsBlack[4] = append(zg.posPointsBlack[4], 6, 1, -8, -104, 69, 24, 88, 26, 14, 32, 60, -10, 20, 76, 57, 24, -2, 43, 32, 60, 72, 63, 43, 2, 1, -16, 22, 17, 25, 20, -13, -6, -14, -15, -2, -5, -1, -10, -20, -22, -30, -6, -13, -11, -16, -11, -16, -27, -36, -18, 0, -19, -15, -15, -21, -38, -39, -30, -31, -13, -31, -36, -34, -42)
	// zg.posPointsBlack[5] = append(zg.posPointsBlack[5], 4, 54, 47, -99, -99, 60, 83, -62, -32, 10, 55, 56, 56, 55, 10, 3, -62, 12, -57, 44, -67, 28, 37, -31, -55, 50, 11, -4, -19, 13, 0, -49, -55, -43, -52, -28, -51, -47, -8, -50, -47, -42, -43, -79, -64, -32, -29, -32, -4, 3, -14, -50, -57, -18, 13, 4, 17, 30, -3, -14, 6, -1, 40, 18)

	// zg.posPointsWhite[0] = append(zg.posPointsWhite[0], 0, 0, 0, 0, 0, 0, 0, 0, -10, 8, -7, -37, -36, -14, 3, -10, 10, 9, 5, 10, 10, -2, 3, 10, 10, 3, 10, 20, 20, 1, 0, 10, 20, 16, -2, 15, 14, 0, 15, 20, 20, 29, 21, 44, 40, 31, 44, 20, 78, 83, 86, 73, 102, 82, 85, 90, 900, 900, 900, 900, 900, 900, 900, 900)
	// zg.posPointsWhite[1] = append(zg.posPointsWhite[1], -74, -23, -26, -24, -19, -35, -22, -69, -23, -15, 2, 0, 2, 0, -23, -20, -18, 10, 25, 22, 18, 25, 11, -14, -1, 5, 31, 21, 22, 35, 2, 0, 24, 24, 45, 37, 33, 41, 25, 17, 10, 67, 100, 74, 73, 100, 62, -2, -3, -6, 100, -36, 4, 62, -4, -14, -66, -53, -75, -75, -10, -55, -58, -70)
	// zg.posPointsWhite[2] = append(zg.posPointsWhite[2], -7, 2, -15, -12, -14, -15, -10, -10, 19, 20, 11, 6, 7, 6, 20, 16, 14, 25, 24, 15, 8, 25, 20, 15, 13, 10, 17, 23, 17, 16, 0, 7, 25, 17, 20, 34, 26, 25, 15, 10, -9, 39, -32, 41, 52, -10, 28, -14, -11, 20, 35, -42, -39, 31, 2, -22, -59, -78, -82, -76, -23, -107, -37, -50)
	// zg.posPointsWhite[3] = append(zg.posPointsWhite[3], -30, -24, -18, 5, -2, -18, -31, -32, -53, -38, -31, -26, -29, -43, -44, -53, -42, -28, -42, -25, -25, -35, -26, -46, -28, -35, -16, -21, -13, -29, -46, -30, 0, 5, 16, 13, 18, -4, -9, -6, 19, 35, 28, 33, 45, 27, 25, 15, 55, 29, 56, 67, 55, 62, 34, 60, 35, 29, 33, 4, 37, 33, 56, 50)
	// zg.posPointsWhite[4] = append(zg.posPointsWhite[4], -39, -30, -31, -13, -31, -36, -34, -42, -36, -18, 0, -19, -15, -15, -21, -38, -30, -6, -13, -11, -16, -11, -16, -27, -14, -15, -2, -5, -1, -10, -20, -22, 1, -16, 22, 17, 25, 20, -13, -6, -2, 43, 32, 60, 72, 63, 43, 2, 14, 32, 60, -10, 20, 76, 57, 24, 6, 1, -8, -104, 69, 24, 88, 26)
	// zg.posPointsWhite[5] = append(zg.posPointsWhite[5], 17, 30, -3, -14, 6, -1, 40, 18, -4, 3, -14, -50, -57, -18, 13, 4, -47, -42, -43, -79, -64, -32, -29, -32, -55, -43, -52, -28, -51, -47, -8, -50, -55, 50, 11, -4, -19, 13, 0, -49, -62, 12, -57, 44, -67, 28, 37, -31, -32, 10, 55, 56, 56, 55, 10, 3, 4, 54, 47, -99, -99, 60, 83, -62)

	zg.piecePoints[0] = 100
	zg.piecePoints[1] = 280
	zg.piecePoints[2] = 320
	zg.piecePoints[3] = 479
	zg.piecePoints[4] = 929
	zg.piecePoints[5] = 60000
}

func (zg *ZimuaGame) createMoveScore(move chess.Move, score int, killer bool) MoveScore {
	return MoveScore{
		move:      move,
		score:     score,
		killer:    killer,
		inCheck:   move.HasTag(chess.Check),
		capture:   false,
		promotion: false,
	}
}

func (zg *ZimuaGame) getMoves(pos *chess.Position, depth int) []MoveScore {

	var moves []MoveScore

	validMoves := pos.ValidMoves()
	board := pos.Board()

	for _, mv := range validMoves {
		isPromo := mv.Promo() != chess.NoPieceType
		if isPromo && mv.Promo() != chess.Queen {
			continue
		}

		score := 0
		pieceFrom := board.Piece(chess.Square(mv.S1()))
		pieceTo := board.Piece(chess.Square(mv.S2())).Type()
		pieceType := pieceFrom.Type()
		isCapture := mv.HasTag(chess.Capture)
		isEnPassant := mv.HasTag(chess.EnPassant)

		if isCapture {
			score += 100
		}
		if isEnPassant {
			score += 10
		}

		if pieceType == chess.King {
			score -= 10
		} else if pieceType == chess.Bishop || pieceType == chess.Knight {
			score += 9
		} else if pieceType == chess.Rook {
			score += 8
		} else if pieceType == chess.Queen {
			score += 7
		}
		if pieceTo != chess.King {
			pieceIdx := 0
			if pieceTo == chess.Knight {
				pieceIdx = 1
			} else if pieceTo == chess.Bishop {
				pieceIdx = 2
			} else if pieceTo == chess.Rook {
				pieceIdx = 3
			} else if pieceTo == chess.Queen {
				pieceIdx = 4
			} else if pieceTo == chess.King {
				pieceIdx = 5
			}
			score += zg.piecePoints[pieceIdx]
		}

		ms := zg.createMoveScore(*mv, score, false)
		ms.capture = isCapture
		ms.promotion = isPromo

		moves = append(moves, ms)
	}

	sort.Slice(moves, func(i, j int) bool {
		return moves[i].score > moves[j].score
	})

	return moves
}

func (zg *ZimuaGame) pieceScoring(b *chess.Board) int {

	var pieceScoreWhite int = 0
	var piecePosWhite int = 0
	var pieceScoreBlack int = 0
	var piecePosBlack int = 0

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
	var pos uint64 = 1

	for i := 0; i < 64; i++ {

		if allWhiteBBs&pos > 0 {
			if bbWhiteKing&pos > 0 {
				pieceScoreWhite += zg.piecePoints[5]
				piecePosWhite += zg.posPointsWhite[5][i]
			} else if bbWhiteQueen&pos > 0 {
				pieceScoreWhite += zg.piecePoints[4]
				piecePosWhite += zg.posPointsWhite[4][i]
			} else if bbWhiteRook&pos > 0 {
				pieceScoreWhite += zg.piecePoints[3]
				piecePosWhite += zg.posPointsWhite[3][i]
			} else if bbWhiteBishop&pos > 0 {
				pieceScoreWhite += zg.piecePoints[2]
				piecePosWhite += zg.posPointsWhite[2][i]
			} else if bbWhiteKnight&pos > 0 {
				pieceScoreWhite += zg.piecePoints[1]
				piecePosWhite += zg.posPointsWhite[1][i]
			} else if bbWhitePawn&pos > 0 {
				pieceScoreWhite += zg.piecePoints[0]
				piecePosWhite += zg.posPointsWhite[0][i]
			}
		}

		if allBlackBBs&pos > 0 {
			if bbBlackKing&pos > 0 {
				pieceScoreBlack += zg.piecePoints[5]
				piecePosBlack += zg.posPointsBlack[5][i]
			} else if bbBlackQueen&pos > 0 {
				pieceScoreBlack += zg.piecePoints[4]
				piecePosBlack += zg.posPointsBlack[4][i]
			} else if bbBlackRook&pos > 0 {
				pieceScoreBlack += zg.piecePoints[3]
				piecePosBlack += zg.posPointsBlack[3][i]
			} else if bbBlackBishop&pos > 0 {
				pieceScoreBlack += zg.piecePoints[2]
				piecePosBlack += zg.posPointsBlack[2][i]
			} else if bbBlackKnight&pos > 0 {
				pieceScoreBlack += zg.piecePoints[1]
				piecePosBlack += zg.posPointsBlack[1][i]
			} else if bbBlackPawn&pos > 0 {
				pieceScoreBlack += zg.piecePoints[0]
				piecePosBlack += zg.posPointsBlack[0][i]
			}
		}

		pos = pos << 1
	}

	scoreWhite := pieceScoreWhite + piecePosWhite
	scoreBlack := pieceScoreBlack + piecePosBlack

	return scoreWhite - scoreBlack
}

func (zg *ZimuaGame) qsearch(pos *chess.Position, standPat int) int {

	legalMoves := pos.ValidMoves()

	for _, move := range legalMoves {
		if !move.HasTag(chess.Capture) {
			continue
		}

		newPos := pos.Update(move)
		score := zg.pieceScoring(newPos.Board())

		if pos.Turn() == chess.Black {
			if score < standPat {
				return zg.minValue
			}
		} else {
			if score > standPat {
				return zg.minValue
			}
		}
	}

	return standPat
}

func (zg *ZimuaGame) alphaBetaNM(pos *chess.Position, depth int, alpha int, beta int, maxPlayer bool, startDepth int, inCheck bool, isNull bool, siblings []MoveScore) MoveScore {

	if depth == 0 {
		score := 0
		if pos.Status() == chess.Checkmate {
			score = checkmate
		} else {
			score = zg.pieceScoring(pos.Board())
		}

		if pos.Turn() == chess.Black {
			score = score * -1
		}

		mv := zg.createMoveScore(zg.nilMove, score, false)
		return mv
	}

	legalMoves := zg.getMoves(pos, depth)

	if len(legalMoves) == 0 {
		if pos.Turn() == chess.White {
			return MoveScore{
				score: zg.minValue,
			}
		}
		return MoveScore{
			score: zg.maxValue,
		}
	}

	moveCount := 0
	bestMove := legalMoves[0]
	value := -99999999

	allowLMR := depth >= 3 && !inCheck

	if !inCheck && depth >= 3 && depth != startDepth && !isNull {
		newPos := pos.NullMove()
		status := newPos.Status()
		if status != chess.Stalemate && status != chess.FivefoldRepetition {
			newSiblings := make([]MoveScore, depth-3)
			nmRes := zg.alphaBetaNM(newPos, depth-3, -beta, -beta+1, true, startDepth, false, true, newSiblings)

			if -nmRes.score >= beta {
				nmRes.score = nmRes.score * -1
				return nmRes
			}
		}
	}

	for _, mv := range legalMoves {
		moveCount++
		zg.moveSearched++

		newDepth := depth - 1
		isLMR := false

		if allowLMR && moveCount >= 4 && !mv.capture && !mv.promotion {
			isLMR = true
			newDepth--

			if moveCount >= 6 {
				newDepth--
			}
		}

		newPos := pos.Update(&mv.move)

		newSiblings := make([]MoveScore, newDepth)
		res := zg.alphaBetaNM(newPos, newDepth, -beta, -alpha, false, startDepth, mv.inCheck, false, newSiblings)
		score := -res.score
		if score > alpha && isLMR {
			newSiblings = make([]MoveScore, depth-1)
			res = zg.alphaBetaNM(newPos, depth-1, -beta, -alpha, false, startDepth, mv.inCheck, false, newSiblings)
			score = -res.score
		}
		status := newPos.Status()
		if status == chess.Checkmate {
			score = checkmate
		}

		newValue := max(value, score)
		if newValue > value {

			if status != chess.Stalemate && status != chess.FivefoldRepetition {

				value = newValue
				bestMove.move = mv.move
				bestMove.score = newValue

				for i, s := range newSiblings {
					if s.move.String() != "a1a1" {
						siblings[i] = s
					}
				}

				siblings[depth-1] = bestMove
			}
		}
		alpha = max(alpha, value)

		if alpha >= beta {
			zg.moveSearched += len(legalMoves) - moveCount
			break
		}

	}

	return bestMove
}

func (zg *ZimuaGame) calcMove(g *chess.Game, depth int, alpha int, beta int, inCheck bool, siblings []MoveScore) MoveScore {
	maxPlayer := true
	if g.Position().Turn() == chess.Black {
		maxPlayer = false
	}
	res := zg.alphaBetaNM(g.Position(), depth, alpha, beta, maxPlayer, depth, inCheck, false, siblings)

	return res
}

func (zg *ZimuaGame) evaluate(g *chess.Game, inCheck bool) (bool, chess.Move) {

	if zg.doOpen {

		openMove := zg.openingMove(g)

		if openMove != nil {
			g.Move(openMove)
			return false, *openMove
		}
		zg.doOpen = false
	}

	minEval := zg.minValue
	maxEval := zg.maxValue
	alpha := minEval
	beta := maxEval

	ply := 0
	maxPly := 99

	totalStart := time.Now()
	totalElapsed := 0.0
	bestMove := zg.nilMove
	moveInCheck := false

	response(fmt.Sprintf("# times @ %v\n", zg.timeControl.timePerMove))
	response("# 16+16 pieces, centr = (1,1) R=40\n")

	log.Println("start loop")
	for ply < maxPly {
		start := time.Now()
		zg.moveSearched = 0
		ply++

		siblings := make([]MoveScore, ply)
		res := zg.calcMove(g, ply, alpha, beta, inCheck, siblings)

		if math.Abs(float64(res.score)) == float64(checkmate) {
			bestMove = res.move
			inCheck = true
			break
		}

		if res.score > alpha && res.score < beta {
			alpha = res.score - 500
			beta = res.score + 500

			if alpha < minEval {
				alpha = minEval
			}
			if beta > maxEval {
				beta = maxEval
			}
		} else if alpha != minEval {
			alpha = minEval
			beta = maxEval
			ply--
		}

		t := time.Now()
		elapsed := t.Sub(start)

		moves := ""
		for i := len(siblings) - 1; i >= 0; i-- {
			moves += siblings[i].move.String() + " "
		}

		if g.Position().Turn() == chess.Black {
			res.score *= -1
		}

		response(fmt.Sprintf("%3v %6v %8v %10v %v\n", ply, res.score, int(elapsed)/1000000000, zg.moveSearched, moves))

		totalElapsed = t.Sub(totalStart).Seconds()
		zg.timeControl.totalNodes += zg.moveSearched

		bestMove = res.move
		moveInCheck = res.inCheck

		if totalElapsed > zg.timeControl.timePerMove {
			break
		}
	}
	zg.timeControl.totalElapsed += totalElapsed

	nps := float64(zg.timeControl.totalNodes) / zg.timeControl.totalElapsed
	log.Printf("move chosen: %v\tnps: %.2f\n", bestMove.String(), nps)

	g.Move(&bestMove)

	zg.timeControl.moveCount++

	timeRemaining := zg.timeControl.remainingTime - totalElapsed
	timePerMove := timeRemaining / (zg.timeControl.expectedMoves - zg.timeControl.moveCount)

	zg.timeControl.timePerMove = timePerMove
	zg.timeControl.remainingTime = timeRemaining

	log.Println("returning to xBoard")

	return moveInCheck, bestMove
}

func (zg *ZimuaGame) openingMove(game *chess.Game) *chess.Move {
	prevMoves := game.Moves()
	moveIndex := len(prevMoves)
	opennings := Possible(game.Moves())

	_ = opennings
	_ = prevMoves

	possibleOpenings := []*Opening{}
	for _, o := range opennings {
		moves := o.Game().Moves()
		if len(moves) > moveIndex {
			possibleOpenings = append(possibleOpenings, o)
		}
	}

	if len(possibleOpenings) > 0 {
		opening := possibleOpenings[rand.Intn(len(possibleOpenings))]
		moves := opening.Game().Moves()
		return moves[moveIndex]
	}

	return nil
}
