package main

import (
	"time"

	"github.com/dominicplouffe/chess"
)

func (zg *ZimuaGame) qsearch(alpha int, beta int, pos *chess.Position) int {

	if zg.moveSearched%10000 == 0 {
		t := time.Now()
		zg.timeControl.totalElapsed = t.Sub(zg.timeControl.startTime).Seconds()
		if zg.timeControl.totalElapsed > zg.timeControl.timePerMove {
			zg.stopped = true
		}
	}

	zg.moveSearched++

	if pos.Status() == chess.FivefoldRepetition || pos.Status() == chess.ThreefoldRepetition {
		return 0
	}

	score := zg.pieceScoring(pos)

	if score >= beta {
		return beta
	}

	if score > alpha {
		alpha = score
	}

	legalMoves := pos.ValidMoves()
	score = zg.minValue
	for _, move := range legalMoves {
		isCapture := move.HasTag(chess.Capture)

		if !isCapture {
			continue
		}

		newPos := pos.Update(move)

		score = zg.qsearch(-beta, -alpha, newPos)
		score *= -1

		if zg.stopped {
			return 0
		}

		if score > alpha {
			if score >= beta {
				return beta
			}
			alpha = score
		}
	}

	return alpha
}
