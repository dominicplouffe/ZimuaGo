package main

import (
	"fmt"
	"time"

	"github.com/dominicplouffe/chess"
)

func (zg *ZimuaGame) alphaBeta(pos *chess.Position, depth int, alpha int, beta int, startDepth int, inCheck bool, isNull bool) MoveScore {

	// fmt.Printf("Depth %v\n", depth)
	if depth == 0 {
		qscore := zg.qsearch(alpha, beta, pos) // zg.pieceScoring(pos)
		return MoveScore{
			score: qscore,
		}
	}

	if zg.moveSearched%10000 == 0 {
		t := time.Now()
		zg.timeControl.totalElapsed = t.Sub(zg.timeControl.startTime).Seconds()
		fmt.Printf("%v: %v\n", zg.timeControl.totalElapsed, zg.timeControl.timePerMove)
		if zg.timeControl.totalElapsed > zg.timeControl.timePerMove {
			zg.stopped = true
		}
	}

	zg.moveSearched++

	if pos.Status() == chess.FivefoldRepetition || pos.Status() == chess.ThreefoldRepetition {
		return MoveScore{
			score: 0,
		}
	}

	if depth == zg.maxDepth {
		return MoveScore{
			score: zg.pieceScoring(pos),
		}
	}

	if inCheck {
		depth++
	}

	score := zg.maxValue

	thash := zg.transposiiton[depth][pos.Hash()]
	if thash != 0 {
		return MoveScore{
			score: thash,
		}
	}

	// Null Move Code Block
	if !isNull && !inCheck && depth >= 4 {
		fmt.Println("Null Move")
		res := zg.alphaBeta(pos, depth-4, -beta, -beta+1, startDepth, inCheck, true)

		score = res.score * -1
		pos.Update(&zg.nilMove)

		if zg.stopped {
			return MoveScore{
				score: 0,
			}
		}

		if score >= beta && score < checkmate {
			return MoveScore{
				score: beta,
			}
		}
	}

	legalMoves := zg.getMoves(pos, depth)
	bestMove := MoveScore{
		score: zg.minValue,
	}
	bestPos := pos
	moveCount := 0

	for _, mv := range legalMoves {
		moveCount++
		zg.moveSearched++

		if zg.stopped {
			return MoveScore{
				score: 0,
			}
		}

		newPos := pos.Update(&mv.move)
		res := zg.alphaBeta(
			newPos,
			depth-1,
			-beta,
			-alpha,
			startDepth,
			inCheck,
			false,
		)
		// res.score = res.score * -1

		if res.score > bestMove.score {
			bestMove = mv

			if res.score > alpha {
				if res.score >= beta {
					if mv.capture {
						posMap := map[[16]byte]int{newPos.Hash(): res.score}
						zg.transposiiton[depth] = posMap

						bestMove.score = beta

						return bestMove
					}
				}
				alpha = -res.score
			}
		}
	}

	posMap := map[[16]byte]int{bestPos.Hash(): bestMove.score}
	zg.transposiiton[depth] = posMap

	// bestMove.score = alpha
	// fmt.Println("return 333")
	return bestMove

}
