package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"time"

	"github.com/dominicplouffe/chess"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var checkmate = 99999999
var stalemate = 88888888
var doubleBishopBonus = 75
var connectedRooksBonus = 25

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
	minValue       int
	maxValue       int
	name           string
	gamestage      int
	transposiiton  map[int]map[[16]byte]int
	killer         map[[16]byte]MoveScore
}

//Zimua creates an instance of the Zimua chess engine
func Zimua(name string, maxMinutes float64) ZimuaGame {
	initSquareIndexes()

	zg := ZimuaGame{
		posPointsBlack: make(map[int][]int),
		posPointsWhite: make(map[int][]int),
		piecePoints:    make(map[int]int),
		moveSearched:   0,
		nilMove:        chess.Move{},
		minValue:       -checkmate,
		maxValue:       checkmate,
		timeControl:    getTimeControl(maxMinutes),
		name:           name,
		moveCount:      0,
		gamestage:      0,
		transposiiton:  make(map[int]map[[16]byte]int),
		killer:         make(map[[16]byte]MoveScore),
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

	return tc
}

func (zg ZimuaGame) initGame() {
	zg.posPointsWhite[0] = append(zg.posPointsWhite[0],
		900, 900, 900, 900, 900, 900, 900, 900,
		78, 83, 86, 73, 102, 82, 85, 90,
		20, 29, 21, 44, 40, 31, 44, 20,
		20, 16, -2, 15, 14, 0, 15, 20,
		10, 3, 10, 20, 20, 1, 0, 10,
		10, 9, 5, 10, 10, -2, 3, 10,
		-10, 8, -7, -37, -36, -14, 3, -10,
		0, 0, 0, 0, 0, 0, 0, 0,
	)
	zg.posPointsWhite[1] = append(zg.posPointsWhite[1],
		-66, -53, -75, -75, -10, -55, -58, -70,
		-3, -6, 100, -36, 4, 62, -4, -14,
		10, 67, 100, 74, 73, 100, 62, -2,
		24, 24, 45, 37, 33, 41, 25, 17,
		-1, 5, 31, 21, 22, 35, 2, 0,
		-18, 10, 25, 22, 18, 25, 11, -14,
		-23, -15, 2, 0, 2, 0, -23, -20,
		-74, -23, -26, -24, -19, -35, -22, -69,
	)
	zg.posPointsWhite[2] = append(zg.posPointsWhite[2],
		-59, -78, -82, -76, -23, -107, -37, -50,
		-11, 20, 35, -42, -39, 31, 2, -22,
		-9, 39, -32, 41, 52, -10, 28, -14,
		25, 17, 20, 34, 26, 25, 15, 10,
		13, 10, 17, 23, 17, 16, 0, 7,
		14, 25, 24, 15, 8, 25, 20, 15,
		19, 20, 11, 6, 7, 6, 20, 16,
		-7, 2, -15, -12, -14, -15, -10, -10,
	)
	zg.posPointsWhite[3] = append(zg.posPointsWhite[3],
		35, 29, 33, 4, 37, 33, 56, 50,
		55, 29, 56, 67, 55, 62, 34, 60,
		19, 35, 28, 33, 45, 27, 25, 15,
		0, 5, 16, 13, 18, -4, -9, -6,
		-28, -35, -16, -21, -13, -29, -46, -30,
		-42, -28, -42, -25, -25, -35, -26, -46,
		-53, -38, -31, -26, -29, -43, -44, -53,
		-30, -24, -18, 5, -2, -18, -31, -32,
	)
	zg.posPointsWhite[4] = append(zg.posPointsWhite[4],
		6, 1, -8, -104, 69, 24, 88, 26,
		14, 32, 60, -10, 20, 76, 57, 24,
		-2, 43, 32, 60, 72, 63, 43, 2,
		1, -16, 22, 17, 25, 20, -13, -6,
		-14, -15, -2, -5, -1, -10, -20, -22,
		-30, -6, -13, -11, -16, -11, -16, -27,
		-36, -18, 0, -19, -15, -15, -21, -38,
		-39, -30, -31, -13, -31, -36, -34, -42,
	)
	zg.posPointsWhite[5] = append(zg.posPointsWhite[5],
		4, 54, 47, -99, -99, 60, 83, -62,
		-32, 10, 55, 56, 56, 55, 10, 3,
		-62, 12, -57, 44, -67, 28, 37, -31,
		-55, 50, 11, -4, -19, 13, 0, -49,
		-55, -43, -52, -28, -51, -47, -8, -50,
		-47, -42, -43, -79, -64, -32, -29, -32,
		-4, 3, -14, -50, -57, -18, 13, 4,
		17, 30, -3, -14, 6, -1, 40, 18,
	)

	zg.posPointsBlack[0] = append(zg.posPointsBlack[0],
		0, 0, 0, 0, 0, 0, 0, 0,
		-10, 8, -7, -37, -36, -14, 3, -10,
		10, 9, 5, 10, 10, -2, 3, 10,
		10, 3, 10, 20, 20, 1, 0, 10,
		20, 16, -2, 15, 14, 0, 15, 20,
		20, 29, 21, 44, 40, 31, 44, 20,
		78, 83, 86, 73, 102, 82, 85, 90,
		900, 900, 900, 900, 900, 900, 900, 900,
	)

	zg.posPointsBlack[1] = append(zg.posPointsBlack[1],
		-74, -23, -26, -24, -19, -35, -22, -69,
		-23, -15, 2, 0, 2, 0, -23, -20,
		-18, 10, 25, 22, 18, 25, 11, -14,
		-1, 5, 31, 21, 22, 35, 2, 0,
		24, 24, 45, 37, 33, 41, 25, 17,
		10, 67, 100, 74, 73, 100, 62, -2,
		-3, -6, 100, -36, 4, 62, -4, -14,
		-66, -53, -75, -75, -10, -55, -58, -70,
	)
	zg.posPointsBlack[2] = append(zg.posPointsBlack[2],
		-7, 2, -15, -12, -14, -15, -10, -10,
		19, 20, 11, 6, 7, 6, 20, 16,
		14, 25, 24, 15, 8, 25, 20, 15,
		13, 10, 17, 23, 17, 16, 0, 7,
		25, 17, 20, 34, 26, 25, 15, 10,
		-9, 39, -32, 41, 52, -10, 28, -14,
		-11, 20, 35, -42, -39, 31, 2, -22,
		-59, -78, -82, -76, -23, -107, -37, -50,
	)

	zg.posPointsBlack[3] = append(zg.posPointsBlack[3],
		-30, -24, -18, 5, -2, -18, -31, -32,
		-53, -38, -31, -26, -29, -43, -44, -53,
		-42, -28, -42, -25, -25, -35, -26, -46,
		-28, -35, -16, -21, -13, -29, -46, -30,
		0, 5, 16, 13, 18, -4, -9, -6,
		19, 35, 28, 33, 45, 27, 25, 15,
		55, 29, 56, 67, 55, 62, 34, 60,
		35, 29, 33, 4, 37, 33, 56, 50,
	)
	zg.posPointsBlack[4] = append(zg.posPointsBlack[4],
		-39, -30, -31, -13, -31, -36, -34, -42,
		-36, -18, 0, -19, -15, -15, -21, -38,
		-30, -6, -13, -11, -16, -11, -16, -27,
		-14, -15, -2, -5, -1, -10, -20, -22,
		1, -16, 22, 17, 25, 20, -13, -6,
		-2, 43, 32, 60, 72, 63, 43, 2,
		14, 32, 60, -10, 20, 76, 57, 24,
		6, 1, -8, -104, 69, 24, 88, 26,
	)
	zg.posPointsBlack[5] = append(zg.posPointsBlack[5],
		17, 30, -3, -14, 6, -1, 40, 18,
		-4, 3, -14, -50, -57, -18, 13, 4,
		-47, -42, -43, -79, -64, -32, -29, -32,
		-55, -43, -52, -28, -51, -47, -8, -50,
		-55, 50, 11, -4, -19, 13, 0, -49,
		-62, 12, -57, 44, -67, 28, 37, -31,
		-32, 10, 55, 56, 56, 55, 10, 3,
		4, 54, 47, -99, -99, 60, 83, -62,
	)

	zg.piecePoints[0] = 100   //PAWN
	zg.piecePoints[1] = 280   // KNIGHT
	zg.piecePoints[2] = 320   // BISHOP
	zg.piecePoints[3] = 479   // ROOK
	zg.piecePoints[4] = 929   // QUEEN
	zg.piecePoints[5] = 60000 //KING
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

	killer_move := zg.killer[pos.Hash()]

	if killer_move.score > 0 {
		killer_move.killer = true
		killer_move.score = 900
		moves = append(moves, killer_move)
	}

	validMoves := pos.ValidMoves()
	board := pos.Board()

	for _, mv := range validMoves {
		isPromo := mv.Promo() != chess.NoPieceType
		if isPromo && mv.Promo() != chess.Queen {
			continue
		}

		score := 0

		// newPos := pos.Update(mv)
		// thash := zg.transposiiton[depth][newPos.Hash()]
		thash := 0

		if thash != 0 {
			score += 1000
		}

		pieceFrom := board.Piece(chess.Square(mv.S1()))
		pieceTo := board.Piece(chess.Square(mv.S2())).Type()
		pieceType := pieceFrom.Type()
		isCapture := mv.HasTag(chess.Capture)
		kindSideCastle := mv.HasTag(chess.KingSideCastle)
		queenSideCastle := mv.HasTag(chess.QueenSideCastle)
		toCheck := mv.HasTag(chess.Check)
		pieceColor := pieceFrom.Color()

		if toCheck {
			score += 800
		}

		if isCapture {
			score += 100
		}
		if isPromo {
			score += 10
		}

		if kindSideCastle {
			score += 0
		}
		if queenSideCastle {
			score += 0
		}

		// if pieceType == chess.King {
		// 	score -= 10
		// } else if pieceType == chess.Bishop || pieceType == chess.Knight {
		// 	score += 9
		// } else if pieceType == chess.Rook {
		// 	score += 7
		// } else if pieceType == chess.Queen {
		// 	score += 8
		// }
		// if pieceTo != chess.King {
		// 	pieceIdx := 0
		// 	if pieceTo == chess.Knight {
		// 		pieceIdx = 1
		// 	} else if pieceTo == chess.Bishop {
		// 		pieceIdx = 2
		// 	} else if pieceTo == chess.Rook {
		// 		pieceIdx = 3
		// 	} else if pieceTo == chess.Queen {
		// 		pieceIdx = 4
		// 	} else if pieceTo == chess.King {
		// 		pieceIdx = 5
		// 	}
		// 	score += zg.piecePoints[pieceIdx]
		// }

		f := mv.S2().File()
		r := mv.S2().Rank()
		idx := (int(r) * 8) + int(f)

		pieceTypeIdx := 0
		if pieceType == chess.Knight {
			pieceTypeIdx = 1
		} else if pieceType == chess.Bishop {
			pieceTypeIdx = 2
		} else if pieceType == chess.Rook {
			pieceTypeIdx = 3
		} else if pieceType == chess.Queen {
			pieceTypeIdx = 4
		} else if pieceType == chess.King {
			pieceTypeIdx = 5
		}

		if pieceColor == chess.White {
			score += zg.posPointsWhite[pieceTypeIdx][idx]
		} else {
			score += zg.posPointsBlack[pieceTypeIdx][idx]
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

func (zg *ZimuaGame) scorePiece(p *chess.Position) int {
	return 0
}

func (zg *ZimuaGame) pieceScoring(p *chess.Position) int {

	var pieceScoreWhite int = 0
	var piecePosWhite int = 0
	var pieceScoreBlack int = 0
	var piecePosBlack int = 0

	var bishopWhite int = 0
	var bishopBlack int = 0
	var majorPieceCount int = 0

	b := p.Board()

	if p.Status() == chess.Checkmate {
		return checkmate
	}

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

	// _ = allWhiteBBs
	// _ = allBlackBBs

	var wpsqs []uint64
	var wnsqs []uint64
	var wrsqs []uint64
	var wbsqs []uint64
	var wqsqs []uint64
	var wksqs []uint64

	var bpsqs []uint64
	var bnsqs []uint64
	var brsqs []uint64
	var bbsqs []uint64
	var bqsqs []uint64
	var bksqs []uint64

	var sq uint64 = 1
	for i := 0; i < 64; i++ {
		isqr := bbWhitePawn & sq
		if isqr > 0 {
			wpsqs = append(wpsqs, isqr)
			pieceScoreWhite += zg.piecePoints[0]
			piecePosWhite += zg.posPointsWhite[0][squareIndexes[sq]]
		}

		isqr = bbWhiteKnight & sq
		if isqr > 0 {
			wnsqs = append(wnsqs, isqr)
			pieceScoreWhite += zg.piecePoints[1]
			piecePosWhite += zg.posPointsWhite[1][squareIndexes[sq]]
			majorPieceCount++
		}

		isqr = bbWhiteBishop & sq
		if isqr > 0 {
			wbsqs = append(wbsqs, isqr)
			pieceScoreWhite += zg.piecePoints[2]
			piecePosWhite += zg.posPointsWhite[2][squareIndexes[sq]]
			bishopWhite++
			majorPieceCount++
		}

		isqr = bbWhiteRook & sq
		if isqr > 0 {
			wrsqs = append(wrsqs, isqr)
			pieceScoreWhite += zg.piecePoints[3]
			piecePosWhite += zg.posPointsWhite[3][squareIndexes[sq]]
			majorPieceCount++
		}

		isqr = bbWhiteQueen & sq
		if isqr > 0 {
			wqsqs = append(wqsqs, isqr)
			pieceScoreWhite += zg.piecePoints[4]
			piecePosWhite += zg.posPointsWhite[4][squareIndexes[sq]]
			majorPieceCount++
		}

		isqr = bbWhiteKing & sq
		if isqr > 0 {
			wksqs = append(wksqs, isqr)
			pieceScoreWhite += zg.piecePoints[5]
			piecePosWhite += zg.posPointsWhite[5][squareIndexes[sq]]
		}

		// Black

		isqr = bbBlackPawn & sq
		if isqr > 0 {
			bpsqs = append(bpsqs, isqr)
			pieceScoreBlack += zg.piecePoints[0]
			piecePosBlack += zg.posPointsBlack[0][squareIndexes[sq]]
		}

		isqr = bbBlackKnight & sq
		if isqr > 0 {
			bnsqs = append(bnsqs, isqr)
			pieceScoreBlack += zg.piecePoints[1]
			piecePosBlack += zg.posPointsBlack[1][squareIndexes[sq]]
			majorPieceCount++
		}

		isqr = bbBlackBishop & sq
		if isqr > 0 {
			bbsqs = append(bbsqs, isqr)
			pieceScoreBlack += zg.piecePoints[2]
			piecePosBlack += zg.posPointsBlack[2][squareIndexes[sq]]
			bishopBlack++
			majorPieceCount++
		}

		isqr = bbBlackRook & sq
		if isqr > 0 {
			brsqs = append(brsqs, isqr)
			pieceScoreBlack += zg.piecePoints[3]
			piecePosBlack += zg.posPointsBlack[3][squareIndexes[sq]]
			majorPieceCount++
		}

		isqr = bbBlackQueen & sq
		if isqr > 0 {
			bqsqs = append(bqsqs, isqr)
			pieceScoreBlack += zg.piecePoints[4]
			piecePosBlack += zg.posPointsBlack[4][squareIndexes[sq]]
			majorPieceCount++
		}

		isqr = bbBlackKing & sq
		if isqr > 0 {
			bksqs = append(bksqs, isqr)
			pieceScoreBlack += zg.piecePoints[5]
			piecePosBlack += zg.posPointsBlack[5][squareIndexes[sq]]
		}

		sq = sq << 1
	}

	if majorPieceCount <= 5 {
		zg.gamestage = 2
	}

	wnmob := getKnightMobilitySquares(wnsqs, allWhiteBBs)
	wrmob, wrcon := getRookMobilitySquares(wrsqs, bbWhiteRook, allWhiteBBs, allBlackBBs)
	wbmob := getBishopMobilitySquares(wbsqs, allWhiteBBs, allBlackBBs)
	wqmob := getQueenMobilitySquares(wqsqs, allWhiteBBs, allBlackBBs)

	bnmob := getKnightMobilitySquares(bnsqs, allBlackBBs)
	brmob, brcon := getRookMobilitySquares(brsqs, bbBlackRook, allBlackBBs, allWhiteBBs)
	bbmob := getBishopMobilitySquares(bbsqs, allBlackBBs, allWhiteBBs)
	bqmob := getQueenMobilitySquares(bqsqs, allBlackBBs, allWhiteBBs)

	scoreWhite := pieceScoreWhite + piecePosWhite
	scoreBlack := pieceScoreBlack + piecePosBlack

	queenMobility := 0
	if zg.timeControl.moveCount > 10 {
		queenMobility = 4
	}
	scoreWhite += ((wnmob * 3) + (wrmob * 4) + (wbmob * 2) + (wqmob * queenMobility))
	scoreBlack += ((bnmob * 3) + (brmob * 4) + (bbmob * 2) + (bqmob * queenMobility))

	if wrcon {
		scoreWhite += connectedRooksBonus
	}
	if brcon {
		scoreBlack += connectedRooksBonus
	}

	// Double Bishop Bonus
	if bishopWhite == 2 {
		scoreWhite += doubleBishopBonus
	}
	if bishopBlack == 2 {
		scoreBlack += doubleBishopBonus
	}

	return scoreWhite - scoreBlack
}

func (zg *ZimuaGame) qsearch(pos *chess.Position) int {
	standPat := zg.pieceScoring(pos)

	legalMoves := pos.ValidMoves()
	for _, move := range legalMoves {
		isCapture := move.HasTag(chess.Capture)

		if !isCapture {
			continue
		}

		newPos := pos.Update(move)
		score := zg.pieceScoring(newPos)

		if score < standPat {
			return score
		}

	}
	return standPat
}

func (zg *ZimuaGame) alphaBetaNM(pos *chess.Position, depth int, alpha int, beta int, startDepth int, inCheck bool, isNull bool, siblings []MoveScore) MoveScore {

	if depth == 0 {
		score := 0

		if pos.Status() == chess.FivefoldRepetition {
			score = stalemate
		} else if pos.Status() == chess.ThreefoldRepetition {
			score = stalemate
		} else {
			score = zg.qsearch(pos) // zg.pieceScoring(pos.Board())
		}

		if pos.Turn() == chess.Black {
			score = score * -1
		}

		mv := zg.createMoveScore(zg.nilMove, score, false)
		return mv
	}

	legalMoves := zg.getMoves(pos, depth)

	if len(legalMoves) == 0 {

		score := zg.qsearch(pos)
		if pos.Turn() == chess.Black {
			score = score * -1
		}

		return MoveScore{
			score: score,
		}
	}

	moveCount := 0
	bestMove := legalMoves[0]
	value := -(checkmate + 1)

	allowLMR := depth >= 3 && !inCheck

	if !inCheck && depth >= 3 && depth != startDepth && !isNull {
		newPos := pos.NullMove()
		status := newPos.Status()
		if status != chess.Stalemate && status != chess.FivefoldRepetition {
			newSiblings := make([]MoveScore, depth-3)
			nmRes := zg.alphaBetaNM(newPos, depth-3, -beta, -beta+1, startDepth, false, true, newSiblings)

			if -nmRes.score >= beta {
				nmRes.score = nmRes.score * -1
				return nmRes
			}
		}
	}

	for _, mv := range legalMoves {

		// fmt.Println(depth, mv.move.S1(), mv.move.S2())
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

			if moveCount >= 8 && depth >= 6 {
				newDepth--
			}
		}

		newPos := pos.Update(&mv.move)

		// fmt.Println(newPos.Hash())
		// fmt.Println(zg.transposiiton)

		score := 0
		newSiblings := make([]MoveScore, newDepth)
		if newPos.Status() == chess.Checkmate {
			score = checkmate
		} else {

			thash := zg.transposiiton[depth][newPos.Hash()]
			if thash == 0 {
				// fmt.Println("a")
				res := zg.alphaBetaNM(newPos, newDepth, -beta, -alpha, startDepth, mv.inCheck, false, newSiblings)
				score = -res.score
			} else {
				// fmt.Println(thash)
				score = thash
			}
			// fmt.Println("score", depth, score, alpha, score > alpha, pos.Turn())
			if score > alpha && isLMR { //
				// fmt.Println("bbb")
				newSiblings = make([]MoveScore, depth-1)
				res := zg.alphaBetaNM(newPos, depth-1, -beta, -alpha, startDepth, mv.inCheck, false, newSiblings)
				score = -res.score
			}
			posMap := map[[16]byte]int{newPos.Hash(): score}
			zg.transposiiton[depth] = posMap

			if alpha >= beta {
				zg.killer[newPos.Hash()] = mv
				break
			}
		}

		if score == stalemate || score == -stalemate {
			continue
		}

		newValue := max(value, score)

		// fmt.Println("***", newValue, value)
		if (newValue > value) || (mv.killer && Abs(score) > 250) {
			value = newValue
			bestMove.move = mv.move
			bestMove.score = newValue

			// fmt.Println("ccc", value, score, mv.move.String())

			for i, s := range newSiblings {
				if s.move.String() != "a1a1" {
					siblings[i] = s
				}
			}

			siblings[depth-1] = bestMove
		}
		alpha = max(alpha, value)

		if alpha >= beta {
			zg.moveSearched += len(legalMoves) - moveCount
			break
		}

	}
	return bestMove
}

func (zg *ZimuaGame) search(g *chess.Game, inCheck bool) (bool, chess.Move) {

	bookMove := zg.openingMove(g)

	if bookMove != nil {
		validMode := false
		for _, move := range g.ValidMoves() {
			if move.String() == bookMove.String() {
				g.Move(bookMove)
				zg.inCheck = bookMove.HasTag(chess.Check)
				zg.moveCount++
				validMode = true
			}
		}

		if validMode {
			return zg.inCheck, *bookMove
		}
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
		// fmt.Println(ply)
		// fmt.Println(g.Position().Board().Draw())
		res := zg.alphaBetaNM(g.Position(), ply, alpha, beta, ply, inCheck, false, siblings)

		if math.Abs(float64(res.score)) == float64(checkmate) {
			bestMove = res.move
			inCheck = true
			break
		}

		if res.score == checkmate {
			alpha = minEval
			beta = maxEval
		} else if res.score > alpha && res.score < beta {
			alpha = res.score - 50
			beta = res.score + 50

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
		if siblings[0].move.String() != "a1a1" {
			for i := len(siblings) - 1; i >= 0; i-- {
				moves += siblings[i].move.String() + " "
			}
		} else {
			res.score = checkmate
		}

		if g.Position().Turn() == chess.Black {
			res.score *= -1
		}

		response(fmt.Sprintf("%3v %6v %8v %10v %v\n", ply, res.score, int(elapsed)/10000000, zg.moveSearched, moves))

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

	var moves []string
	for _, mv := range game.Moves() {
		moves = append(moves, mv.String())
	}
	nextMove, found := findMove(moves)

	if !found {
		return nil
	}

	for _, mv := range game.ValidMoves() {
		if mv.String() == nextMove {
			return mv
		}
	}

	return nil
}
