package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/dominicplouffe/chess"
)

func (zg *ZimuaGame) search(g *chess.Game, inCheck bool) (bool, chess.Move) {

	// fmt.Println("Searching.............")
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
	ply := 0
	maxPly := 99
	zg.timeControl.elapsedTime = 0
	zg.timeControl.startTime = time.Now()
	zg.stopped = false

	bestMove := zg.nilMove
	moveInCheck := false

	response(fmt.Sprintf("# times @ %v\n", zg.timeControl.timePerMove))
	response("# 16+16 pieces, centr = (1,1) R=40\n")

	log.Println("start loop")
	for ply < maxPly {
		// fmt.Printf("%v %v\n", ply, maxPly)
		zg.moveSearched = 0
		ply++

		res := zg.alphaBeta(g.Position(), ply, zg.minValue, zg.maxValue, ply, inCheck, false)
		// fmt.Printf("%v\n", res.breakMove)

		if zg.stopped {
			break
		}

		fmt.Printf("%v\n", res.score)
		if math.Abs(float64(res.score)) == float64(checkmate) {
			bestMove = res.move
			inCheck = true
			break
		}

		// moves := ""
		// if siblings[0].move.String() != "a1a1" {
		// 	for i := len(siblings) - 1; i >= 0; i-- {
		// 		moves += siblings[i].move.String() + " "
		// 	}
		// } else {
		// 	res.score = checkmate
		// }

		if g.Position().Turn() == chess.Black {
			res.score *= -1
		}

		response(fmt.Sprintf("%3v %6v %8v %10v\n", ply, res.score, int(zg.timeControl.totalElapsed)/10000000, zg.moveSearched))

		bestMove = res.move
		moveInCheck = res.inCheck

		// fmt.Printf("End of Loop\n")
	}

	nps := float64(zg.timeControl.totalNodes) / zg.timeControl.totalElapsed
	log.Printf("move chosen: %v\tnps: %.2f\n", bestMove.String(), nps)

	g.Move(&bestMove)

	zg.timeControl.moveCount++

	timeRemaining := zg.timeControl.remainingTime - zg.timeControl.totalElapsed
	timePerMove := timeRemaining / (zg.timeControl.expectedMoves - zg.timeControl.moveCount)

	zg.timeControl.timePerMove = timePerMove
	zg.timeControl.remainingTime = timeRemaining

	log.Println("returning to xBoard")

	return moveInCheck, bestMove
}
