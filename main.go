package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dominicplouffe/chess"
	"github.com/pkg/profile"
)

var wrt = bufio.NewWriter(os.Stdout)

func main() {

	args := os.Args[1:]

	if len(args) > 1 && args[1] == "-profile" {
		defer profile.Start().Stop()
	}
	rand.Seed(time.Now().UnixNano())

	f, err := os.OpenFile("zimua.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	if len(args) >= 1 && args[0] == "-uci" {
		xBoard()
	} else if len(args) >= 1 && args[0] == "-cpu" {
		computerVSComputer()
	} else if len(args) >= 1 && args[0] == "-human" {
		computerVSHuman()
	} else {
		// fmt.Println("Usage: ./engine.go [-uci|-cpu|-human] [-profile]")
		xBoard()
	}

}

func response(value string) {
	log.Printf(fmt.Sprintf("<< %v", value))
	wrt.WriteString(fmt.Sprintf("%v\n", value))
	wrt.Flush()
}

func computerVSHuman() {

	// fen, _ := chess.FEN("3R4/1r3ppk/7p/p7/1B6/2P5/P1b2PPP/6K1 w - - 0 1")
	// game := chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

	game := chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))
	zg := Zimua("White", 15.0)

	reader := bufio.NewReader(os.Stdin)

	for {
		zg.inCheck, _ = zg.search(game, zg.inCheck)
		fmt.Println(game.Position().Board().Draw())
		fmt.Println(game.Position().String())

		foundMove := false
		for {
			fmt.Print("Your Move -> ")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			fmt.Println(text)

			for _, move := range game.ValidMoves() {
				if move.String() == text {
					game.Move(move)
					zg.inCheck = move.HasTag(chess.Check)
					foundMove = true
				}
			}

			if foundMove {
				fmt.Println(game.Position().Board().Draw())
				fmt.Println(game.Position().String())
				break
			} else {
				fmt.Println("Move not found in list of moves")
			}
		}
		if game.Outcome() != chess.NoOutcome {
			print(game.Outcome)
			break
		}
	}
}

func computerVSComputer() {
	zg := Zimua("White", 5.0)
	zg2 := Zimua("Black", 5.0)

	game := chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))

	for {
		zg2.inCheck, _ = zg.search(game, zg.inCheck)
		fmt.Println(game.Position().Board().Draw())
		fmt.Println(game.Position().String())

		zg.inCheck, _ = zg2.search(game, zg2.inCheck)
		fmt.Println(game.Position().Board().Draw())
		fmt.Println(game.Position().String())

		if game.Outcome() != chess.NoOutcome {
			print(game.Outcome())
			print(game.String())
			break
		}
	}
}

func xBoard() {

	isForceGame := false
	maxTime := 5
	color := "white"
	game := chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))
	_ = game
	zg := Zimua("Zimua v2 bishop", 5.0)

	reader := bufio.NewReader(os.Stdin)

	for {
		cmd, _ := reader.ReadString('\n')
		cmd = strings.Replace(cmd, "\n", "", -1)

		log.Printf("Command: %v\n", cmd)

		if cmd == "xboard" {

			response("tellics say     %%s Zimua Engine\n")
			response("tellics say     (c) dplouffe Analytics Inc.\n")
		} else if cmd == "new" || cmd == "post" {
			game = chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))
			zg = Zimua("Zimua v2 bishop", 5.0)
			response("Zimua Ready\n")
		} else if cmd == "protover 2" {
			response(fmt.Sprintf("feature myname=\"%v\"\n", zg.name))
			response("feature ping=1\n")
			response("feature san=0\n")
			response("feature sigint=0\n")
			// response("feature sigterm=0\n")
			response("feature setboard=1\n")
			response("feature debug=1\n")
			response("feature time=0\n")
			response("feature done=1\n")
		} else if strings.HasPrefix(cmd, "ping") {
			parts := strings.Split(cmd, " ")
			n := parts[len(parts)-1]
			response(fmt.Sprintf("pong %v\n", n))
		} else if strings.HasPrefix(cmd, "setboard") {
			parts := []rune(cmd)
			fenStr := string(parts[9:])
			fen, _ := chess.FEN(fenStr)
			game = chess.NewGame(fen, chess.UseNotation(chess.LongAlgebraicNotation{}))

		} else if cmd == "go" {
			response("#go\n")
			if color == "white" {
				xBoardPlay(game, &zg)
			}
		} else if cmd == "fen" {
			response(fmt.Sprintf("%v\n", game.Position().String()))

		} else if cmd == "white" {
			response("#Playing white\n")
		} else if cmd == "black" {
			color = "black"
			response("#Playing black\n")
		} else if cmd == "quit" {
			break

		} else if cmd == "force" {
			isForceGame = true
		} else if strings.HasPrefix(cmd, "level") {
			maxTime, _ = strconv.Atoi(strings.Split(cmd, " ")[2])
			game = chess.NewGame(chess.UseNotation(chess.LongAlgebraicNotation{}))
			zg = Zimua("Zimua Chess Engine v2", float64(maxTime))
		} else {
			matched, _ := regexp.MatchString(`^[a-h][1-8][a-h][1-8].?$`, cmd)

			if matched {
				foundMove := false
				for _, move := range game.ValidMoves() {
					if move.String() == cmd {
						game.Move(move)
						zg.inCheck = move.HasTag(chess.Check)
						zg.moveCount++
						foundMove = true
					}
				}

				if !foundMove {
					response(fmt.Sprintf("illegal move: %v", cmd))
					continue
				}

				if game.Outcome() != chess.NoOutcome {
					response("#game_over\n")
				} else if !isForceGame {
					xBoardPlay(game, &zg)
				}
			} else {
				response(fmt.Sprintf("#ignored command : '%v'", cmd))
			}

		}
	}
}

func xBoardPlay(game *chess.Game, zg *ZimuaGame) {
	log.Println("start")

	// if zg.moveCount < 4 {
	// 	moves := game.ValidMoves()
	// 	move := moves[rand.Intn(len(moves))]
	// 	game.Move(move)
	// 	response(fmt.Sprintf("move %v\n", move.String()))
	// 	zg.inCheck = false
	// } else {
	// 	inCheck, move := zg.search(game, zg.inCheck)
	// 	response(fmt.Sprintf("move %v\n", move.String()))
	// 	zg.inCheck = inCheck
	// }

	inCheck, move := zg.search(game, zg.inCheck)
	response(fmt.Sprintf("move %v\n", move.String()))
	zg.inCheck = inCheck

	zg.moveCount++
	log.Println("finished")

	if game.Position().Status() == chess.Checkmate {
		response("#checkmate\n")
	} else if game.Position().Status() == chess.Stalemate {
		response("#stalemate\n")
	} else if game.Position().Status() == chess.FivefoldRepetition {
		response("#result : draw {stalemate}\n")
	} else if game.Position().Status() == chess.InsufficientMaterial {
		response("#result : draw {stalemate}\n")
	} else if game.Position().Status() == chess.DrawOffer {
		response("#result : draw {stalemate}\n")
	}
}
